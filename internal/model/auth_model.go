package model

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
}
