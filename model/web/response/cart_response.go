package response

import "time"

type CartResponse struct {
	ID        uint                `json:"id"`
	UserID    uint                `json:"user_id"`
	CartItems []CartItemResponse `json:"cart_items"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}
