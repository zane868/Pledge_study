package middlewares

import (
	"net/http"
	"pledge-backend/config"
	"pledge-backend/db"
	"pledge-backend/schedule/models"
	"time"

	"github.com/gin-gonic/gin"
)

// RequireSnapshot prevents an unsynchronized, failed or stale index from looking like an empty market.
func RequireSnapshot() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var state models.SyncState
		if db.Mysql == nil {
			ctx.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"code": 503, "message": "sync unavailable"})
			return
		}
		result := db.Mysql.Where("chain_id=? AND pool_contract=?", config.Config.TestNet.ChainId, config.Config.TestNet.PledgePoolToken).Limit(1).Find(&state)
		updated, err := time.Parse(time.RFC3339, state.LastSuccess)
		if result.Error != nil || result.RowsAffected == 0 || !state.Complete || err != nil || time.Since(updated) > 5*time.Minute {
			ctx.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"code": 503, "message": "pool snapshot is unavailable or stale"})
			return
		}
		ctx.Next()
	}
}
