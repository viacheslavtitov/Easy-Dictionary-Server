package utils

import (
	"flag"
	"os"

	"go.uber.org/zap"
)

func ReadFile(filepath string) (string, error) {
	zap.S().Debugf("Trying read file %s", filepath)
	fileContent, err := os.ReadFile(filepath)
	if err != nil {
		zap.S().Error(err)
		zap.S().Errorf("Failed to read file %s", filepath)
		return "", err
	}

	// Convert []byte to string
	text := string(fileContent)
	// zap.S().Debug(text)
	return text, nil
}

func GetMigrationsDir() string {
	if IsRunningTest() {
		return "file://../../migrations"
	} else {
		return "file://../db/migrations"
	}
}

func IsRunningTest() bool {
	return flag.Lookup("test.v") != nil
}
