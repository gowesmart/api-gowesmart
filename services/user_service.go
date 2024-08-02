package services

import (
	"html"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gowesmart/api-gowesmart/exceptions"
	"github.com/gowesmart/api-gowesmart/model/entity"
	"github.com/gowesmart/api-gowesmart/model/web/request"
	"github.com/gowesmart/api-gowesmart/model/web/response"
	"github.com/gowesmart/api-gowesmart/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (service *UserService) Register(c *gin.Context, userReq *request.RegisterRequest) (*response.RegisterResponse, error) {
	db, logger := utils.GetDBAndLogger(c)

	newUser := service.toUserEntity(userReq)

	hashedPassword, err := utils.HashPassword(userReq.Password)
	if err != nil {
		return nil, err
	}
	newUser.Password = string(hashedPassword)
	newUser.Username = html.EscapeString(strings.TrimSpace(userReq.Username))
	newUser.RoleID = 2 // USER

	err = db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(newUser).Error; err != nil {
			return exceptions.NewCustomError(http.StatusConflict, "Username or email already exists")
		}

		if err = tx.Create(&entity.Profile{
			UserID: newUser.ID,
		}).Error; err != nil {
			return exceptions.NewCustomError(http.StatusConflict, "Username or email already exists")
		}

		if err = tx.Model(&entity.User{}).
			Preload("Role", "id = ?", newUser.RoleID).
			Where("id = ?", newUser.ID).Take(&newUser).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	logger.Info("user registered successfully", zap.Uint("userID", newUser.ID))

	return service.toRegisterResponse(newUser), nil
}

func (service *UserService) Login(c *gin.Context, userReq *request.LoginRequest) (*response.LoginResponse, error) {
	db, _ := utils.GetDBAndLogger(c)

	loginUser := service.toUserEntity(userReq)

	err := db.Transaction(func(tx *gorm.DB) error {

		err := db.Model(&entity.User{}).Where("email = ?", loginUser.Email).Take(&loginUser).Error
		if err != nil {
			return exceptions.NewCustomError(http.StatusUnauthorized, "Email or password is incorrect")
		}

		err = db.Model(&entity.User{}).
			Preload("Role", "id = ?", loginUser.RoleID).
			Where("id = ?", loginUser.ID).Take(&loginUser).Error

		if err != nil {
			return exceptions.NewCustomError(http.StatusUnauthorized, err.Error())
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	err = utils.VerifyPassword(userReq.Password, loginUser.Password)

	if err != nil {
		return nil, exceptions.NewCustomError(http.StatusUnauthorized, "Email or password is incorrect")
	}

	token, err := utils.GenerateToken(loginUser.ID, loginUser.RoleID)
	if err != nil {
		return nil, err
	}

	return service.toLoginResponse(loginUser, token), nil
}

// func (service *UserService) UpdatePassword(c *gin.Context, userID uint, newPassword string) error {
// 	db, logger := utils.GetDBAndLogger(c)

// 	hashedPassword, err := utils.HashPassword(newPassword)
// 	if err != nil {
// 		return err
// 	}

// 	if err := db.Model(&entity.User{}).Where("id = ?", userID).Update("password", hashedPassword).Error; err != nil {
// 		return err
// 	}

// 	logger.Info("user password updated successfully", zap.Uint("userID", userID))

// 	return nil
// }

// func (service *UserService) GetUserProfile(c *gin.Context, userID uint) (*response.UserProfileResponse, error) {
// 	db, _ := utils.GetDBAndLogger(c)

// 	var responseUser response.UserProfileResponse

// 	if err := db.Model(&entity.User{}).
// 		Select("users.id, users.username, users.role, users.email, profiles.user_id, profiles.full_name, profiles.bio, profiles.age, profiles.gender").
// 		Joins("left join profiles on users.id = profiles.user_id").
// 		Where("users.id = ?", userID).
// 		Scan(&responseUser).Error; err != nil {
// 		return nil, err
// 	}

// 	if responseUser.ID == 0 {
// 		return nil, exceptions.NewCustomError(http.StatusNotFound, "User not found")
// 	}

// 	return &responseUser, nil
// }

// func (service *UserService) UpdateUserProfile(c *gin.Context, updatedUser *entity.User, userID uint) (*response.UpdateUserProfileResponse, error) {
// 	db, logger := utils.GetDBAndLogger(c)

// 	var responseUser response.UpdateUserProfileResponse

// 	err := db.Transaction(func(tx *gorm.DB) error {
// 		if err := tx.Model(&entity.User{ID: userID}).
// 			Omit("password").
// 			Updates(updatedUser).Error; err != nil {
// 			return err
// 		}

// 		if err := tx.Model(&entity.Profile{}).
// 			Where("user_id = ?", userID).
// 			Updates(updatedUser.Profile).Error; err != nil {
// 			return err
// 		}

// 		if err := tx.Model(&entity.User{}).
// 			Select("profiles.*, users.id, users.username, users.role, users.email").
// 			Joins("left join profiles on users.id = profiles.user_id").
// 			Take(&responseUser, "users.id = ?", userID).Error; err != nil {
// 			return err
// 		}
// 		return nil
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	logger.Info("success updating user profile", zap.Uint("userID", userID))

// 	return &responseUser, nil
// }

// func (service *UserService) DeleteUserProfile(c *gin.Context, userID uint) error {
// 	db, logger := utils.GetDBAndLogger(c)

// 	err := db.Transaction(func(tx *gorm.DB) error {
// 		// if err := tx.Where("user_id = ?", userID).Delete(&entity.Profile{}).Error; err != nil {
// 		// 	return err
// 		// }

// 		if err := tx.Delete(&entity.User{ID: userID}).Error; err != nil {
// 			return err
// 		}

// 		return nil
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	logger.Info("user profile and user deleted successfully", zap.Uint("userID", userID))
// 	return nil
// }

// func (service *UserService) ForgotPassword(c *gin.Context, username string, email string) (*response.ForgotPasswordResponse, error) {
// 	db, _ := utils.GetDBAndLogger(c)

// 	var user entity.User
// 	if err := db.Where("username = ?", username).Where("email = ?", email).First(&user).Error; err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return nil, exceptions.NewCustomError(http.StatusNotFound, "Username or email not found")
// 		}
// 		return nil, err
// 	}

// 	token, err := utils.GenerateResetPasswordToken(user.ID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &response.ForgotPasswordResponse{
// 		Token: token,
// 	}, nil
// }

// func (service *UserService) ResetPassword(c *gin.Context, token string, newPassword string) error {
// 	db, _ := utils.GetDBAndLogger(c)

// 	claims, err := utils.ParseResetToken(token)
// 	if err != nil {
// 		return exceptions.NewCustomError(http.StatusBadRequest, "Invalid or expired token")
// 	}

// 	var user entity.User
// 	if err := db.First(&user, claims.UserID).Error; err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return exceptions.NewCustomError(http.StatusNotFound, "User not found")
// 		}
// 		return err
// 	}

// 	hashedPassword, err := utils.HashPassword(newPassword)
// 	if err != nil {
// 		return err
// 	}

// 	user.Password = string(hashedPassword)
// 	if err := db.Save(&user).Error; err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (service *UserService) GetCurrentUser(c *gin.Context, userID uint) (*response.GetUserCurrentResponse, error) {
// 	db, _ := utils.GetDBAndLogger(c)

// 	var user entity.User
// 	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return nil, exceptions.NewCustomError(http.StatusNotFound, "User not found")
// 		}
// 		return nil, err
// 	}

// 	return &response.GetUserCurrentResponse{
// 		ID:       user.ID,
// 		Username: user.Username,
// 		Email:    user.Email,
// 		Role:     user.Role,
// 	}, nil
// }

func (*UserService) toUserEntity(req any) *entity.User {
	switch r := req.(type) {
	case *request.RegisterRequest:
		return &entity.User{
			Username: r.Username,
			Email:    r.Email,
			Password: r.Password,
		}
	case *request.LoginRequest:
		return &entity.User{
			Email:    r.Email,
			Password: r.Password,
		}
	default:
		return nil
	}
}

func (*UserService) toRegisterResponse(user *entity.User) *response.RegisterResponse {
	return &response.RegisterResponse{
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role.Name,
	}
}

func (*UserService) toLoginResponse(user *entity.User, token string) *response.LoginResponse {
	return &response.LoginResponse{
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role.Name,
		Token:    token,
	}
}
