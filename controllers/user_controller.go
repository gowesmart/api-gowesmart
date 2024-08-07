package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gowesmart/api-gowesmart/model/web"
	"github.com/gowesmart/api-gowesmart/model/web/request"
	_ "github.com/gowesmart/api-gowesmart/model/web/response"
	"github.com/gowesmart/api-gowesmart/services"
	"github.com/gowesmart/api-gowesmart/utils"
)

type UserController struct {
	userService        *services.UserService
	profileService     *services.ProfileService
	transactionService *services.TransactionService
	cartItemService    *services.CartItemService
}

func NewUserController(userService *services.UserService, profileService *services.ProfileService, transactionService *services.TransactionService, cartItemService *services.CartItemService) *UserController {
	return &UserController{
		userService,
		profileService,
		transactionService,
		cartItemService,
	}
}

// Register godoc
// @Summary User register.
// @Description	Registering a user from public access.
// @Tags Auth
// @Param Body body	request.RegisterRequest	true "the body to register a user"
// @Produce json
// @Success 201	{object} web.WebSuccess[response.RegisterResponse]
// @Failure 400	{object} web.WebBadRequestError
// @Failure 500	{object} web.WebInternalServerError
// @Router /api/auth/register [post]
func (controller *UserController) Register(c *gin.Context) {
	var registerReq request.RegisterRequest

	err := c.ShouldBindJSON(&registerReq)
	utils.PanicIfError(err)

	res, err := controller.userService.Register(c, &registerReq)
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusCreated, res, nil)
}

// LoginUser godoc
// @Summary User login.
// @Description Logging in to get jwt token to access admin or user api by roles.
// @Tags	Auth
// @Param Body	body request.LoginRequest	true "the body to login a user"
// @Produce json
// @Success 200	{object} web.WebSuccess[response.LoginResponse]
// @Failure 400	{object} web.WebBadRequestError
// @Failure 401	{object} web.WebUnauthorizedError
// @Failure 500	{object} web.WebInternalServerError
// @Router /api/auth/login [post]
func (controller *UserController) Login(c *gin.Context) {
	var loginReq request.LoginRequest

	err := c.ShouldBindJSON(&loginReq)
	utils.PanicIfError(err)

	res, err := controller.userService.Login(c, &loginReq)
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, res, nil)
}

// ForgotPassword godoc
// @Summary Forgot password.
// @Description Request forgot password.
// @Tags Auth
// @Param Body body	request.ForgotPasswordRequest	true	"the body to request forgot password"
// @Produce json
// @Success 200	{object} web.WebSuccess[response.ForgotPasswordResponse]
// @Failure 404	{object} web.WebNotFoundError
// @Failure 500	{object} web.WebInternalServerError
// @Router /api/auth/forgot-password [post]
func (controller *UserController) ForgotPassword(c *gin.Context) {
	var forgotPasswordReq request.ForgotPasswordRequest
	err := c.ShouldBindJSON(&forgotPasswordReq)
	utils.PanicIfError(err)

	res, err := controller.userService.ForgotPassword(c, &forgotPasswordReq)
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, res, nil)
}

// ResetPassword godoc
// @Summary Reset password.
// @Description Reset password.
// @Tags Auth
// @Param Authorization header string true	"Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param Body	body	request.ResetPasswordRequest	true	"the body to reset password"
// @Produce	json
// @Success 200	{object} web.WebSuccess[string]
// @Failure 400	{object} web.WebBadRequestError
// @Failure 404	{object} web.WebNotFoundError
// @Failure 500	{object} web.WebInternalServerError
// @Router /api/auth/reset-password [post]
func (controller *UserController) ResetPassword(c *gin.Context) {
	var resetPasswordReq request.ResetPasswordRequest
	err := c.ShouldBindJSON(&resetPasswordReq)
	utils.PanicIfError(err)

	err = controller.userService.ResetPassword(c, &resetPasswordReq)
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, "password updated", nil)
}

// GetCurrentUser godoc
// @Summary Get current user.
// @Description	Get current user.
// @Tags Users
// @Param Authorization	header	string	true	"Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200	{object} web.WebSuccess[response.GetUserCurrentResponse]
// @Failure 404	{object} web.WebNotFoundError
// @Failure 400	{object} web.WebBadRequestError
// @Failure 500	{object} web.WebInternalServerError
// @Router /api/users/current [get]
func (controller *UserController) GetCurrentUser(c *gin.Context) {
	claims, err := utils.ExtractTokenClaims(c)
	utils.PanicIfError(err)

	res, err := controller.userService.GetCurrentUser(c, uint(claims.UserID))
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, res, nil)
}

// UpdateUserProfile godoc
// @Summary		Update user profile.
// @Description	Update user profile.
// @Tags Users
// @Param Authorization	header string true	"Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param Body body	request.ProfileUpdateRequest true	"the body to reset password"
// @Produce	json
// @Success	200	{object}	web.WebSuccess[response.ProfileResponse]
// @Failure	404	{object}	web.WebNotFoundError
// @Failure	400	{object}	web.WebBadRequestError
// @Failure	401	{object}	web.WebUnauthorizedError
// @Failure	500	{object}	web.WebInternalServerError
// @Router /api/users/profile [patch]
func (controller *UserController) UpdateUserProfile(c *gin.Context) {
	var profileUpdateReq request.ProfileUpdateRequest

	err := c.ShouldBindJSON(&profileUpdateReq)
	utils.PanicIfError(err)

	claims, err := utils.ExtractTokenClaims(c)
	utils.PanicIfError(err)

	res, err := controller.profileService.UpdateProfile(c, &profileUpdateReq, uint(claims.UserID))
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, res, nil)
}

// Find user profile godoc
// @Summary		Find user profile.
// @Description	Find a user profile by username.
// @Tags Users
// @Param username path	string true	"username"
// @Produce json
// @Success 200	{object} web.WebSuccess[response.ProfileResponse]
// @Failure 404	{object} web.WebNotFoundError
// @Failure 500	{object} web.WebInternalServerError
// @Router /api/users/profile/{username} [get]
func (controller *UserController) FindProfileByUsername(c *gin.Context) {
	res, err := controller.profileService.FindProfileByUsername(c, c.Param("username"))
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, res, nil)
}

// Get user transactions godoc
// @Summary Find user transactions.
// @Description	Find a user transactions by username.
// @Tags Users
// @Produce json
// @Param Authorization	header string true	"Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200	{object} web.WebSuccess[response.UserTransactionResponse]
// @Failure 404	{object} web.WebNotFoundError
// @Failure 500	{object} web.WebInternalServerError
// @Router /api/users/current/transactions [get]
func (controller *UserController) FindUserTransaction(c *gin.Context) {
	claims, err := utils.ExtractTokenClaims(c)
	utils.PanicIfError(err)

	res, err := controller.transactionService.GetTransactionByUserID(c, claims.UserID)
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, res, nil)
}

// Get user carts godoc
// @Summary Find user carts.
// @Description	Find a user carts by username.
// @Tags Users
// @Produce json
// @Param Authorization	header string true	"Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200	{object} web.WebSuccess[response.CartResponse]
// @Failure 404	{object} web.WebNotFoundError
// @Failure 500	{object} web.WebInternalServerError
// @Router /api/users/current/carts [get]
func (controller *UserController) FindCart(c *gin.Context) {
	claims, err := utils.ExtractTokenClaims(c)
	utils.PanicIfError(err)

	res, err := controller.cartItemService.GetByUserID(c, claims.UserID)
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, res, nil)
}

// GetAllUsers godoc
// @Summary Get all users
// @Description	Get all users
// @Tags Users
// @Produce json
// @Param Authorization	header string true	"Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param limit query int false "Limit" default(10)
// @Param page query int false "Page" default(1)
// @Success 200	{object} web.WebSuccess[[]response.UserResponse]
// @Failure 500	{object} web.WebInternalServerError
// @Router /api/users [get]
func (controller *UserController) GetAllUsers(c *gin.Context) {
	utils.UserRoleMustAdmin(c)

	var pagination web.PaginationRequest

	err := c.ShouldBindQuery(&pagination)
	utils.PanicIfError(err)

	res, metadata, err := controller.userService.GetAllUsers(c, &pagination)
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, res, metadata)
}
