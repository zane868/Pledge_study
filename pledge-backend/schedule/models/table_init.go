package models

import "pledge-backend/db"

func InitTable() error {
	return db.Mysql.AutoMigrate(&PoolBase{}, &PoolData{}, &RedisTokenInfo{}, &TokenInfo{}, &SyncState{})
}
