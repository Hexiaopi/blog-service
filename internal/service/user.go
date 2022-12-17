package service

import (
	"context"
	"errors"

	"github.com/hexiaopi/blog-service/internal/store"
)

type UserSrv interface {
	CheckAuth(ctx context.Context, param *AuthRequest) error
}

type UserService struct {
	store store.Factory
}

var _ UserSrv = (*UserService)(nil)

func NewUserService(factory store.Factory) *UserService {
	return &UserService{
		store: factory,
	}
}

type AuthRequest struct {
	UserName string
	PassWord string
}

func (svc *UserService) CheckAuth(ctx context.Context, param *AuthRequest) error {
	user, err := svc.store.Users().Get(ctx, param.UserName)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not exists")
	}
	return user.Compare(param.PassWord)
}
