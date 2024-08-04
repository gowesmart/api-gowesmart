package response

type ProfileResponse struct {
	ID       uint   `json:"id" example:"1" extensions:"x-order=0"`
	Username string `json:"username" example:"luigi" extensions:"x-order=1"`
	Email    string `json:"email" example:"luigi@sam.com" extensions:"x-order=2"`
	Name     string `json:"name,omitempty" example:"Luigi Di Caprio" extensions:"x-order=3"`
	Bio      string `json:"bio,omitempty" example:"I am Luigi" extensions:"x-order=4"`
	Age      int    `json:"age,omitempty" example:"18" extensions:"x-order=5"`
	Gender   string `json:"gender,omitempty" example:"MALE" extensions:"x-order=6"`
}
