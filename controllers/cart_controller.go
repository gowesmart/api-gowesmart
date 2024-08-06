package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gowesmart/api-gowesmart/model/web/request"
	"github.com/gowesmart/api-gowesmart/services"
	"github.com/gowesmart/api-gowesmart/utils"
)

type CartController struct {
	service services.CartService
}

func NewCartController(service services.CartService) CartController {
	return CartController{service: service}
}

// GetAll godoc
// @Summary Get all carts
// @Description Get all carts
// @Tags Cart
// @Accept json
// @Produce json
// @Success 200 {array} response.CartResponse
// @Failure 400 {object} web.WebBadRequestError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/carts [get]
func (ctrl CartController) GetAll(c *gin.Context) {
	res, err := ctrl.service.GetAll(c)
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, res, nil)
}

// GetById godoc
// @Summary Get cart by ID
// @Description Get cart by ID
// @Tags Cart
// @Accept json
// @Produce json
// @Param id path int true "Cart ID"
// @Success 200 {object} response.CartResponse
// @Failure 400 {object} web.WebBadRequestError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/carts/{id} [get]
func (ctrl CartController) GetById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	res, err := ctrl.service.GetById(c, id)
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, res, nil)
}

// Create godoc
// @Summary Create a new cart
// @Description Create a new cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param cart body request.CartCreateRequest true "Cart Create"
// @Success 201 {object} response.CartResponse
// @Failure 400 {object} web.WebBadRequestError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/carts [post]
func (ctrl CartController) Create(c *gin.Context) {
	var req request.CartCreateRequest
	err := c.ShouldBindJSON(&req)
	utils.PanicIfError(err)

	res, err := ctrl.service.Create(c, req)
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, res, nil)
}

// Update godoc
// @Summary Update a cart
// @Description Update a cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param id path int true "Cart ID"
// @Param cart body request.CartUpdateRequest true "Cart Update"
// @Success 200 {object} response.CartResponse
// @Failure 400 {object} web.WebBadRequestError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/carts/{id} [patch]
func (ctrl CartController) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req request.CartUpdateRequest
	err := c.ShouldBindJSON(&req)
	utils.PanicIfError(err)

	res, err := ctrl.service.Update(c, id, req)
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, res, nil)
}

// Delete godoc
// @Summary Delete a cart
// @Description Delete a cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param id path int true "Cart ID"
// @Success 204
// @Failure 400 {object} web.WebBadRequestError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/carts/{id} [delete]
func (ctrl CartController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := ctrl.service.Delete(c, id)
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, "Cart successfully deleted", nil)
}
