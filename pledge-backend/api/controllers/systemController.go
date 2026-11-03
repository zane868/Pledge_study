package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pledge-backend/config"
	"pledge-backend/db"
	syncmodels "pledge-backend/schedule/models"
	"time"
)

func Health(ctx *gin.Context) {
	if db.Mysql == nil || db.RedisConn == nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"ready": false})
		return
	}
	sql, err := db.Mysql.DB()
	if err != nil || sql.PingContext(ctx.Request.Context()) != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"ready": false, "database": "unavailable"})
		return
	}
	conn := db.RedisConn.Get()
	defer conn.Close()
	if _, err = conn.Do("PING"); err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"ready": false, "cache": "unavailable"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"ready": true, "chainId": config.Config.TestNet.ChainId})
}

func Deployment(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"code": 0, "message": "success", "data": config.Config.Deployment})
}

func SyncStatus(ctx *gin.Context) {
	var state syncmodels.SyncState
	result := db.Mysql.Where("chain_id=? AND pool_contract=?", config.Config.TestNet.ChainId, config.Config.TestNet.PledgePoolToken).Limit(1).Find(&state)
	if result.Error != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"code": 503, "message": "sync status unavailable"})
		return
	}
	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"code": 503, "message": "initial sync pending", "data": gin.H{"complete": false}})
		return
	}
	updated, err := time.Parse(time.RFC3339, state.LastSuccess)
	stale := err != nil || time.Since(updated) > 5*time.Minute
	ctx.JSON(http.StatusOK, gin.H{"code": 0, "message": "success", "data": state, "stale": stale})
}

func DisabledProjectPrice(ctx *gin.Context) {
	ctx.JSON(http.StatusServiceUnavailable, gin.H{"code": 503, "message": "project-token price feed is not enabled; use the configured pool oracle for lending prices"})
}
