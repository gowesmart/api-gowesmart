package request

type CartCreateRequest struct {
	UserID int `json:"userId" binding:"required"`
}

type CartUpdateRequest struct {
	ID     int `json:"id" binding:"required"`
	UserID int `json:"userId" binding:"required"`
}
