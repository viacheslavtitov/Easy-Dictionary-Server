package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	middleware "easy-dictionary-server/api/middleware"
	route "easy-dictionary-server/api/router"
	db "easy-dictionary-server/db"
	internalenv "easy-dictionary-server/internalenv"
	utils "easy-dictionary-server/internalenv/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @title EasyDictionary API
// @version 1.0
// @description This is the REST API for EasyDictionary app
// @host localhost:8080
// @BasePath /
func main() {
	//load environment configuration
	env := internalenv.LoadEnv()
	if env == nil {
		zap.S().Panic("Environment didnt initialize. Server will stop")
		os.Exit(1)
	}
	//init logger
	internalenv.InitLogger(env)
	zap.S().Debug("Debug log")
	zap.S().Info("Info log")
	//init database
	database := db.Setup(env)
	db.RunMigrations(database.SQLDB, utils.GetMigrationsDir(env.AppEnv))
	//init http routers
	routeGin := gin.Default()
	zap.S().Info("Trying to start http server by address " + env.CombineServerAddress())
	server := &http.Server{
		Addr:         env.CombineServerAddress(),
		Handler:      routeGin,
		ReadTimeout:  time.Duration(env.TimeOut) * time.Second,
		WriteTimeout: time.Duration(env.TimeOut) * time.Second,
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.S().Fatal("Server error")
			zap.Error(err)
		}
	}()
	routeGin.Use(middleware.RequestLogger())
	limiter := middleware.NewClientLimiter(5, 10) //max 5 requests per second, max 10 requests at the same time
	routeGin.Use(middleware.RateLimitMiddleware(limiter))
	routeGin.Use(gin.Recovery())
	route.Setup(env.TimeOut, &routeGin.RouterGroup, database, env)
	zap.S().Info("Server started")
	<-ctx.Done()
	zap.S().Info("Server is stopping...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		zap.S().Fatal("Server forced to shutdown", err)
	}
	zap.S().Info("Server is stopped")
}
