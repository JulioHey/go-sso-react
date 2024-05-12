package postgres

import (
	"authorizer/core/authentication"
	user "authorizer/core/user"
	"context"
	"gorm.io/gorm"
	"log"
)

type AuthStore struct{}

// Authenticate authenticates a user by checking if the provided email and password match the user's credentials.
//
// Parameters:
// - ctx: The context.Context object for the authentication request.
// - tx: The *gorm.DB object for the database transaction.
// - req: The *authentication.Request object containing the user's email and password.
//
// Returns:
// - error: An error if the authentication fails, otherwise nil.
func (a *AuthStore) Authenticate(ctx context.Context, tx *gorm.DB, req *authentication.Request) error {
	var u *user.User
	err := tx.Where("email = ?", req.Extra["email"]).First(&u).Error
	if err != nil {
		return err
	}
	var password *user.Password
	err = tx.Where("user_id = ?", u.ID).First(&password).Error
	if err != nil {
		return err
	}
	log.Printf("PORRA %+v", password)

	if !user.CheckPasswordHash(req.Extra["password"].(string), password.Password) {
		return user.InvalidUserPasswordError
	}

	return nil
}
