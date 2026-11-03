package testutil

import (
	"github.com/alicebob/miniredis/v2"
	"github.com/gomodule/redigo/redis"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"pledge-backend/config"
	"pledge-backend/db"
	"testing"
)

func Database(t *testing.T) *miniredis.Miniredis {
	t.Helper()
	sql, err := gorm.Open(sqlite.Open("file:"+t.Name()+"?mode=memory&cache=shared"), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}, Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		t.Fatal(err)
	}
	server := miniredis.RunT(t)
	pool := &redis.Pool{Dial: func() (redis.Conn, error) { return redis.Dial("tcp", server.Addr()) }}
	oldSQL, oldRedis := db.Mysql, db.RedisConn
	oldConfig := *config.Config
	db.Mysql, db.RedisConn = sql, pool
	t.Cleanup(func() {
		_ = pool.Close()
		raw, _ := sql.DB()
		_ = raw.Close()
		db.Mysql, db.RedisConn = oldSQL, oldRedis
		*config.Config = oldConfig
	})
	return server
}
