package user

import (
	"gateway/internal/model"
	"github.com/google/uuid"
	"time"
)

type SignUpInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type SignUpOutput struct {
	model.Response
}

type SignInInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type SignInOutput struct {
	model.Response
	Token string `json:"token"`
}

type UpdateInfoInput struct {
	UserId      uuid.UUID  `json:"userId"`
	Name        *string    `json:"name,omitempty"`
	Surname     *string    `json:"surname,omitempty"`
	Birthday    *time.Time `json:"birthday,omitempty"`
	Email       *string    `json:"email,omitempty"`
	PhoneNumber *string    `json:"phone_number,omitempty"`
}

type UpdateInfoOutput struct {
	model.Response
	User *User `json:"user"`
}
