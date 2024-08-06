package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gowesmart/api-gowesmart/exceptions"
	_ "github.com/gowesmart/api-gowesmart/model/web"
	"github.com/gowesmart/api-gowesmart/model/web/request"
	_ "github.com/gowesmart/api-gowesmart/model/web/response"
	"github.com/gowesmart/api-gowesmart/services"
	"github.com/gowesmart/api-gowesmart/utils"
)

type BikeController struct {
	bikeService services.BikeService
}

func NewBikeController(bikeService *services.BikeService) *BikeController {
	return &BikeController{
		*bikeService,
	}
}

// CreateBike godoc
// @Summary Create a bike
// @Description Create a new bike
// @Tags Bikes
// @Accept json
// @Produce json
// @Param Authorization	header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param bike body request.CreateBikeRequest true "Bike body"
// @Success 201	{object} web.WebSuccess[response.BikeResponse]
// @Failure 400	{object} web.WebBadRequestError
// @Failure 500	{object} web.WebInternalServerError
// @Router /api/bikes [post]
func (controller *BikeController) CreateBike(c *gin.Context) {
	utils.UserRoleMustAdmin(c)

	var bikeReq request.CreateBikeRequest
	err := c.ShouldBindJSON(&bikeReq)
	utils.PanicIfError(err)

	res, err := controller.bikeService.CreateBike(c, &bikeReq)
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusCreated, res, nil)
}

// UpdateBike godoc
// @Summary Update a bike
// @Description	Update an existing bike
// @Tags Bikes
// @Accept json
// @Produce json
// @Param Authorization	header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param id path uint true	"Bike ID"
// @Param bike body request.UpdateBikeRequest	true	"Bike body"
// @Success 200 {object} web.WebSuccess[response.BikeResponse]
// @Failure 400 {object} web.WebBadRequestError
// @Failure 404 {object} web.WebNotFoundError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/bikes/{id} [patch]
func (controller *BikeController) UpdateBike(c *gin.Context) {
	utils.UserRoleMustAdmin(c)

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.PanicIfError(exceptions.NewCustomError(http.StatusBadRequest, "id must be an integer"))
	}

	var bikeReq request.UpdateBikeRequest
	err = c.ShouldBindJSON(&bikeReq)
	utils.PanicIfError(err)

	res, err := controller.bikeService.UpdateBike(c, uint(id), &bikeReq)
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, res, nil)
}

// DeleteBike godoc
// @Summary Delete a bike
// @Description	Delete a bike by ID
// @Tags Bikes
// @Produce json
// @Param Authorization	header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param id path uint true "Bike ID"
// @Success 200 {object} web.WebSuccess[response.BikeResponse]
// @Failure 404 {object} web.WebNotFoundError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/bikes/{id} [delete]
func (controller *BikeController) DeleteBike(c *gin.Context) {
	utils.UserRoleMustAdmin(c)

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.PanicIfError(exceptions.NewCustomError(http.StatusBadRequest, "id must be an integer"))
	}

	err = controller.bikeService.DeleteBike(c, uint(id))
	utils.PanicIfError(err)

	c.Status(http.StatusNoContent)
}

// GetAllBikes godoc
// @Summary Get all bikes
// @Description	Get all bikes
// @Tags Bikes
// @Produce json
// @Success 200 {object} web.WebSuccess[[]response.BikeResponse]
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/bikes [get]
func (controller *BikeController) GetAllBikes(c *gin.Context) {
	res, err := controller.bikeService.GetAllBikes(c)
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, res, nil)
}

// GetBikeByID godoc
// @Summary Get a bike by ID
// @Description	Get a bike by ID
// @Tags Bikes
// @Produce json
// @Param id path uint true	"Bike ID"
// @Success 200 {object} web.WebSuccess[response.BikeResponse]
// @Failure 404 {object} web.WebNotFoundError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/bikes/{id} [get]
func (controller *BikeController) GetBikeByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.PanicIfError(exceptions.NewCustomError(http.StatusBadRequest, "id must be an integer"))
	}

	res, err := controller.bikeService.GetBikeByID(c, uint(id))
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, res, nil)
}
