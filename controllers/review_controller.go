package controllers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gowesmart/api-gowesmart/model/web"
	"github.com/gowesmart/api-gowesmart/model/web/request"
	_ "github.com/gowesmart/api-gowesmart/model/web/response"
	"github.com/gowesmart/api-gowesmart/services"
	"github.com/gowesmart/api-gowesmart/utils"
	"net/http"
	"strconv"
	_ "strconv"
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
//
//	@Summary		Create a review
//	@Description	Create a new review
//	@Tags			Reviews
//	@Accept			json
//	@Produce		json
//	@Param			review		body		request.CreateReviewRequest	true	"Review body"
//	@Success		201		{object}	web.WebSuccess[response.ReviewResponse]
//	@Failure		400		{object}	web.WebBadRequestError
//	@Failure		500		{object}	web.WebInternalServerError
//	@Router			/api/reviews [post]
func (controller *ReviewController) CreateReview(c *gin.Context) {
	var reviewReq request.CreateReviewRequest

	err := c.ShouldBindJSON(&reviewReq)
	utils.PanicIfError(err)

	res, err := controller.reviewService.CreateReview(c, &reviewReq)
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusCreated, res, nil)
}

// UpdateReview godoc
//
//	@Summary		Update a review
//	@Description	Update an existing review
//	@Tags			Reviews
//	@Accept			json
//	@Produce		json
//	@Param			id		path		uint						true	"Review ID"
//	@Param			review		body		request.UpdateReviewRequest	true	"Review body"
//	@Success		200		{object}	web.WebSuccess[response.ReviewResponse]
//	@Failure		400		{object}	web.WebBadRequestError
//	@Failure		404		{object}	web.WebNotFoundError
//	@Failure		500		{object}	web.WebInternalServerError
//	@Router			/api/reviews/{id} [put]
func (controller *ReviewController) UpdateReview(c *gin.Context) {
	var reviewReq request.UpdateReviewRequest

	err := c.ShouldBindJSON(&reviewReq)
	utils.PanicIfError(err)

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	res, err := controller.reviewService.UpdateReview(c, uint(id), &reviewReq)
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, res, nil)
}

// DeleteReview godoc
//
//	@Summary		Delete a review
//	@Description	Delete a review by ID
//	@Tags			Reviews
//	@Produce		json
//	@Param			id		path		uint	true	"Review ID"
//	@Success		204		{object}	web.WebSuccess[response.ReviewResponse]
//	@Failure		404		{object}	web.WebNotFoundError
//	@Failure		500		{object}	web.WebInternalServerError
//	@Router			/api/reviews/{id} [delete]
func (controller *ReviewController) DeleteReview(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = controller.reviewService.DeleteReview(c, uint(id))
	utils.PanicIfError(err)

	c.Status(http.StatusNoContent)
}

// GetAllReviews godoc
//
//	@Summary		Get all reviews
//	@Description	Get all reviews
//	@Tags			Reviews
//	@Produce		json
//	@Success		200		{object}	web.WebSuccess[[]response.ReviewResponse]
//	@Failure		500		{object}	web.WebInternalServerError
//	@Router			/api/reviews [get]
func (controller *ReviewController) GetAllReviews(c *gin.Context) {
	res, err := controller.reviewService.GetAllReviews(c)
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, res, nil)
}

// GetReviewByID godoc
//
//	@Summary		Get a review by ID
//	@Description	Get a review by ID
//	@Tags			Reviews
//	@Produce		json
//	@Param			id		path		uint	true	"Review ID"
//	@Success		200		{object}	web.WebSuccess[response.ReviewResponse]
//	@Failure		404		{object}	web.WebNotFoundError
//	@Failure		500		{object}	web.WebInternalServerError
//	@Router			/api/reviews/{id} [get]
func (controller *ReviewController) GetReviewByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	res, err := controller.reviewService.GetReviewByID(c, uint(id))
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, res, nil)
}
