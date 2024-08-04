package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gowesmart/api-gowesmart/exceptions"
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

	var res response.ProfileResponse

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
			Select("user_id as id, name, bio, age, gender, users.username, users.email").
			Joins("JOIN users ON profiles.user_id = users.id").
			Take(&res, "profiles.user_id = ?", userID).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	logger.Info("success updating user profile", zap.Uint("userID", userID))

	return &res, nil
}

func (service *ProfileService) FindProfileByUsername(c *gin.Context, username string) (*response.ProfileResponse, error) {
	db, _ := utils.GetDBAndLogger(c)

	var res response.ProfileResponse

	err := db.Model(&entity.Profile{}).
		Select("user_id as id, name, bio, age, gender, users.username, users.email").
		Joins("JOIN users ON profiles.user_id = users.id").
		Take(&res, "users.username = ?", username).Error

	if err != nil {
		return nil, exceptions.NewCustomError(http.StatusNotFound, "User not found")
	}

	return &res, nil
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
