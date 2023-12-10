package domain

type BusGetUserByEmailResponse struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
