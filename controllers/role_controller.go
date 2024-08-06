package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gowesmart/api-gowesmart/model/web/request"
	"github.com/gowesmart/api-gowesmart/services"
	"net/http"
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
// @Param request body request.UpdateRoleRequest true "Update Role Request"
// @Success 200 {object} response.RoleResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /roles/update [put]
func (controller *RoleController) UpdateRoleByUserID(c *gin.Context) {
	var roleReq request.UpdateRoleRequest
	if err := c.ShouldBindJSON(&roleReq); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	res, err := controller.roleService.UpdateRoleByUserID(c, &roleReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// ErrorResponse defines the structure for error responses
type ErrorResponse struct {
	Error string `json:"error"`
}
