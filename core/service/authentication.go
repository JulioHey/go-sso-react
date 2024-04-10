package service

import (
	"authorizer/core/authentication"
	"context"
)

type AuthManager interface {
	Authenticate(ctx context.Context, req *authentication.Request) error
}

type CodeManager interface {
	CreateCode(ctx context.Context,
		req *authentication.Request) (*authentication.Response, error)
}

type AuthService struct {
	authManager AuthManager
	sessManager CodeManager
}

func (a *AuthService) AuthenticationFlow(ctx context.Context,
	request *authentication.Request) (*authentication.Response, error) {
	err := a.authManager.Authenticate(ctx, request)
	if err != nil {
		return nil, err
	}

	res, err := a.sessManager.CreateCode(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}
