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
	database "easy-dictionary-server/db"
	internalenv "easy-dictionary-server/internalenv"

	"github.com/gin-gonic/gin"
	"github.com/tillberg/autorestart"
	"go.uber.org/zap"
)

func main() {
	//load environment configuration
	changeEnvChan := make(chan internalenv.Env)
	env := internalenv.LoadEnv(os.Args[1], changeEnvChan)
	if env == nil {
		zap.S().Panic("Config file didn't initialize. Server will stop")
		close(changeEnvChan)
		os.Exit(1)
	}
	go func() {
		newEnv := <-changeEnvChan
		zap.S().Info(newEnv.AppEnv + " new config is not equal previous. Server will restart")
		go autorestart.RestartViaExec() //work only on Unix systems
	}()
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	//init logger
	internalenv.InitLogger(env)
	zap.S().Debug("Debug log")
	zap.S().Info("Info log")
	//init database
	// database := database.Setup(env)
	database.Setup(env)
	//init http routers
	routeGin := gin.Default()
	zap.S().Info("Trying to start http server by address " + env.CombineServerAddress())
	server := &http.Server{
		Addr:         env.CombineServerAddress(),
		Handler:      routeGin,
		ReadTimeout:  time.Duration(env.TimeOut) * time.Second,
		WriteTimeout: time.Duration(env.TimeOut) * time.Second,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.S().Fatal("Server error")
			zap.Error(err)
		}
	}()
	routeGin.Use(middleware.RequestLogger())
	routeGin.Use(gin.Recovery())
	route.Setup(env.TimeOut, &routeGin.RouterGroup)
	zap.S().Info("Server started")
	<-done
	zap.S().Info("Server is stopping")
	close(changeEnvChan)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		zap.S().Fatal("Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
		zap.S().Info("timeout of 5 seconds.")
	}
	zap.S().Info("Server stopped")
}
