package user

import (
	"authorizer/core/repository"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// InvalidUserPasswordError error when password and email dont match
var InvalidUserPasswordError = errors.New("invalid email or password")

var validate = validator.New()

// User model
type User struct {
	repository.Model
	Username  string `gorm:"unique;not null" json:"username" validate:"required"`
	Email     string `gorm:"unique;not null" json:"email" validate:"required,email"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}

// Validate user fields
func (u *User) Validate() error {
	return validate.Struct(u)
}

// CreateUserRequest model
type CreateUserRequest struct {
	User     User   `json:"user" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (c *CreateUserRequest) Validate() error {
	return validate.Struct(c)
}

// Password model saves user encrypted password
type Password struct {
	repository.Model
	Password string    `json:"password" validate:"required"`
	UserID   uuid.UUID `json:"user_id" gorm:"index:userID;not null,column:user_id" validate:"required"`
	User     User      `json:"-" validate:"-"`
}

// Validate user password fields
func (u *Password) Validate() error {
	return validate.Struct(u)
}

// HashPassword hashes user password and returns userPassword Model
func HashPassword(password, userID string) (*Password, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return nil, err
	}

	id, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	return &Password{
		Password: string(bytes),
		UserID:   id,
	}, nil
}

// CheckPasswordHash checks if password matches hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
