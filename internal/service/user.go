package service

import (
	"context"
	"errors"
	"log"

	"github.com/hexiaopi/blog-service/internal/model"
	"github.com/hexiaopi/blog-service/internal/store"
)

type UserSrv interface {
	CheckAuth(ctx context.Context, param *AuthRequest) error
	Get(ctx context.Context, name string) (*model.User, error)
	List(ctx context.Context, param *ListUserRequest) ([]model.User, int64, error)
	Create(ctx context.Context, param *CreateUserRequest) error
	Update(ctx context.Context, param *UpdateUserRequest) error
	Delete(ctx context.Context, param *DeleteUserRequest) error
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

func (svc *UserService) Get(ctx context.Context, name string) (*model.User, error) {
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

type ListUserRequest struct {
	model.ListOption
}

func (svc *UserService) List(ctx context.Context, param *ListUserRequest) ([]model.User, int64, error) {
	users, err := svc.store.Users().List(ctx, &param.ListOption)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}
	for i := range users {
		users[i].PassWord = ""
	}
	total, err := svc.store.Users().Count(ctx, &param.ListOption)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}
	return users, total, nil
}

type CreateUserRequest struct {
	model.User
}

func (svc *UserService) Create(ctx context.Context, param *CreateUserRequest) error {
	if err := svc.store.Users().Create(ctx, &param.User); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

type UpdateUserRequest struct {
	model.User
}

func (svc *UserService) Update(ctx context.Context, param *UpdateUserRequest) error {
	if err := svc.store.Users().Update(ctx, &param.User); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

type DeleteUserRequest struct {
	model.OneOption
}

func (svc *UserService) Delete(ctx context.Context, param *DeleteUserRequest) error {
	return svc.store.Users().Delete(ctx, param.Id)
}
