package service

import (
	"context"
	"errors"

	"github.com/hexiaopi/blog-service/internal/model"
	"github.com/hexiaopi/blog-service/internal/store"
)

type UserSrv interface {
	CheckAuth(ctx context.Context, param *AuthRequest) error
	GetUser(ctx context.Context, name string) (*model.User, error)
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
	UserId   int    `json:"user_id"`
	UserName string `json:"username"`
	PassWord string `json:"password"`
	Captcha  string `json:"captcha"`
	Cid      string `json:"cid"`
}

func (svc *UserService) CheckAuth(ctx context.Context, param *AuthRequest) error {
	user, err := svc.store.Users().Get(ctx, param.UserName)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not exists")
	}
	param.UserId = user.ID
	if err := user.Compare(param.PassWord); err != nil {
		return err
	}
	return nil
}

func (svc *UserService) GetUser(ctx context.Context, name string) (*model.User, error) {
	user, err := svc.store.Users().Get(ctx, name)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not exists")
	}
	user.PassWord = ""
	return user, nil
}
