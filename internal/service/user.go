package service

import (
	"context"
	"errors"
	"log"
	"time"

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
	user, err := svc.store.Users().Get(ctx, param.User.Name)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}
	roleExist := make(map[int]string)
	for _, role := range user.Roles {
		roleExist[role.ID] = "delete" //默认删除
	}
	for _, role := range param.Roles {
		if _, ok := roleExist[role.ID]; ok {
			roleExist[role.ID] = "exist" //原本角色
		} else {
			roleExist[role.ID] = "create" //新的角色
		}
	}
	return svc.store.Tx(ctx, func(ctx context.Context, factory store.Factory) error {
		if err := factory.Users().Update(ctx, &param.User); err != nil {
			return err
		}
		for k, v := range roleExist {
			if v == "exist" {
				continue
			}
			if v == "create" {
				if err := factory.UserRole().Create(ctx, &model.UserRole{
					UserId:     user.ID,
					RoleId:     k,
					CreateTime: time.Now(),
				}); err != nil {
					return err
				}
			}
			if v == "delete" {
				if err := factory.UserRole().Delete(ctx, &model.UserRole{
					UserId: user.ID,
					RoleId: k,
				}); err != nil {
					return err
				}
			}
		}
		return nil
	})
}

type DeleteUserRequest struct {
	model.OneOption
}

func (svc *UserService) Delete(ctx context.Context, param *DeleteUserRequest) error {
	return svc.store.Users().Delete(ctx, param.Id)
}
