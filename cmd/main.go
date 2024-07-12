package main

import (
	"github.com/mpt-tiendc/golang-template/internal/config"
	"github.com/mpt-tiendc/golang-template/internal/infrastructure/logger"
	"github.com/mpt-tiendc/golang-template/internal/interfaces/api/http"
	"go.uber.org/zap"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	httpServer := http.NewHttpServer(config)

	logger.Log.Info("Starting http server ....", zap.String("Host", config.Host), zap.Int("Port", config.Port))
	logger.Log.Error(httpServer.Start().Error())
}
