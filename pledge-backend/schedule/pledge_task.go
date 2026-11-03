package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"pledge-backend/config"
	"pledge-backend/db"
	"pledge-backend/schedule/models"
	"pledge-backend/schedule/tasks"
	"syscall"
	"time"
)

func main() {
	if config.Config.TestNet.NetUrl == "" {
		fmt.Fprintln(os.Stderr, "SEPOLIA_RPC_URL is required")
		os.Exit(1)
	}
	db.InitMysql()
	db.InitRedis()
	if err := models.InitTable(); err != nil {
		fmt.Fprintln(os.Stderr, "schema migration failed")
		os.Exit(1)
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	tasks.Run(ctx, 120*time.Second)
}
