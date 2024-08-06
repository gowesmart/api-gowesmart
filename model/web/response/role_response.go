package response

type RoleResponse struct {
	ID     uint `json:"id"`
	Role   uint `json:"role"`
	UserID uint `json:"user_id"`
}

type RoleListResponse struct {
	Roles []RoleResponse `json:"roles"`
}
