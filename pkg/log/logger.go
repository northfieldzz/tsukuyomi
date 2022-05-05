package log

import (
	"go.uber.org/zap"
	"os"
)

var logger *zap.Logger

func Init() error {
	var err error
	if os.Getenv("GIN_MODE") == "release" {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		return err
	}
	return nil
}

func GetLogger() *zap.Logger {
	return logger
}
