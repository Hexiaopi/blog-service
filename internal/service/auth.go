package service

import (
	"context"
	"errors"

	"github.com/hexiaopi/blog-service/internal/store"
)

type AuthService struct {
	store store.Factory
}

func NewAuthService(factory store.Factory) AuthService {
	return AuthService{
		store: factory,
	}
}

type AuthRequest struct {
	AppKey    string
	AppSecret string
}

func (svc *AuthService) CheckAuth(ctx context.Context, param *AuthRequest) error {
	auth, err := svc.store.Auths().Get(ctx, param.AppKey, param.AppSecret)
	if err != nil {
		return err
	}
	if auth == nil {
		return errors.New("auth info does not exists")
	}
	return nil
}
