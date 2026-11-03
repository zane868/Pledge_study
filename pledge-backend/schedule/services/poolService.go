package services

import (
	"context"
	"encoding/json"
	"fmt"
	"pledge-backend/config"
	"pledge-backend/db"
	"pledge-backend/internal/chain"
	"pledge-backend/log"
	"pledge-backend/schedule/models"
	"strconv"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type poolService struct{ Reader chain.Reader }

func NewPool() *poolService {
	return &poolService{Reader: chain.Client{URL: config.Config.TestNet.NetUrl, Deployment: config.Config.Deployment, Confirmations: 12}}
}

func (s *poolService) UpdateAllPoolInfo() {
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()
	if err := s.Sync(ctx); err != nil {
		log.Logger.Error("pool sync failed: " + err.Error())
	}
}

func (s *poolService) Sync(ctx context.Context) error {
	snapshot, err := s.Reader.Read(ctx)
	if err != nil {
		// Preserve the last successful snapshot but expose that the latest attempt failed.
		_ = db.Mysql.Model(&models.SyncState{}).Where("chain_id=? AND pool_contract=?", config.Config.TestNet.ChainId, config.Config.TestNet.PledgePoolToken).Updates(map[string]interface{}{"last_error": err.Error(), "complete": false}).Error
		return err
	}
	err = s.Persist(snapshot)
	if err != nil {
		_ = db.Mysql.Model(&models.SyncState{}).Where("chain_id=? AND pool_contract=?", snapshot.ChainID, snapshot.PoolContract).Updates(map[string]interface{}{"last_error": err.Error(), "complete": false}).Error
	}
	return err
}

// Persist publishes base data, distribution data and the sync watermark as one transaction.
// Redis is a derived cache; it is advanced only after the SQL commit succeeds.
func (s *poolService) Persist(snapshot *chain.Snapshot) error {
	now := time.Now().UTC().Format(time.RFC3339)
	err := db.Mysql.Transaction(func(tx *gorm.DB) error {
		for _, token := range snapshot.Tokens {
			row := models.TokenInfo{Token: token.Address, ChainId: snapshot.ChainID, Symbol: token.Symbol, Decimals: int(token.Decimals), Price: token.Price, CreatedAt: now, UpdatedAt: now}
			var existing models.TokenInfo
			result := tx.Where("chain_id=? AND token=?", snapshot.ChainID, token.Address).Limit(1).Find(&existing)
			if result.Error != nil {
				return result.Error
			}
			if result.RowsAffected == 0 {
				if err := tx.Create(&row).Error; err != nil {
					return err
				}
			} else {
				if err := tx.Model(&existing).Updates(map[string]interface{}{"symbol": row.Symbol, "decimals": row.Decimals, "price": row.Price, "updated_at": now}).Error; err != nil {
					return err
				}
			}
		}
		for _, pool := range snapshot.Pools {
			lend := snapshot.Tokens[pool.LendToken]
			borrow := snapshot.Tokens[pool.BorrowToken]
			lendJSON, _ := json.Marshal(models.LendToken{LendFee: snapshot.LendFee, TokenName: lend.Symbol, TokenPrice: lend.Price})
			borrowJSON, _ := json.Marshal(models.BorrowToken{BorrowFee: snapshot.BorrowFee, TokenName: borrow.Symbol, TokenPrice: borrow.Price})
			base := models.PoolBase{PoolContract: snapshot.PoolContract, SyncBlockNumber: snapshot.BlockNumber, SyncBlockHash: snapshot.BlockHash, PoolId: int(pool.DisplayPoolID), ChainId: snapshot.ChainID, SettleTime: pool.SettleTime, EndTime: pool.EndTime, InterestRate: pool.InterestRate, MaxSupply: pool.MaxSupply, LendSupply: pool.LendSupply, BorrowSupply: pool.BorrowSupply, MartgageRate: pool.MartgageRate, LendToken: pool.LendToken, BorrowToken: pool.BorrowToken, LendTokenSymbol: lend.Symbol, BorrowTokenSymbol: borrow.Symbol, LendTokenInfo: string(lendJSON), BorrowTokenInfo: string(borrowJSON), State: strconv.Itoa(int(pool.State)), SpCoin: pool.SP, JpCoin: pool.JP, AutoLiquidateThreshold: pool.AutoLiquidateThreshold, CreatedAt: now, UpdatedAt: now}
			data := models.PoolData{PoolContract: snapshot.PoolContract, SyncBlockNumber: snapshot.BlockNumber, SyncBlockHash: snapshot.BlockHash, PoolId: strconv.FormatUint(pool.DisplayPoolID, 10), ChainId: snapshot.ChainID, SettleAmountLend: pool.SettleAmountLend, SettleAmountBorrow: pool.SettleAmountBorrow, FinishAmountLend: pool.FinishAmountLend, FinishAmountBorrow: pool.FinishAmountBorrow, LiquidationAmounLend: pool.LiquidationAmountLend, LiquidationAmounBorrow: pool.LiquidationAmountBorrow, CreatedAt: now, UpdatedAt: now}
			if err := tx.Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "chain_id"}, {Name: "pool_contract"}, {Name: "pool_id"}}, UpdateAll: true}).Create(&base).Error; err != nil {
				return err
			}
			if err := tx.Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "chain_id"}, {Name: "pool_contract"}, {Name: "pool_id"}}, UpdateAll: true}).Create(&data).Error; err != nil {
				return err
			}
		}
		state := models.SyncState{ChainID: snapshot.ChainID, PoolContract: snapshot.PoolContract, BlockNumber: snapshot.BlockNumber, BlockHash: snapshot.BlockHash, PoolCount: len(snapshot.Pools), Complete: true, LastSuccess: now}
		return tx.Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "chain_id"}, {Name: "pool_contract"}}, UpdateAll: true}).Create(&state).Error
	})
	if err != nil {
		return fmt.Errorf("persist snapshot: %w", err)
	}
	if db.RedisConn != nil {
		key := "snapshot:" + snapshot.ChainID + ":" + snapshot.PoolContract
		if err := db.RedisSet(key, snapshot, 1800); err != nil {
			log.Logger.Warn("snapshot persisted; Redis refresh failed")
		}
	}
	return nil
}
