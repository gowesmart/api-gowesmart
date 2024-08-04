package response

import (
	"time"
)

type ReviewResponse struct {
	ID        uint      `json:"id"`
	Comment   string    `json:"comment"`
	Rating    int       `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	BikeID    uint      `json:"bike_id"`
	UserID    uint      `json:"user_id"`
}
