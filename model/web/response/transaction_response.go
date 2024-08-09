package response

type TransactionResponse struct {
	ID         int             `json:"id"`
	TotalPrice int             `json:"total_price"`
	UserID     int             `json:"user_id"`
	Status     string          `json:"status"`
	Orders     []OrderResponse `json:"orders"`
	CreatedAt  string       `json:"created_at"`
	UpdatedAt  string       `json:"upodated_at"`
}

type UserTransactionResponse struct {
	ID          uint                  `json:"id"`
	Username    string                `json:"username"`
	Transaction []TransactionResponse `json:"transaction"`
}

type CreateTransactionResponse struct {
	TransactionID int `json:"transaction_id"`
}