package models

type SyncState struct {
	ID           uint   `gorm:"primaryKey" json:"-"`
	ChainID      string `gorm:"column:chain_id;uniqueIndex:idx_sync_identity" json:"chainId"`
	PoolContract string `gorm:"column:pool_contract;uniqueIndex:idx_sync_identity;size:42" json:"poolContract"`
	BlockNumber  string `json:"blockNumber"`
	BlockHash    string `json:"blockHash"`
	PoolCount    int    `json:"poolCount"`
	Complete     bool   `json:"complete"`
	LastSuccess  string `json:"lastSuccess"`
	LastError    string `json:"lastError,omitempty"`
}

func (SyncState) TableName() string { return "sync_state" }
