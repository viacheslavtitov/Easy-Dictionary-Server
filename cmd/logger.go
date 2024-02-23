package main

import (
	"go.uber.org/zap"
)

func initLogger(env Env) {
	var config zap.Config
	switch env.AppEnv {
	case "local":
		{
			config = zap.NewDevelopmentConfig()
		}
	case "development":
		{
			config = zap.NewDevelopmentConfig()
		}
	case "production":
		{
			config = zap.NewProductionConfig()
		}
	}
	logger, _ := config.Build()
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
}
