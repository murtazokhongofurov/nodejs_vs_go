package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
	"gitlab.com/nodejs_vs_go/go_api_gateway/api/http"
	"gitlab.com/nodejs_vs_go/go_api_gateway/api/models"

	"gitlab.com/nodejs_vs_go/go_api_gateway/genproto/driver_service"
)

// Create Driver info
// @ID create_driver
// @Router /v1/driver [POST]
// @Summary Create driver
// @Description Create driver info
// @Tags Driver
// @Accept json
// @Produce json
// @Param driver body models.DriverRequest true "CreateDriverRequestBody"
// @Success 201 {object} http.Response{data=models.DriverResponse} "Driver data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateDriver(c *gin.Context) {
	var body models.DriverRequest

	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return 
	}

	resp, err := h.service.DriverService().Create(context.Background(), &driver_service.DriverRequest{
		FirstName: body.FirstName,
		LastName: body.LastName,
		PhoneNumber: body.PhoneNumber,
		BirthDate: body.BirthDate,
	})

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}



// Get Driver info
// @ID get_driver
// @Router /v1/driver/{id} [DELETE]
// @Summary Get driver
// @Description Get driver info
// @Tags Driver
// @Accept json
// @Produce json
// @Param id path string true "driver_id"
// @Success 200 {object} http.Response{data=models.DriverResponse} "Driver data"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetDriver(c *gin.Context) {

	driverId := c.Param("id")

	resp, err := h.service.DriverService().GetById(context.Background(), &driver_service.DriverId{
		Id: driverId,
	})
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}



// Update Driver info
// @ID update_driver
// @Router /v1/driver [PUT]
// @Summary Update driver
// @Description Update driver info
// @Tags Driver
// @Accept json
// @Produce json
// @Param driver body models.DriverResponse true "UpdateDriverRequestBody"
// @Success 201 {object} http.Response{data=models.DriverResponse} "Driver data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateDriver(c *gin.Context) {
	var body models.DriverResponse
	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	res, err := h.service.DriverService().Update(context.Background(), &driver_service.DriverResponse{
		Id: body.Id,
		FirstName: body.FirstName,
		LastName: body.LastName,
		PhoneNumber: body.PhoneNumber,
		BirthDate: body.BirthDate,
	})
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, res)
}



// Delete Driver info
// @ID delete_driver
// @Router /v1/driver/{id} [GET]
// @Summary Get driver
// @Description Get driver info
// @Tags Driver
// @Accept json
// @Produce json
// @Param id path string true "driver_id"
// @Success 204 
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteDriver(c *gin.Context) {
	driverId := c.Param("id")

	res, err := h.service.DriverService().Delete(context.Background(), &driver_service.DriverId{
		Id: driverId,
	})

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error()) 
		return
	}
	h.handleResponse(c, http.NoContent, res)
}