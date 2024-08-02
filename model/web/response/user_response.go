package response

type RegisterResponse struct {
	Username string `json:"username" example:"luigi" extensions:"x-order=0"`
	Email    string `json:"email" example:"luigi@sam.com" extensions:"x-order=1"`
	Role     string `json:"role" example:"USER" extensions:"x-order=2"`
}

type GetUserCurrentResponse struct {
	ID       uint   `json:"id" example:"1" extensions:"x-order=0"`
	Username string `json:"username" example:"luigi" extensions:"x-order=2"`
	Email    string `json:"email" example:"luigi@sam.com" extensions:"x-order=3"`
	Role     string `json:"role" example:"USER" extensions:"x-order=4"`
}

type ForgotPasswordResponse struct {
	ForgotPasswordToken string `json:"forgot_password_token" example:"token"`
}

type LoginResponse struct {
	Username string `json:"username" example:"luigi" extensions:"x-order=0"`
	Email    string `json:"email" example:"luigi@sam.com" extensions:"x-order=1"`
	Role     string `json:"role" example:"USER" extensions:"x-order=2"`
	Token    string `json:"token" example:"token" extensions:"x-order=3"`
}
