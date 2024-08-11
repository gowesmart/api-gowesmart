package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gowesmart/api-gowesmart/exceptions"
	"github.com/gowesmart/api-gowesmart/model/web"
	"github.com/gowesmart/api-gowesmart/model/web/request"
	_ "github.com/gowesmart/api-gowesmart/model/web/response"
	"github.com/gowesmart/api-gowesmart/services"
	"github.com/gowesmart/api-gowesmart/utils"
)

type TransactionController struct {
	service services.TransactionService
}

func NewTransactionController(service services.TransactionService) TransactionController {
	return TransactionController{service: service}
}

// Register godoc
// @Summary Get all transaction.
// @Description Registering a user from public access.
// @Tags Transactions
// @Produce json
// @Param Authorization	header string true	"Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param limit query int false "Limit" default(10)
// @Param page query int false "Page" default(1)
// @Success 200 {object} web.WebSuccess[[]response.TransactionResponse]
// @Failure 400 {object} web.WebBadRequestError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/transactions [get]
func (t TransactionController) GetAll(c *gin.Context) {
	utils.UserRoleMustAdmin(c)

	var pagination web.PaginationRequest

	err := c.ShouldBindQuery(&pagination)
	utils.PanicIfError(err)

	res, metadata, err := t.service.GetAll(c, &pagination)
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, res, metadata)
}

// GetById godoc
// @Summary Get transaction by ID
// @Description Get transaction by ID
// @Tags Transactions
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 200 {object} web.WebSuccess[response.TransactionResponse]
// @Failure 400 {object} web.WebBadRequestError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/transactions/{id} [get]
func (t TransactionController) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.PanicIfError(exceptions.NewCustomError(http.StatusBadRequest, "id must be an integer"))
	}

	res, err := t.service.GetById(c, id)
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, res, nil)
}

// Create godoc
// @Summary Create a new transaction
// @Description Create a new transaction
// @Tags Transactions
// @Accept json
// @Produce json
// @Param Authorization	header string true	"Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param payload body []request.TransactionCreate true "Transaction payload"
// @Success 200 {object} web.WebSuccess[response.CreateTransactionResponse]
// @Failure 400 {object} web.WebBadRequestError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/transactions [post]
func (t TransactionController) Create(c *gin.Context) {
	claims, err := utils.ExtractTokenClaims(c)
	utils.PanicIfError(err)

	var payload []request.TransactionCreate
	err = c.ShouldBindJSON(&payload)
	utils.PanicIfError(err)

	data, err := t.service.Create(c, payload, int(claims.UserID))
	utils.PanicIfError(exceptions.NewCustomError(500, err.Error()))

	utils.ToResponseJSON(c, http.StatusOK, data, nil)
}

// Update godoc
// @Summary Update a transaction
// @Description Update a transaction
// @Tags Transactions
// @Accept json
// @Produce json
// @Param Authorization	header string true	"Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param id path int true "Transaction ID"
// @Param payload body []request.TransactionUpdate true "Transaction update payload"
// @Success 200 {object} web.WebSuccess[string]
// @Failure 400 {object} web.WebBadRequestError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/transactions/{id} [patch]
func (t TransactionController) Update(c *gin.Context) {
	claims, err := utils.ExtractTokenClaims(c)
	utils.PanicIfError(err)

	transactionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.PanicIfError(exceptions.NewCustomError(http.StatusBadRequest, "id must be an integer"))
	}

	var payload []request.TransactionUpdate
	err = c.ShouldBindJSON(&payload)
	utils.PanicIfError(err)

	err = t.service.Update(c, payload, transactionID, int(claims.UserID))
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, "data successfully updated", nil)
}

// Delete godoc
// @Summary Delete a transaction
// @Description Delete a transaction
// @Tags Transactions
// @Produce json
// @Param Authorization	header string true	"Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param id path int true "Transaction ID"
// @Success 200 {object} web.WebSuccess[string]
// @Failure 400 {object} web.WebBadRequestError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/transactions/{id} [delete]
func (t TransactionController) Delete(c *gin.Context) {
	claims, err := utils.ExtractTokenClaims(c)
	utils.PanicIfError(err)

	transactionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.PanicIfError(exceptions.NewCustomError(http.StatusBadRequest, "id must be an integer"))
	}

	err = t.service.Delete(c, transactionID, int(claims.UserID))
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, "data successfully deleted", nil)
}

// Pay godoc
// @Summary Pay for a transaction
// @Description Pay for a transaction
// @Tags Transactions
// @Produce json
// @Param Authorization	header string true	"Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param id path int true "Transaction ID"
// @Success 200 {object} web.WebSuccess[string]
// @Failure 400 {object} web.WebBadRequestError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/transactions/payment/{id} [patch]
func (t TransactionController) Pay(c *gin.Context) {
	claims, err := utils.ExtractTokenClaims(c)
	utils.PanicIfError(err)

	transactionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.PanicIfError(exceptions.NewCustomError(http.StatusBadRequest, "id must be an integer"))
	}

	err = t.service.Pay(c, transactionID, int(claims.UserID))
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, "payment success", nil)
}
