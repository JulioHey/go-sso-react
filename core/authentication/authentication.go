package authentication

import (
	. "authorizer/core/autherror"
	"github.com/go-playground/validator/v10"
	"github.com/hashicorp/errwrap"
)

var (
	validate = validator.New()
)

// Request to authenticate user
type Request struct {
	Scope        string         `json:"scope" validate:"required"`
	ClientID     string         `json:"client_id" validate:"required"`
	ResponseType string         `json:"response_type" validate:"isResponseTypeValid"`
	RedirectURI  string         `json:"redirect_uri" validate:"required"`
	State        string         `json:"state" validate:"required"`
	Nonce        string         `json:"nonce" validate:"required"`
	Extra        map[string]any `json:"extra"`
}

// isResponseTypeValid checks if the response type is valid.
func isResponseTypeValid(fl validator.FieldLevel) bool {
	return fl.Field().String() == "code"
}

// Validate request
func (r *Request) Validate() error {
	validate.RegisterValidation("isResponseTypeValid", isResponseTypeValid)
	err := validate.Struct(r)
	if err != nil {
		return errwrap.Wrap(ErrInvalidRequest, err)
	}
	return nil
}

// Response to authenticate user
type Response struct {
	Code         string `json:"code" `
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
