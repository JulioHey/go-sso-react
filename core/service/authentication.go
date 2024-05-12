package service

import (
	"authorizer/core/authentication"
	"context"
	"gorm.io/gorm"
)

type AuthManager interface {
	Authenticate(ctx context.Context, tx *gorm.DB, req *authentication.Request) error
}

type CodeManager interface {
	CreateCode(ctx context.Context, tx *gorm.DB,
		req *authentication.Request) (*authentication.Response, error)
}

type AuthService struct {
	db          *gorm.DB
	authManager AuthManager
	sessManager CodeManager
}

func (a *AuthService) AuthenticationFlow(ctx context.Context,
	request *authentication.Request) (*authentication.Response, error) {
	err := a.authManager.Authenticate(ctx, a.db, request)
	if err != nil {
		return nil, err
	}

	res, err := a.sessManager.CreateCode(ctx, a.db, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}
