package user

import (
	"authorizer/core/repository"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type User struct {
	repository.Model
	Username  string `gorm:"unique;not null" json:"username" validate:"required"`
	Email     string `gorm:"unique;not null" json:"email" validate:"required,email"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}

func (u *User) Validate() error {
	return validate.Struct(u)
}
