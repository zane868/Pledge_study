package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"pledge-backend/config"
	"pledge-backend/internal/chain"
	"time"
)

func main() {
	if config.Config.TestNet.NetUrl == "" {
		fmt.Fprintln(os.Stderr, "SEPOLIA_RPC_URL is required")
		os.Exit(1)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()
	value, err := (chain.Client{URL: config.Config.TestNet.NetUrl, Deployment: config.Config.Deployment, Confirmations: 12}).Read(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, "chain check failed:", err)
		os.Exit(1)
	}
	out := json.NewEncoder(os.Stdout)
	out.SetIndent("", "  ")
	if err = out.Encode(value); err != nil {
		os.Exit(1)
	}
}
