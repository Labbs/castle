package domain

type LoginRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}
