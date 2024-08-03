package services

import (
	"github.com/gin-gonic/gin"
	"github.com/gowesmart/api-gowesmart/model/entity"
	"github.com/gowesmart/api-gowesmart/model/web/request"
	"github.com/gowesmart/api-gowesmart/model/web/response"
	"github.com/gowesmart/api-gowesmart/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProfileService struct{}

func NewProfileService() *ProfileService {
	return &ProfileService{}
}

func (service *ProfileService) UpdateProfile(c *gin.Context, profileReq *request.ProfileUpdateRequest, userID uint) (*response.ProfileResponse, error) {
	db, logger := utils.GetDBAndLogger(c)

	profileWithUser := service.toProfileEntity(profileReq, userID)

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&entity.User{ID: userID}).
			Omit("password").
			Updates(profileWithUser.User).Error; err != nil {
			return err
		}

		if err := tx.Model(&entity.Profile{}).
			Where("user_id = ?", userID).
			Updates(profileWithUser).Error; err != nil {
			return err
		}

		if err := tx.Model(&entity.Profile{}).
			Select("profiles.*, users.id, users.username, users.role, users.email").
			Joins("left join users on users.id = profiles.user_id").
			Take(&profileWithUser, "users.id = ?", userID).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	logger.Info("success updating user profile", zap.Uint("userID", userID))

	return service.toProfileResponse(profileWithUser), nil
}

func (*ProfileService) toProfileEntity(req *request.ProfileUpdateRequest, userID uint) *entity.Profile {
	return &entity.Profile{
		Name:   req.Name,
		Bio:    req.Bio,
		Age:    req.Age,
		Gender: req.Gender,
		User: entity.User{
			ID:       userID,
			Username: req.Username,
			Email:    req.Email,
		},
	}
}

func (*ProfileService) toProfileResponse(profile *entity.Profile) *response.ProfileResponse {
	return &response.ProfileResponse{
		ID:       profile.ID,
		Username: profile.User.Username,
		Email:    profile.User.Email,
		Role:     profile.User.Role.Name,
		Name:     &profile.Name,
		Bio:      &profile.Bio,
		Age:      &profile.Age,
		Gender:   &profile.Gender,
	}
}
