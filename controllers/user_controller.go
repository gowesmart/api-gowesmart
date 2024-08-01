package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gowesmart/api-gowesmart/exceptions"
	"github.com/gowesmart/api-gowesmart/model/entity"
	_ "github.com/gowesmart/api-gowesmart/model/web"
	"github.com/gowesmart/api-gowesmart/model/web/request"
	_ "github.com/gowesmart/api-gowesmart/model/web/response"
	"github.com/gowesmart/api-gowesmart/services"
	"github.com/gowesmart/api-gowesmart/utils"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		*userService,
	}
}

// Register godoc
// @Summary User register.
// @Description Registering a user from public access.
// @Tags Auth
// @Param Body body request.RegisterRequest true "the body to register a user"
// @Produce json
// @Success 201 {object} web.WebSuccess[response.RegisterResponse]
// @Failure 400 {object} web.WebBadRequestError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/auth/register [post]
func (controller *UserController) Register(c *gin.Context) {
	var registerReq request.RegisterRequest

	err := c.ShouldBindJSON(&registerReq)
	utils.PanicIfError(err)

	newUser := entity.User{
		Username: registerReq.Username,
		Email:    registerReq.Email,
		Password: registerReq.Password,
	}

	registerRes, err := controller.UserService.Register(c, &newUser)
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusCreated, registerRes, nil)
}

// LoginUser godoc
// @Summary User login.
// @Description Logging in to get jwt token to access admin or user api by roles.
// @Tags Auth
// @Param Body body request.LoginRequest true "the body to login a user"
// @Produce json
// @Success 200 {object} web.WebSuccess[response.LoginResponse]
// @Failure 400 {object} web.WebBadRequestError
// @Failure 401 {object} web.WebUnauthorizedError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/auth/login [post]
func (controller *UserController) Login(c *gin.Context) {
	var loginReq request.LoginRequest

	err := c.ShouldBindJSON(&loginReq)
	utils.PanicIfError(err)

	loginRes, err := controller.UserService.Login(c, loginReq.Email, loginReq.Password)
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, loginRes, nil)
}

// UpdatePassword godoc
// @Summary Update user password.
// @Description Update the current user's password.
// @Tags Users
// @Param Body body request.UpdatePasswordRequest true "the body to update a password"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} web.WebSuccess[string]
// @Failure 400 {object} web.WebBadRequestError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/users/password [patch]
func (controller *UserController) UpdatePassword(c *gin.Context) {
	var updatePasswordReq request.UpdatePasswordRequest

	err := c.ShouldBindJSON(&updatePasswordReq)
	utils.PanicIfError(err)

	userID, _, err := utils.ExtractTokenClaims(c)
	utils.PanicIfError(err)

	err = controller.UserService.UpdatePassword(c, userID, updatePasswordReq.Password)
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, "password updated", nil)
}

// GetUserProfile godoc
// @Summary Get user profile.
// @Description Get user profile data.
// @Tags Users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} web.WebSuccess[response.UserProfileResponse]
// @Failure 400 {object} web.WebBadRequestError
// @Failure 404 {object} web.WebNotFoundError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/users/profile/{id} [get]
func (controller *UserController) GetUserProfile(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		panic(exceptions.NewCustomError(http.StatusBadRequest, "Id must be an integer"))
	}

	user, err := controller.UserService.GetUserProfile(c, uint(userID))
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, user, nil)
}

// UpdateUserProfile godoc
// @Summary Update user profile.
// @Description Update the profile of a user.
// @Tags Users
// @Param Body body request.UpdateUserProfileRequest true "the body to update user profile"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} web.WebSuccess[response.UpdateUserProfileResponse]
// @Failure 400 {object} web.WebBadRequestError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/users/profile [patch]
func (controller *UserController) UpdateUserProfile(c *gin.Context) {
	var updateUserReq request.UpdateUserProfileRequest

	err := c.ShouldBindJSON(&updateUserReq)
	utils.PanicIfError(err)

	userID, _, err := utils.ExtractTokenClaims(c)
	utils.PanicIfError(err)

	updatedProfile := &entity.User{
		Username: updateUserReq.Username,
		Email:    updateUserReq.Email,

		ID: uint(userID),
	}

	userResponse, err := controller.UserService.UpdateUserProfile(c, updatedProfile, userID)
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, userResponse, nil)
}

// DeleteUserProfile godoc
// @Summary Delete user.
// @Description Delete a user profile by ID.
// @Tags Users
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} web.WebSuccess[string]
// @Failure 400 {object} web.WebBadRequestError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/users [delete]
func (controller *UserController) DeleteUserProfile(c *gin.Context) {
	userID, _, err := utils.ExtractTokenClaims(c)
	utils.PanicIfError(err)

	err = controller.UserService.DeleteUserProfile(c, userID)
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, "user deleted", nil)
}

// ForgotPassword godoc
// @Summary Forgot password.
// @Description Request forgot password.
// @Tags Auth
// @Param Body body request.ForgotPasswordRequest true "the body to request forgot password"
// @Produce json
// @Success 200 {object} web.WebSuccess[response.ForgotPasswordResponse]
// @Failure 404 {object} web.WebNotFoundError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/auth/forgot-password [post]
func (controller *UserController) ForgotPassword(c *gin.Context) {
	var req request.ForgotPasswordRequest
	err := c.ShouldBindJSON(&req)
	utils.PanicIfError(err)

	userResponse, err := controller.UserService.ForgotPassword(c, req.Username, req.Email)
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, userResponse, nil)
}

// ResetPassword godoc
// @Summary Reset password.
// @Description Reset password.
// @Tags Auth
// @Param Body body request.ResetPasswordRequest true "the body to reset password"
// @Produce json
// @Success 200 {object} web.WebSuccess[string]
// @Failure 400 {object} web.WebBadRequestError
// @Failure 404 {object} web.WebNotFoundError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/auth/reset-password [post]
func (controller *UserController) ResetPassword(c *gin.Context) {
	var req request.ResetPasswordRequest
	err := c.ShouldBindJSON(&req)
	utils.PanicIfError(err)

	err = controller.UserService.ResetPassword(c, req.Token, req.NewPassword)
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, "password updated", nil)
}

// GetCurrentUser godoc
// @Summary Get current user.
// @Description Get current user.
// @Tags Users
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} web.WebSuccess[response.GetUserCurrentResponse]
// @Failure 404 {object} web.WebNotFoundError
// @Failure 400 {object} web.WebBadRequestError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/users/current [get]
func (controller *UserController) GetCurrentUser(c *gin.Context) {
	userID, _, err := utils.ExtractTokenClaims(c)
	utils.PanicIfError(err)

	userResponse, err := controller.UserService.GetCurrentUser(c, uint(userID))
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, userResponse, nil)
}
