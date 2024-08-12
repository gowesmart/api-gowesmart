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

type ReviewController struct {
	reviewService services.ReviewService
}

func NewReviewController(reviewService *services.ReviewService) *ReviewController {
	return &ReviewController{
		*reviewService,
	}
}

// CreateReview godoc
// @Summary Create a review
// @Description	Create a new review
// @Tags Reviews
// @Accept json
// @Produce json
// @Param Authorization	header string true	"Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param review body request.CreateReviewRequest true	"Review body"
// @Success 201 {object} web.WebSuccess[response.ReviewResponse]
// @Failure 400	{object} web.WebBadRequestError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/reviews [post]
func (controller *ReviewController) CreateReview(c *gin.Context) {
	claims, err := utils.ExtractTokenClaims(c)
	utils.PanicIfError(err)

	var reviewReq request.CreateReviewRequest
	err = c.ShouldBindJSON(&reviewReq)
	utils.PanicIfError(err)

	res, err := controller.reviewService.CreateReview(c, &reviewReq, claims.UserID)
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusCreated, res, nil)
}

// UpdateReview godoc
// @Summary Update a review
// @Description	Update an existing review
// @Tags Reviews
// @Accept json
// @Produce json
// @Param Authorization	header string true	"Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param id path uint true	"Review ID"
// @Param review body request.UpdateReviewRequest	true	"Review body"
// @Success 200	{object} web.WebSuccess[response.ReviewResponse]
// @Failure 400	{object} web.WebBadRequestError
// @Failure 404	{object} web.WebNotFoundError
// @Failure 500	{object} web.WebInternalServerError
// @Router /api/reviews/{id} [patch]
func (controller *ReviewController) UpdateReview(c *gin.Context) {
	claims, err := utils.ExtractTokenClaims(c)
	utils.PanicIfError(err)

	var reviewReq request.UpdateReviewRequest
	err = c.ShouldBindJSON(&reviewReq)
	utils.PanicIfError(err)

	reviewID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.PanicIfError(exceptions.NewCustomError(http.StatusBadRequest, "id must be an integer"))
	}

	res, err := controller.reviewService.UpdateReview(c, &reviewReq, uint(reviewID), claims.UserID)
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, res, nil)
}

// DeleteReview godoc
// @Summary Delete a review
// @Description	Delete a review by ID
// @Tags Reviews
// @Produce json
// @Param Authorization	header string true	"Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param id path uint true	"Review ID"
// @Success 204	{object} web.WebSuccess[response.ReviewResponse]
// @Failure 404	{object} web.WebNotFoundError
// @Failure 500	{object} web.WebInternalServerError
// @Router /api/reviews/{id} [delete]
func (controller *ReviewController) DeleteReview(c *gin.Context) {
	_, err := utils.ExtractTokenClaims(c)
	utils.PanicIfError(err)

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.PanicIfError(exceptions.NewCustomError(http.StatusBadRequest, "id must be an integer"))
	}

	err = controller.reviewService.DeleteReview(c, uint(id))
	utils.PanicIfError(err)

	c.Status(http.StatusNoContent)
}

// GetAllReviews godoc
// @Summary Get all reviews
// @Description	Get all reviews
// @Tags Reviews
// @Produce json
// @Param Authorization	header string true	"Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param limit query int false "Limit" default(10)
// @Param page query int false "Page" default(1)
// @Success 200	{object} web.WebSuccess[[]response.ReviewResponse]
// @Failure 500	{object} web.WebInternalServerError
// @Router /api/reviews [get]
func (controller *ReviewController) GetAllReviews(c *gin.Context) {
	utils.UserRoleMustAdmin(c)

	var pagination web.PaginationRequest

	err := c.ShouldBindQuery(&pagination)
	utils.PanicIfError(err)

	res, metadata, err := controller.reviewService.GetAllReviews(c, &pagination)
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, res, metadata)
}

// GetReviewByID godoc
// @Summary Get a review by ID
// @Description	Get a review by ID
// @Tags Reviews
// @Produce json
// @Param Authorization	header string true	"Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param id path uint true	"Review ID"
// @Success 200	{object} web.WebSuccess[response.ReviewResponse]
// @Failure 404	{object} web.WebNotFoundError
// @Failure 500	{object} web.WebInternalServerError
// @Router /api/reviews/{id} [get]
func (controller *ReviewController) GetReviewByID(c *gin.Context) {
	utils.UserRoleMustAdmin(c)

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.PanicIfError(exceptions.NewCustomError(http.StatusBadRequest, "id must be an integer"))
	}

	res, err := controller.reviewService.GetReviewByID(c, uint(id))
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, res, nil)
}

// GetReviewByOrderID godoc
// @Summary Get a review by ID
// @Description	Get a review by ID
// @Tags Reviews
// @Produce json
// @Param Authorization	header string true	"Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param Id path uint true	"Order ID"
// @Success 200	{object} web.WebSuccess[response.ReviewResponse]
// @Failure 404	{object} web.WebNotFoundError
// @Failure 500	{object} web.WebInternalServerError
// @Router /api/reviews/order/{orderId} [get]
func (controller *ReviewController) GetReviewByOrderID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.PanicIfError(exceptions.NewCustomError(http.StatusBadRequest, "id must be an integer"))
	}

	res, err := controller.reviewService.GetReviewByOrderID(c, uint(id))
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, res, nil)
}
