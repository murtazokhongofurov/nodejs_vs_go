package api

import (
	"github.com/gin-gonic/gin"
	swaggerfile "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "gitlab.com/nodejs_vs_go/go_api_gateway/api/docs"
	"gitlab.com/nodejs_vs_go/go_api_gateway/api/handlers"
	"gitlab.com/nodejs_vs_go/go_api_gateway/config"
)

// @description This is a api-gateway
// @termsOfService https://murtazoxon.uz
func SetUpAPI(r *gin.Engine, h handlers.Handler, cfg config.Config) {
	r.GET("/ping", )
	api := r.Group("/v1", h.Ping)

	api.POST("/driver", h.CreateDriver)
	api.GET("/driver/:id", h.GetDriver)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfile.Handler))
}