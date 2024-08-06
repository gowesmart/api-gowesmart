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
	newUser.RoleID = uint(entity.IDRoleUser) // USER

	err = db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(newUser).Error; err != nil {
			return exceptions.NewCustomError(http.StatusConflict, "Username or email already exists")
		}

		err = tx.Create(&entity.Profile{
			UserID: newUser.ID,
		}).Error
		if err != nil {
			return err
		}

		err = tx.Create(&entity.Cart{
			UserID: newUser.ID,
		}).Error
		if err != nil {
			return err
		}

		err = tx.Model(&entity.User{}).
			Preload("Role").
			Where("id = ?", newUser.ID).Take(&newUser).Error
		if err != nil {
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
		err := tx.Model(&entity.User{}).
			Preload("Role").
			Where("email = ?", loginUser.Email).Take(&loginUser).Error
		if err != nil {
			return exceptions.NewCustomError(http.StatusUnauthorized, "Email or password is incorrect")
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

func (service *UserService) ForgotPassword(c *gin.Context, userReq *request.ForgotPasswordRequest) (*response.ForgotPasswordResponse, error) {
	db, _ := utils.GetDBAndLogger(c)

	user := service.toUserEntity(userReq)

	err := db.Where("username = ?", user.Username).Where("email = ?", user.Email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exceptions.NewCustomError(http.StatusNotFound, "Username or email not found")
		}
		return nil, err
	}

	token, err := utils.GenerateResetPasswordToken(user.ID)
	if err != nil {
		return nil, err
	}

	return service.toForgotPasswordResponse(token), nil
}

func (service *UserService) ResetPassword(c *gin.Context, userReq *request.ResetPasswordRequest) error {
	db, _ := utils.GetDBAndLogger(c)

	claims, err := utils.ExtractTokenClaims(c)
	if err != nil {
		return exceptions.NewCustomError(http.StatusBadRequest, "Invalid or expired token")
	}

	var user entity.User

	err = db.First(&user, claims.UserID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return exceptions.NewCustomError(http.StatusNotFound, "User not found")
		}
		return err
	}

	hashedPassword, err := utils.HashPassword(userReq.NewPassword)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	err = db.Save(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (service *UserService) GetCurrentUser(c *gin.Context, userID uint) (*response.GetUserCurrentResponse, error) {
	db, _ := utils.GetDBAndLogger(c)

	var user entity.User

	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&entity.User{}).
			Preload("Role").
			Where("id = ?", userID).Take(&user).Error
		if err != nil {
			return exceptions.NewCustomError(http.StatusUnauthorized, err.Error())
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return service.toGetCurrentUserResponse(&user), nil
}

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
	case *request.ForgotPasswordRequest:
		return &entity.User{
			Username: r.Username,
			Email:    r.Email,
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

func (*UserService) toForgotPasswordResponse(token string) *response.ForgotPasswordResponse {
	return &response.ForgotPasswordResponse{
		ForgotPasswordToken: token,
	}
}

func (*UserService) toGetCurrentUserResponse(user *entity.User) *response.GetUserCurrentResponse {
	return &response.GetUserCurrentResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role.Name,
	}
}
