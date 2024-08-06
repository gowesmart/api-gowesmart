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

type RoleService struct{}

func NewRoleService() *RoleService {
	return &RoleService{}
}

func (service *RoleService) UpdateRoleByUserID(c *gin.Context, roleReq *request.UpdateRoleRequest) (*response.RoleResponse, error) {
	db, logger := utils.GetDBAndLogger(c)

	var res response.RoleResponse
	var user entity.User
	var role entity.Role

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&user, roleReq.UserID).Error; err != nil {
			return err
		}

		if err := tx.First(&role, roleReq.Role).Error; err != nil {
			return err
		}

		user.RoleID = role.ID

		if err := tx.Save(&user).Error; err != nil {
			return err
		}

		res.ID = role.ID
		res.Role = role.ID
		res.UserID = user.ID

		return nil
	})

	if err != nil {
		return nil, err
	}

	logger.Info("success updating role for user", zap.Uint("userID", roleReq.UserID))

	return &res, nil
}
