package request

type ProfileUpdateRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20,no_space,lowercase" extensions:"x-order=0"`
	Email    string `json:"email" binding:"required,email" extensions:"x-order=1"`
	Name     string `json:"name" binding:"omitempty,min=3,max=150" extensions:"x-order=2"`
	Bio      string `json:"bio" binding:"omitempty,max=700" extensions:"x-order=3"`
	Age      int    `json:"age" binding:"omitempty,min=0" extension:"z-order=4"`
	Gender   string `json:"gender" binding:"omitempty,uppercase,oneof=MALE FEMALE" extensions:"x-order=5"`
}
