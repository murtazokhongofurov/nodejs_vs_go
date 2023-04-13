package handlers

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/nodejs_vs_go/go_api_gateway/api/http"
	"gitlab.com/nodejs_vs_go/go_api_gateway/config"
	"gitlab.com/nodejs_vs_go/go_api_gateway/pkg/logger"
	"gitlab.com/nodejs_vs_go/go_api_gateway/service"
)
type Handler struct {
	cfg config.Config
	log logger.LoggerI
	service service.ServiceManagerI
}

func NewHandler(cfg config.Config, log logger.LoggerI, srvs service.ServiceManagerI) Handler {
	return Handler{
		cfg: cfg,
		log: log,
		service: srvs,
	}
}


func (h *Handler) handleResponse(c *gin.Context, status http.Status, data interface{}) {
	switch code := status.Code; {
	case code < 300:
		h.log.Info(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	case code < 400:
		h.log.Warn(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	default:
		h.log.Error(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	}

	c.JSON(status.Code, http.Response{
		Status:      status.Status,
		Description: status.Description,
		Data:        data,
	})
}