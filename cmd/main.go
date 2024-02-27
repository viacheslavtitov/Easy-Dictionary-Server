package main

import (
	"os"

	internal "github.com/viacheslavtitov/easy-dictionary-server/internal"
	"go.uber.org/zap"
)

func main() {
	//load environment configuration
	env := internal.LoadEnv(os.Args[1])
	if env == nil {
		os.Exit(1)
	}
	//init logger
	internal.InitLogger(*env)
	zap.L().Debug("Debug log")
	zap.L().Info("Info log")
}
