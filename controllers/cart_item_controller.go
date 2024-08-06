package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/gowesmart/api-gowesmart/model/web"
	"github.com/gowesmart/api-gowesmart/model/web/request"
	_ "github.com/gowesmart/api-gowesmart/model/web/response"
	"github.com/gowesmart/api-gowesmart/services"
	"github.com/gowesmart/api-gowesmart/utils"
)

type CartController struct {
	service services.CartItemService
}

func NewCartController(service services.CartItemService) CartController {
	return CartController{service: service}
}

// Create godoc
// @Summary Create a new cart item
// @Description Create a new cart item
// @Tags Carts
// @Accept json
// @Produce json
// @Param	Authorization	header string	true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param cart body request.CartItemCreateRequest true "Cart Item Create"
// @Success 201 {object} web.WebSuccess[response.CartResponse]
// @Failure 400 {object} web.WebBadRequestError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/carts [post]
func (ctrl CartController) Create(c *gin.Context) {
	var req request.CartItemCreateRequest
	err := c.ShouldBindJSON(&req)
	utils.PanicIfError(err)

	claims, err := utils.ExtractTokenClaims(c)
	utils.PanicIfError(err)

	res, err := ctrl.service.Create(c, req, claims.UserID)
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, res, nil)
}

// Update godoc
// @Summary Update a cart
// @Description Update a cart
// @Tags Carts
// @Accept json
// @Produce json
// @Param	Authorization	header string	true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param cart body request.CartItemUpdateRequest true "Cart Item Update"
// @Success 200 {object} web.WebSuccess[response.CartResponse]
// @Failure 400 {object} web.WebBadRequestError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/carts [patch]
func (ctrl CartController) Update(c *gin.Context) {
	var req request.CartItemUpdateRequest
	err := c.ShouldBindJSON(&req)
	utils.PanicIfError(err)

	claims, err := utils.ExtractTokenClaims(c)
	utils.PanicIfError(err)

	res, err := ctrl.service.Update(c, req, claims.UserID)
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, res, nil)
}

// Delete godoc
// @Summary Delete a cart
// @Description Delete a cart
// @Tags Carts
// @Accept json
// @Produce json
// @Param Authorization	header string true	"Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param cart body request.CartItemDeleteRequest true "Cart Update"
// @Success 204
// @Failure 200 {object} web.WebSuccess[string]
// @Failure 400 {object} web.WebBadRequestError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/carts [delete]
func (ctrl CartController) Delete(c *gin.Context) {
	var req request.CartItemDeleteRequest
	err := c.ShouldBindJSON(&req)
	utils.PanicIfError(err)

	claims, err := utils.ExtractTokenClaims(c)
	utils.PanicIfError(err)

	err = ctrl.service.Delete(c, req.BikeID, claims.UserID)
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, "Cart item successfully deleted", nil)
}
