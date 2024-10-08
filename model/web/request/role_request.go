package request

type UpdateRoleRequest struct {
	UserID uint `json:"user_id" binding:"required"`
	Role   uint `json:"role" binding:"required,min=1,max=2"`
}
