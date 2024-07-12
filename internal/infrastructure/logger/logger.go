package logger

import (
	"os"

	"go.uber.org/zap"
)

var Log *zap.Logger

func init() {
	if err := InitLogger(); err != nil {
		panic(err)
	}
}

func InitLogger() error {
	var err error
	if os.Getenv("APP_ENV") == "development" {
		Log, err = zap.NewDevelopment()
	} else {
		Log, err = zap.NewProduction()
	}
	if err != nil {
		return err
	}
	zap.ReplaceGlobals(Log)
	return nil
}
