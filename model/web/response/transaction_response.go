package response

import "time"

type TransactionResponse struct {
	ID         int             `json:"id"`
	TotalPrice int             `json:"total_price"`
	UserID     int             `json:"user_id"`
	Status     string          `json:"status"`
	Orders     []OrderResponse `json:"orders"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  time.Time       `json:"upodated_at"`
}

type UserTransactionResponse struct {
	ID          uint                  `json:"id"`
	Username    string                `json:"username"`
	Transaction []TransactionResponse `json:"transaction"`
}

type CreateTransactionResponse struct {
	TransactionID int `json:"transaction_id"`
}

type GetAllTransactionResponse struct {
	ID         int                           `json:"id"`
	TotalPrice int                           `json:"total_price"`
	User       GetAllTransactionUserResponse `json:"user"`
	Status     string                        `json:"status"`
	Orders     []GetAllOrderResponse         `json:"orders"`
	CreatedAt  time.Time                     `json:"created_at"`
	UpdatedAt  time.Time                     `json:"upodated_at"`
}

type GetAllTransactionUserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}
