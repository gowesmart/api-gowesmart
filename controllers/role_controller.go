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

// RoleController handles role-related requests
type RoleController struct {
	roleService *services.RoleService
}

// NewRoleController creates a new RoleController
func NewRoleController(roleService *services.RoleService) *RoleController {
	return &RoleController{roleService: roleService}
}

// UpdateRoleByUserID godoc
// @Summary Update role for a specific user by ID
// @Description Update the role for a user based on user ID and role ID. Role 1 is for admin and role 2 is for user.
// @Tags Roles
// @Accept json
// @Produce json
// @Param Authorization	header string true	"Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param request body request.UpdateRoleRequest true "Update Role Request"
// @Success 200 {object} web.WebSuccess[response.RoleResponse]
// @Failure 400 {object} web.WebBadRequestError
// @Failure 404 {object} web.WebNotFoundError
// @Failure 500 {object} web.WebInternalServerError
// @Router /roles/update [patch]
func (controller *RoleController) UpdateRoleByUserID(c *gin.Context) {
	utils.UserRoleMustAdmin(c)

	var roleReq request.UpdateRoleRequest
	err := c.ShouldBindJSON(&roleReq)
	utils.PanicIfError(err)

	res, err := controller.roleService.UpdateRoleByUserID(c, &roleReq)
	utils.PanicIfError(err)

	c.JSON(http.StatusOK, res)
}
