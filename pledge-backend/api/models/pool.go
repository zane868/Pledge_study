package models

import (
	"encoding/json"
	"errors"
	"pledge-backend/api/models/request"
	"pledge-backend/config"
	"pledge-backend/db"
	"pledge-backend/schedule/models"
)

type Pool struct {
	PoolContract           string   `json:"poolContract"`
	PoolIndex              int      `json:"poolIndex"`
	PoolID                 int      `json:"pool_id"`
	SettleTime             string   `json:"settleTime"`
	EndTime                string   `json:"endTime"`
	InterestRate           string   `json:"interestRate"`
	MaxSupply              string   `json:"maxSupply"`
	LendSupply             string   `json:"lendSupply"`
	BorrowSupply           string   `json:"borrowSupply"`
	MartgageRate           string   `json:"martgageRate"`
	LendToken              string   `json:"lendToken"`
	LendTokenSymbol        string   `json:"lend_token_symbol"`
	BorrowToken            string   `json:"borrowToken"`
	BorrowTokenSymbol      string   `json:"borrow_token_symbol"`
	State                  string   `json:"state"`
	SpCoin                 string   `json:"spCoin"`
	JpCoin                 string   `json:"jpCoin"`
	AutoLiquidateThreshold string   `json:"autoLiquidateThreshold"`
	Pooldata               PoolData `json:"pooldata"`
}

func NewPool() *Pool {
	return &Pool{}
}

func (p *Pool) Pagination(req *request.Search) (error, int64, []Pool) {
	var total int64
	pools := []Pool{}
	poolBase := []models.PoolBase{}

	query := db.Mysql.Table("poolbases").Where("chain_id=? AND pool_contract=?", req.ChainID, config.Config.TestNet.PledgePoolToken)
	if req.LendTokenSymbol != "" {
		query = query.Where("lend_token_symbol=?", req.LendTokenSymbol)
	}
	if req.State != "" {
		query = query.Where("state=?", req.State)
	}
	if err := query.Count(&total).Error; err != nil {
		return err, 0, nil
	}

	err := query.Order("pool_id desc").Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize).Find(&poolBase).Debug().Error
	if err != nil {
		return err, 0, nil
	}

	for _, b := range poolBase {
		poolData := PoolData{}
		err = db.Mysql.Table("pooldata").Where("chain_id=? AND pool_contract=? AND pool_id=?", req.ChainID, b.PoolContract, b.PoolId).First(&poolData).Debug().Error
		if err != nil {
			return err, 0, nil
		}
		if poolData.SyncBlockHash != b.SyncBlockHash {
			return errors.New("snapshot changed during query; retry"), 0, nil
		}
		var lendToken models.LendToken
		_ = json.Unmarshal([]byte(b.LendTokenInfo), &lendToken)
		var borrowToken models.BorrowToken
		_ = json.Unmarshal([]byte(b.BorrowTokenInfo), &borrowToken)
		pools = append(pools, Pool{
			PoolContract: b.PoolContract, PoolIndex: b.PoolId - 1,
			PoolID:       b.PoolId,
			SettleTime:   b.SettleTime,
			EndTime:      b.EndTime,
			InterestRate: b.InterestRate,
			MaxSupply:    b.MaxSupply,
			LendSupply:   b.LendSupply,
			BorrowSupply: b.BorrowSupply,
			MartgageRate: b.MartgageRate,
			LendToken:    b.LendToken, LendTokenSymbol: lendToken.TokenName,
			BorrowToken: b.BorrowToken, BorrowTokenSymbol: borrowToken.TokenName,
			State:                  b.State,
			SpCoin:                 b.SpCoin,
			JpCoin:                 b.JpCoin,
			AutoLiquidateThreshold: b.AutoLiquidateThreshold,
			Pooldata:               poolData,
		})
	}
	return nil, total, pools
}
