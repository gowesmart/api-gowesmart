package request

type CreateReviewRequest struct {
	Comment string `json:"comment" binding:"required"`
	Rating  int    `json:"rating" binding:"required,min=1,max=5"`
	BikeID  uint   `json:"bike_id" binding:"required"`
	OrderID int    `json:"order_id" binding:"required"`
}

type UpdateReviewRequest struct {
	Comment string `json:"comment" binding:"required"`
	Rating  int    `json:"rating" binding:"required,min=1,max=5"`
}
