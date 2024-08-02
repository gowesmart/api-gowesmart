package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

	res, err := controller.UserService.Register(c, &registerReq)
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusCreated, res, nil)
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

	res, err := controller.UserService.Login(c, &loginReq)
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, res, nil)
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
	var forgotPasswordReq request.ForgotPasswordRequest
	err := c.ShouldBindJSON(&forgotPasswordReq)
	utils.PanicIfError(err)

	res, err := controller.UserService.ForgotPassword(c, &forgotPasswordReq)
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, res, nil)
}

// ResetPassword godoc
// @Summary Reset password.
// @Description Reset password.
// @Tags Auth
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param Body body request.ResetPasswordRequest true "the body to reset password"
// @Produce json
// @Success 200 {object} web.WebSuccess[string]
// @Failure 400 {object} web.WebBadRequestError
// @Failure 404 {object} web.WebNotFoundError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/auth/reset-password [post]
func (controller *UserController) ResetPassword(c *gin.Context) {
	var resetPasswordReq request.ResetPasswordRequest
	err := c.ShouldBindJSON(&resetPasswordReq)
	utils.PanicIfError(err)

	err = controller.UserService.ResetPassword(c, &resetPasswordReq)
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
	claims, err := utils.ExtractTokenClaims(c)
	utils.PanicIfError(err)

	userResponse, err := controller.UserService.GetCurrentUser(c, uint(claims.UserID))
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, userResponse, nil)
}
