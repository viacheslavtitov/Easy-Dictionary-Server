package main

import (
	"os"

	"go.uber.org/zap"
)

func main() {
	//load environment configuration
	env := loadEnv(os.Args[1])
	if env == nil {
		os.Exit(1)
	}
	//init logger
	initLogger(*env)
	zap.L().Debug("Debug log")
	zap.L().Info("Info log")
}
