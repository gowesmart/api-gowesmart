package response

import "time"

type TransactionResponse struct {
	ID         int             `json:"id"`
	TotalPrice int             `json:"totalPrice"`
	UserID     int             `json:"userId"`
	Status     string          `json:"status"`
	Orders     []OrderResponse `json:"orders"`
	CreatedAt  time.Time       `json:"autoCreateTime"`
	UpdatedAt  time.Time       `json:"autoUpdateTime"`
}

type UserTransactionResponse struct {
	ID          uint                  `json:"id"`
	Username    string                `json:"username"`
	Transaction []TransactionResponse `json:"transaction"`
}
