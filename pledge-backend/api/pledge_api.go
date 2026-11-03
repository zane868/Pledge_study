package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"os"
	"os/signal"
	"pledge-backend/api/middlewares"
	"pledge-backend/api/models"
	"pledge-backend/api/routes"
	"pledge-backend/api/static"
	"pledge-backend/api/validate"
	"pledge-backend/config"
	"pledge-backend/db"
	"syscall"
	"time"
)

func main() {
	if err := config.ValidateAPI(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	db.InitMysql()
	db.InitRedis()
	if err := models.InitTable(); err != nil {
		fmt.Fprintln(os.Stderr, "schema migration failed")
		os.Exit(1)
	}
	validate.BindingValidator()
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	app.Use(gin.Logger(), gin.Recovery(), middlewares.Cors())
	app.StaticFS("/storage", http.FS(static.Files))
	routes.InitRoute(app)
	server := &http.Server{Addr: net.JoinHostPort(config.Config.Env.Host, config.Config.Env.Port), Handler: app, ReadHeaderTimeout: 5 * time.Second, ReadTimeout: 15 * time.Second, WriteTimeout: 120 * time.Second, IdleTimeout: 60 * time.Second}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	go func() {
		<-ctx.Done()
		shutdown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_ = server.Shutdown(shutdown)
	}()
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Fprintln(os.Stderr, "HTTP server failed")
		os.Exit(1)
	}
}
