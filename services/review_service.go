package services

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gowesmart/api-gowesmart/model/entity"
	"github.com/gowesmart/api-gowesmart/model/web/request"
	"github.com/gowesmart/api-gowesmart/model/web/response"
	"github.com/gowesmart/api-gowesmart/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ReviewService struct{}

func NewReviewService() *ReviewService {
	return &ReviewService{}
}

func validateRating(c *gin.Context, rating int) error {
	if rating < 1 || rating > 5 {
		err := errors.New("rating must be between 1 and 5")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}
	return nil
}

func (service *ReviewService) CreateReview(c *gin.Context, reviewReq *request.CreateReviewRequest) (*response.ReviewResponse, error) {
	db, logger := utils.GetDBAndLogger(c)

	var res response.ReviewResponse

	// Validate the rating
	if err := validateRating(c, reviewReq.Rating); err != nil {
		return nil, nil
	}

	review := entity.Review{
		Comment: reviewReq.Comment,
		Rating:  reviewReq.Rating,
		BikeID:  reviewReq.BikeID,
		UserID:  reviewReq.UserID,
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&review).Error; err != nil {
			return err
		}

		if err := tx.Model(&entity.Review{}).
			Select("id, comment, rating, created_at, updated_at, bike_id, user_id").
			Take(&res, review.ID).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	logger.Info("success creating review", zap.Uint("reviewID", review.ID))

	return &res, nil
}

func (service *ReviewService) UpdateReview(c *gin.Context, id uint, reviewReq *request.UpdateReviewRequest) (*response.ReviewResponse, error) {
	db, logger := utils.GetDBAndLogger(c)

	var res response.ReviewResponse
	var review entity.Review

	// Validate the rating
	if err := validateRating(c, reviewReq.Rating); err != nil {
		return nil, nil // Response already sent, return nil
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&review, id).Error; err != nil {
			return err
		}

		review.Comment = reviewReq.Comment
		review.Rating = reviewReq.Rating

		if err := tx.Save(&review).Error; err != nil {
			return err
		}

		if err := tx.Model(&entity.Review{}).
			Select("id, comment, rating, created_at, updated_at, bike_id, user_id").
			Take(&res, review.ID).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	logger.Info("success updating review", zap.Uint("reviewID", review.ID))

	return &res, nil
}

func (service *ReviewService) DeleteReview(c *gin.Context, id uint) error {
	db, logger := utils.GetDBAndLogger(c)

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&entity.Review{}, id).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	logger.Info("success deleting review", zap.Uint("reviewID", id))

	return nil
}

func (service *ReviewService) GetAllReviews(c *gin.Context) ([]response.ReviewResponse, error) {
	db, logger := utils.GetDBAndLogger(c)

	var reviews []response.ReviewResponse

	if err := db.Model(&entity.Review{}).
		Select("id, comment, rating, created_at, updated_at, bike_id, user_id").
		Find(&reviews).Error; err != nil {
		logger.Error("failed to fetch reviews", zap.Error(err))
		return nil, err
	}

	logger.Info("success fetching all reviews")

	return reviews, nil
}

func (service *ReviewService) GetReviewByID(c *gin.Context, id uint) (*response.ReviewResponse, error) {
	db, logger := utils.GetDBAndLogger(c)

	var res response.ReviewResponse

	if err := db.Model(&entity.Review{}).
		Select("id, comment, rating, created_at, updated_at, bike_id, user_id").
		Take(&res, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.Warn("review not found", zap.Uint("reviewID", id))
			return nil, err
		}

		logger.Error("failed to fetch review", zap.Error(err))
		return nil, err
	}

	logger.Info("success fetching review", zap.Uint("reviewID", id))

	return &res, nil
}
