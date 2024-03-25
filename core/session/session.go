package session

import (
	. "authorizer/core/autherror"
	"github.com/go-playground/validator/v10"
	"github.com/hashicorp/errwrap"
)

var (
	validate = validator.New()
)

type Session struct {
	State string `json:"state" validate:"required"`
}

func (s *Session) Validate() error {
	err := validate.Struct(s)
	if err != nil {
		return errwrap.Wrap(ErrInvalidRequest, err)
	}
	return nil
}
