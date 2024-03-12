package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/tillberg/autorestart"
	internal "github.com/viacheslavtitov/easy-dictionary-server/internal"
	"go.uber.org/zap"
)

func main() {
	//load environment configuration
	changeEnvChan := make(chan internal.Env)
	env := internal.LoadEnv(os.Args[1], changeEnvChan)
	if env == nil {
		log.Default().Panic("Config file didn't initialize. Server will stop")
		close(changeEnvChan)
		os.Exit(1)
	}
	go func() {
		newEnv := <-changeEnvChan
		zap.L().Info(newEnv.AppEnv + " new config is not equal previous. Server will restart")
		go autorestart.RestartViaExec() //work only on Unix systems
	}()
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	//init logger
	internal.InitLogger(*env)
	zap.L().Debug("Debug log")
	zap.L().Info("Info log")
	zap.L().Info("Server started")
	<-done
	zap.L().Info("Server is stopping")
	close(changeEnvChan)
	zap.L().Info("Server stopped")
}
