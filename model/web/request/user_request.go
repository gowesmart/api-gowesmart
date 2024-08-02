package request

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20,no_space,lowercase" extensions:"x-order=0"`
	Email    string `json:"email" binding:"required,email" extensions:"x-order=1"`
	Password string `json:"password" binding:"required,min=8,no_space" example:"password" extensions:"x-order=2"`
}

type ForgotPasswordRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20,no_space,lowercase" extensions:"x-order=0"`
	Email    string `json:"email" binding:"required,email"`
}

type ResetPasswordRequest struct {
	NewPassword string `json:"new_password" binding:"required,min=8" example:"new_password" extensions:"x-order=0"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" extensions:"x-order=0"`
	Password string `json:"password" binding:"required" example:"password" extensions:"x-order=1"`
}
