package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
	"gitlab.com/nodejs_vs_go/go_api_gateway/api/http"
	"gitlab.com/nodejs_vs_go/go_api_gateway/api/models"
	"gitlab.com/nodejs_vs_go/go_api_gateway/genproto/driver_service"
)

// Create Car info
// @ID create_car
// @Router /v1/driver/car [POST]
// @Summary Create Car
// @Description Create car info
// @Tags Driver Car
// @Accept json
// @Produce json
// @Param car body models.CarRequest true "CreateCarRequestBody"
// @Success 201 {object} http.Response{data=models.CarResponse} "Car data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateCar(c *gin.Context) {
	var body models.CarRequest

	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.service.CarService().Create(context.Background(), &driver_service.CarRequest{
		DriverId: body.DriverId,
		CarModel: body.CarModel,
		ManufactureYear: body.ManufactureYear,
		NumberPlate: body.NumberPlate,
		TechnicalLicence: body.TechnicalLicence,
		CarType: body.CarType,
		CarImage: body.CarImage,
		Color: body.Color,
	})
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}