package request

type CreateReviewRequest struct {
	Comment string `json:"comment" binding:"required"`
	Rating  int    `json:"rating" binding:"required"`
	BikeID  uint   `json:"bike_id" binding:"required"`
	UserID  uint   `json:"user_id" binding:"required"`
}

type UpdateReviewRequest struct {
	Comment string `json:"comment" binding:"required"`
	Rating  int    `json:"rating" binding:"required"`
}
