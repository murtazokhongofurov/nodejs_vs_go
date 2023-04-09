package main

import (
	"gitlab.com/nodejs_vs_go/go_api_gateway/api"

	"github.com/gin-gonic/gin"
	"gitlab.com/nodejs_vs_go/go_api_gateway/api/handlers"
	"gitlab.com/nodejs_vs_go/go_api_gateway/config"
	"gitlab.com/nodejs_vs_go/go_api_gateway/pkg/logger"
	"gitlab.com/nodejs_vs_go/go_api_gateway/service"
)


func main() {
	cfg := config.Load()

	grpcSvcs, err := service.NewGrpcClients(cfg)
	if err != nil {
		panic(err)
	}

	var loggerLevel = new(string)
	*loggerLevel = logger.LevelDebug

	log := logger.NewLogger("go_api_gateway", *loggerLevel)
	defer func() {
		err := logger.Cleanup(log) 
		if err != nil {
			return
		}
	}()

	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	h := handlers.NewHandler(cfg, log, grpcSvcs)

	api.SetUpAPI(r, h, cfg)
	
	if err := r.Run(cfg.HttpPort); err != nil {
		return
	}

	
}