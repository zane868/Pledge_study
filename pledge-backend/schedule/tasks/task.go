package tasks

import (
	"context"
	"pledge-backend/schedule/services"
	"time"
)

// Run is a read-only synchronizer. It never clears Redis or signs transactions.
func Run(ctx context.Context, interval time.Duration) {
	service := services.NewPool()
	sync := func() { service.UpdateAllPoolInfo() }
	sync()
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			sync()
		}
	}
}
