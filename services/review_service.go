package services

import (
	"github.com/gin-gonic/gin"
	"github.com/gowesmart/api-gowesmart/model/entity"
	"github.com/gowesmart/api-gowesmart/model/web"
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

func (service *ReviewService) CreateReview(c *gin.Context, reviewReq *request.CreateReviewRequest, userID uint) (*response.ReviewResponse, error) {
	db, logger := utils.GetDBAndLogger(c)

	var res response.ReviewResponse

	review := entity.Review{
		Comment: reviewReq.Comment,
		Rating:  reviewReq.Rating,
		BikeID:  reviewReq.BikeID,
		UserID:  userID,
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

func (service *ReviewService) UpdateReview(c *gin.Context, reviewReq *request.UpdateReviewRequest, reviewID, userID uint) (*response.ReviewResponse, error) {
	db, logger := utils.GetDBAndLogger(c)

	var res response.ReviewResponse
	var review entity.Review

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id = ?", userID).First(&review, reviewID).Error; err != nil {
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

func (service *ReviewService) GetAllReviews(c *gin.Context, pagination *web.PaginationRequest) ([]response.GetAllReviewResponse, *web.Metadata, error) {
	db, logger := utils.GetDBAndLogger(c)

	var reviews []response.GetAllReviewResponse

	query := db.Model(&entity.Review{}).
		Select("reviews.*, bikes.name as bike_name, users.username as user_username").
		Joins("JOIN bikes ON reviews.bike_id = bikes.id").
		Joins("JOIN users ON reviews.user_id = users.id")

	var totalData int64
	if err := query.Count(&totalData).Error; err != nil {
		logger.Error("failed to count reviews", zap.Error(err))
		return nil, nil, err
	}
	pagination.TotalData = totalData

	offset := pagination.GetOffset()
	limit := pagination.GetLimit()
	if err := query.Offset(offset).Limit(limit).Find(&reviews).Error; err != nil {
		logger.Error("failed to fetch reviews", zap.Error(err))
		return nil, nil, err
	}

	pagination.TotalPages = int((totalData + int64(limit) - 1) / int64(limit))

	metadata := &web.Metadata{
		Page:       &pagination.Page,
		Limit:      &pagination.Limit,
		TotalPages: &pagination.TotalPages,
		TotalData:  &pagination.TotalData,
	}

	logger.Info("success fetching all reviews", zap.Int("total_data", int(totalData)), zap.Int("total_pages", pagination.TotalPages))

	return reviews, metadata, nil
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

func (service *ReviewService) GetReviewByBikeID(c *gin.Context, bikeID uint) ([]response.ReviewResponse, error) {
	db, logger := utils.GetDBAndLogger(c)

	var reviews []response.ReviewResponse

	if err := db.Model(&entity.Review{}).
		Where("bike_id = ?", bikeID).
		Select("id, comment, rating, created_at, updated_at, bike_id, user_id").
		Find(&reviews).Error; err != nil {
		logger.Error("failed to fetch reviews by bike id", zap.Error(err))
		return nil, err
	}

	logger.Info("success fetching reviews by bike id")

	return reviews, nil
}
