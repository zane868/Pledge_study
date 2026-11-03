package models

import (
	"pledge-backend/db"
	schedule "pledge-backend/schedule/models"
)

func InitTable() error {
	if err := schedule.InitTable(); err != nil {
		return err
	}
	return db.Mysql.AutoMigrate(&MultiSign{}, &TokenList{})
}
