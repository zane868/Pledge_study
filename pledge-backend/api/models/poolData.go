package models

import (
	"pledge-backend/config"
	"pledge-backend/db"
)

type PoolData struct {
	PoolContract           string `json:"poolContract" gorm:"column:pool_contract"`
	SyncBlockNumber        string `json:"syncBlockNumber" gorm:"column:sync_block_number"`
	SyncBlockHash          string `json:"syncBlockHash" gorm:"column:sync_block_hash"`
	Id                     int    `json:"-" gorm:"column:id;primaryKey;autoIncrement"`
	PoolID                 int    `json:"pool_id" gorm:"column:pool_id;"`
	ChainId                string `json:"chain_id" gorm:"column:chain_id"`
	FinishAmountBorrow     string `json:"finish_amount_borrow" gorm:"column:finish_amount_borrow"`
	FinishAmountLend       string `json:"finish_amount_lend" gorm:"column:finish_amount_lend"`
	LiquidationAmounBorrow string `json:"liquidation_amoun_borrow" gorm:"column:liquidation_amoun_borrow"`
	LiquidationAmounLend   string `json:"liquidation_amoun_lend" gorm:"column:liquidation_amoun_lend"`
	SettleAmountBorrow     string `json:"settle_amount_borrow" gorm:"column:settle_amount_borrow"`
	SettleAmountLend       string `json:"settle_amount_lend" gorm:"column:settle_amount_lend"`
	CreatedAt              string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt              string `json:"updated_at" gorm:"column:updated_at"`
}

type PoolDataInfoRes struct {
	PoolIndex     int      `json:"poolIndex"`
	DisplayPoolID int      `json:"displayPoolId"`
	Index         int      `json:"index"`
	PoolData      PoolData `json:"pool_data"`
}

func NewPoolData() *PoolData {
	return &PoolData{}
}

func (p *PoolData) TableName() string {
	return "pooldata"
}

func (p *PoolData) PoolDataInfo(chainId int, res *[]PoolDataInfoRes) error {
	var poolData []PoolData

	err := db.Mysql.Table("pooldata").Where("chain_id=? AND pool_contract=?", chainId, config.Config.TestNet.PledgePoolToken).Order("pool_id asc").Find(&poolData).Debug().Error
	if err != nil {
		return err
	}

	for _, v := range poolData {
		*res = append(*res, PoolDataInfoRes{
			Index:     v.PoolID - 1,
			PoolIndex: v.PoolID - 1, DisplayPoolID: v.PoolID,
			PoolData: v,
		})
	}
	return nil
}
