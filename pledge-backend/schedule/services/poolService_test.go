package services

import (
	"errors"
	"gorm.io/gorm"
	"math/big"
	api "pledge-backend/api/models"
	"pledge-backend/api/models/request"
	"pledge-backend/config"
	"pledge-backend/db"
	"pledge-backend/internal/chain"
	"pledge-backend/internal/testutil"
	"pledge-backend/schedule/models"
	"strings"
	"testing"
)

func sample() *chain.Snapshot {
	deployment := config.Config.Deployment
	amount := new(big.Int).Lsh(big.NewInt(1), 200).String()
	token := chain.Token{Address: deployment.Stable, Symbol: "TEST", Decimals: 18, Price: "100000000"}
	native := chain.Token{Address: "0x0000000000000000000000000000000000000000", Symbol: "ETH", Decimals: 18, Price: "200000000000"}
	return &chain.Snapshot{ChainID: "11155111", PoolContract: deployment.Pool, BlockNumber: "123", BlockHash: "0x" + strings.Repeat("a", 64), LendFee: "0", BorrowFee: "0", Tokens: map[string]chain.Token{token.Address: token, native.Address: native}, Pools: []chain.Pool{{PoolIndex: 0, DisplayPoolID: 1, LendToken: token.Address, BorrowToken: native.Address, SP: deployment.SP, JP: deployment.JP, LendSupply: amount, SettleAmountLend: amount, FinishAmountLend: amount}, {PoolIndex: 1, DisplayPoolID: 2, LendToken: token.Address, BorrowToken: native.Address, SP: deployment.SP, JP: deployment.JP, LendSupply: "555", FinishAmountLend: "777"}}}
}

func TestSnapshotUpdatesDistributionAndPairsEveryPoolCorrectly(t *testing.T) {
	cache := testutil.Database(t)
	if err := models.InitTable(); err != nil {
		t.Fatal(err)
	}
	_ = cache.Set("session:admin", "keep-me")
	snapshot := sample()
	service := NewPool()
	if err := service.Persist(snapshot); err != nil {
		t.Fatal(err)
	}
	snapshot.Pools[0].FinishAmountLend = "999999999999999999999999999999999999999999999999999999999999999"
	snapshot.BlockNumber = "124"
	if err := service.Persist(snapshot); err != nil {
		t.Fatal(err)
	}
	err, count, pools := api.NewPool().Pagination(&request.Search{ChainID: 11155111, Page: 1, PageSize: 20})
	if err != nil || count != 2 || len(pools) != 2 {
		t.Fatalf("pagination: %v %d %+v", err, count, pools)
	}
	for _, p := range pools {
		if p.PoolID == 1 && p.Pooldata.FinishAmountLend != snapshot.Pools[0].FinishAmountLend {
			t.Fatal("lost large amount or distribution-only update")
		}
		if p.PoolID == 2 && p.Pooldata.FinishAmountLend != "777" {
			t.Fatal("cross-pool data pairing")
		}
	}
	if got, _ := cache.Get("session:admin"); got != "keep-me" {
		t.Fatal("sync modified session")
	}
	// Identical chain and display IDs belonging to another deployment must not leak into this API.
	snapshot.PoolContract = "0x0000000000000000000000000000000000001234"
	if err := service.Persist(snapshot); err != nil {
		t.Fatal(err)
	}
	_, count, _ = api.NewPool().Pagination(&request.Search{ChainID: 11155111, Page: 1, PageSize: 20})
	if count != 2 {
		t.Fatal("cross-contract records leaked")
	}
	_, count, _ = api.NewPool().Pagination(&request.Search{ChainID: 11155111, Page: 1, PageSize: 20, LendTokenSymbol: "' OR 1=1 --"})
	if count != 0 {
		t.Fatal("filter was interpreted as SQL")
	}
}

func TestDatabaseFailureRollsBackAndDoesNotAdvanceCache(t *testing.T) {
	testutil.Database(t)
	if err := models.InitTable(); err != nil {
		t.Fatal(err)
	}
	service := NewPool()
	snapshot := sample()
	if err := service.Persist(snapshot); err != nil {
		t.Fatal(err)
	}
	key := "snapshot:" + snapshot.ChainID + ":" + snapshot.PoolContract
	oldCache, _ := db.RedisGet(key)
	if err := db.Mysql.Callback().Create().Before("gorm:create").Register("fail_distribution", func(tx *gorm.DB) {
		if tx.Statement.Table == "pooldata" {
			tx.AddError(errors.New("injected write failure"))
		}
	}); err != nil {
		t.Fatal(err)
	}
	snapshot.Pools[0].LendSupply = "changed"
	snapshot.BlockNumber = "125"
	if err := service.Persist(snapshot); err == nil {
		t.Fatal("expected rollback")
	}
	var base models.PoolBase
	if err := db.Mysql.Where("pool_id=? AND pool_contract=?", 1, snapshot.PoolContract).First(&base).Error; err != nil {
		t.Fatal(err)
	}
	if base.LendSupply == "changed" || base.SyncBlockNumber != "123" {
		t.Fatal("base update survived rolled-back transaction")
	}
	nextCache, _ := db.RedisGet(key)
	if string(oldCache) != string(nextCache) {
		t.Fatal("cache advanced after failed persistence")
	}
	var state models.SyncState
	db.Mysql.First(&state)
	if state.BlockNumber != "123" {
		t.Fatal("watermark advanced")
	}
	_ = db.Mysql.Callback().Create().Remove("fail_distribution")
	if err := service.Persist(snapshot); err != nil {
		t.Fatalf("retry failed: %v", err)
	}
}
