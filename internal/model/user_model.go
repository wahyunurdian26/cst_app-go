package model

import "github.com/google/uuid"

// UserCreateRequest digunakan untuk pembuatan user baru
type UserCreateRequest struct {
	Email                  string `json:"email" validate:"required,email"`
	Username               string `json:"username" validate:"required"`
	Password               string `json:"password" validate:"required,min=6"`
	IDRole                 string `json:"id_role" validate:"required"`
	IDBusinessGroup        string `json:"id_business_group" validate:"required"`
	IDSubBusinessGroup     string `json:"id_sub_business_group" validate:"required"`
	EmailPIC               string `json:"email_pic" validate:"required,email"`
	StatusActive           bool   `json:"status_active"`
	IDBusinessGroupDigital string `json:"id_business_group_digital" validate:"required"`
}

// UserUpdateRequest digunakan untuk pembaruan data user
type UserUpdateRequest struct {
	Id                     uuid.UUID `json:"id" validate:"omitempty"`
	Username               string    `json:"username" validate:"omitempty"`
	Password               string    `json:"password" validate:"omitempty,min=6"`
	IDRole                 string    `json:"id_role" validate:"omitempty"`
	IDBusinessGroup        string    `json:"id_business_group" validate:"omitempty"`
	IDSubBusinessGroup     string    `json:"id_sub_business_group" validate:"omitempty"`
	EmailPIC               string    `json:"email_pic" validate:"omitempty,email"`
	StatusActive           bool      `json:"status_active" validate:"omitempty"`
	IDBusinessGroupDigital string    `json:"id_business_group_digital" validate:"omitempty"`
}

type UserDeleteRequest struct {
	Id uuid.UUID `json:"id" validate:"omitempty"`
}

type UserGetByIdRequest struct {
	Id uuid.UUID `json:"id" validate:"omitempty"`
}
