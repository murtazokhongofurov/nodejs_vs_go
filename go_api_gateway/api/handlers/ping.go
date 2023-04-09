package handlers

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/nodejs_vs_go/go_api_gateway/api/http"
)

// Ping godoc
// @ID ping
// @Router /ping [GET]
// @Summary returns "pong" message
// @Description this returns "pong" messsage to show service is working
// @Accept json
// @Produce json
// @Success 200 {object} http.Response{data=string} "Response data"
// @Failure 500 {object} http.Response{}
func (h *Handler) Ping(c *gin.Context) {
	h.handleResponse(c, http.OK, "pong")
}