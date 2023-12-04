package service

import (
	"context"
	"errors"

	"github.com/hexiaopi/blog-service/internal/entity"
	"github.com/hexiaopi/blog-service/internal/store"
	log "github.com/hexiaopi/blog-service/pkg/logger"
)

type UserSrv interface {
	CheckAuth(ctx context.Context, param *AuthRequest) error
	Get(ctx context.Context, name string) (*entity.User, error)
	List(ctx context.Context, param *ListUserRequest) ([]entity.User, int64, error)
	Create(ctx context.Context, param *CreateUserRequest) error
	Update(ctx context.Context, param *UpdateUserRequest) error
	Delete(ctx context.Context, param *DeleteUserRequest) error
}

type UserService struct {
	store  store.Factory
	logger log.Logger
}

var _ UserSrv = (*UserService)(nil)

func NewUserService(factory store.Factory, logger log.Logger) *UserService {
	return &UserService{
		store:  factory,
		logger: logger,
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
	svc.logger.Debugf("user check auth request:%+v", param)
	user, err := svc.store.Users().Get(ctx, param.UserName)
	if err != nil {
		svc.logger.Errorf("user store get err:%v", err)
		return err
	}
	if user == nil {
		return errors.New("user not exists")
	}
	if err := user.Compare(param.PassWord); err != nil {
		svc.logger.Errorf("user compare password err:%v", err)
		return err
	}
	return nil
}

func (svc *UserService) Get(ctx context.Context, name string) (*entity.User, error) {
	svc.logger.Debugf("user get request:%s", name)
	user, err := svc.store.Users().Get(ctx, name)
	if err != nil {
		svc.logger.Errorf("user store get err:%v", err)
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not exists")
	}
	roles, err := svc.store.UserRole().ListUserRole(ctx, user.ID)
	if err != nil {
		svc.logger.Errorf("get user role err:%v", err)
		return nil, err
	}
	user.Roles = roles
	return user, nil
}

type ListUserRequest struct {
	entity.ListOption
}

func (svc *UserService) List(ctx context.Context, param *ListUserRequest) ([]entity.User, int64, error) {
	svc.logger.Debugf("user list request:%+v", param)
	total, err := svc.store.Users().Count(ctx, &param.ListOption)
	if err != nil {
		svc.logger.Errorf("user store count err:%v", err)
		return nil, 0, err
	}
	if total == 0 {
		return nil, 0, nil
	}
	users, err := svc.store.Users().List(ctx, &param.ListOption)
	if err != nil {
		svc.logger.Errorf("user store list err:%v", err)
		return nil, 0, err
	}
	for i := 0; i < len(users); i++ {
		roles, err := svc.store.UserRole().ListUserRole(ctx, users[i].ID)
		if err != nil {
			svc.logger.Errorf("get user role err:%v", err)
			return nil, 0, err
		}
		users[i].Roles = roles
	}
	return users, total, nil
}

type CreateUserRequest struct {
	entity.User
}

func (svc *UserService) Create(ctx context.Context, param *CreateUserRequest) error {
	svc.logger.Debugf("user create request:%+v", param)
	if err := param.User.EncryptPassword(); err != nil {
		svc.logger.Errorf("user encrypt password err:%v", err)
		return nil
	}
	return svc.store.Tx(ctx, func(ctx context.Context, factory store.Factory) error {
		id, err := factory.Users().Create(ctx, &param.User)
		if err != nil {
			svc.logger.Errorf("user create err:%v", err)
			return err
		}
		for _, role := range param.Roles {
			if err := factory.UserRole().Create(ctx, id, role.ID); err != nil {
				svc.logger.Errorf("user role create err:%v", err)
				return err
			}
		}
		return nil
	})
}

type UpdateUserRequest struct {
	entity.User
}

func (svc *UserService) Update(ctx context.Context, param *UpdateUserRequest) error {
	svc.logger.Debugf("user update request:%+v", param)
	user, err := svc.store.Users().Get(ctx, param.User.Name)
	if err != nil {
		svc.logger.Errorf("user store get err:%v", err)
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}
	if err := param.User.EncryptPassword(); err != nil {
		svc.logger.Errorf("user encrypt password err:%v", err)
		return nil
	}
	return svc.store.Tx(ctx, func(ctx context.Context, factory store.Factory) error {
		if err := factory.Users().Update(ctx, &param.User); err != nil {
			svc.logger.Errorf("user store update err:%v", err)
			return err
		}
		if err := factory.UserRole().DeleteByUser(ctx, user.ID); err != nil {
			svc.logger.Errorf("user delete role err:%v", err)
			return err
		}
		for _, role := range param.Roles {
			if err := factory.UserRole().Create(ctx, user.ID, role.ID); err != nil {
				svc.logger.Errorf("user role create err:%v", err)
				return err
			}
		}
		return nil
	})
}

type DeleteUserRequest struct {
	entity.OneOption
}

func (svc *UserService) Delete(ctx context.Context, param *DeleteUserRequest) error {
	svc.logger.Debugf("user delete request:%+v", param)
	return svc.store.Tx(ctx, func(ctx context.Context, factory store.Factory) error {
		if err := factory.Users().Delete(ctx, param.Id); err != nil {
			svc.logger.Errorf("user store delete err:%v", err)
			return err
		}
		if err := factory.UserRole().DeleteByUser(ctx, param.Id); err != nil {
			svc.logger.Errorf("user role delete err:%v", err)
			return err
		}
		return nil
	})
}
