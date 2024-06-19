package user

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID          uuid.UUID  `json:"-" db:"id" fieldtag:"pk"`
	Login       string     `json:"login" db:"login"`
	Password    string     `json:"-" db:"password"`
	Name        *string    `json:"name,omitempty" db:"name"`
	Surname     *string    `json:"surname,omitempty" db:"surname"`
	Birthday    *time.Time `json:"birthday,omitempty" db:"birthday"`
	Email       *string    `json:"email,omitempty" db:"email"`
	PhoneNumber *string    `json:"phone_number,omitempty" db:"phone_number"`
}
