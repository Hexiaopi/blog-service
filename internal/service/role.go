package service

import (
	"context"
	"errors"

	"github.com/hexiaopi/blog-service/internal/entity"
	"github.com/hexiaopi/blog-service/internal/store"
	log "github.com/hexiaopi/blog-service/pkg/logger"
)

type RoleSrv interface {
	List(ctx context.Context, param *ListRoleRequest) ([]entity.Role, int64, error)
	Create(ctx context.Context, param *CreateRoleRequest) error
	Update(ctx context.Context, param *UpdateRoleRequest) error
	Delete(ctx context.Context, param *DeleteRoleRequest) error
	UpdateMenu(ctx context.Context, param *UpdateRoleMenuRequest) error
	UpdateRest(ctx context.Context, param *UpdateRoleRestRequest) error
}

type RoleService struct {
	store  store.Factory
	logger log.Logger
}

var _ RoleSrv = (*RoleService)(nil)

func NewRoleService(factory store.Factory, logger log.Logger) *RoleService {
	return &RoleService{
		store:  factory,
		logger: logger,
	}
}

type ListRoleRequest struct {
	entity.ListOption
}

func (svc *RoleService) List(ctx context.Context, param *ListRoleRequest) ([]entity.Role, int64, error) {
	svc.logger.Debugf("role list request:%+v", param)
	total, err := svc.store.Roles().Count(ctx, &param.ListOption)
	if err != nil {
		svc.logger.Errorf("role store count err:%v", err)
		return nil, 0, err
	}
	if total == 0 {
		return nil, 0, err
	}
	roles, err := svc.store.Roles().List(ctx, &param.ListOption)
	if err != nil {
		svc.logger.Errorf("role store list err:%v", err)
		return nil, 0, err
	}
	for i, role := range roles {
		roles[i].MenuIds, err = svc.store.RoleMenus().ListByRole(ctx, role.ID)
		if err != nil {
			svc.logger.Errorf("role menu get err:%v", err)
			return nil, 0, err
		}
		roles[i].RestIds, err = svc.store.RoleRests().ListByRole(ctx, role.ID)
		if err != nil {
			svc.logger.Errorf("role rest get err:%v", err)
			return nil, 0, err
		}
	}
	return roles, total, nil
}

type CreateRoleRequest struct {
	entity.Role
}

func (svc *RoleService) Create(ctx context.Context, param *CreateRoleRequest) error {
	svc.logger.Debugf("role create request:%+v", param)
	if err := svc.store.Roles().Create(ctx, &param.Role); err != nil {
		svc.logger.Errorf("role store create err:%v", err)
		return err
	}
	return nil
}

type UpdateRoleRequest struct {
	entity.Role
}

func (svc *RoleService) Update(ctx context.Context, param *UpdateRoleRequest) error {
	svc.logger.Debugf("role update request:%+v", param)
	if err := svc.store.Roles().Update(ctx, &param.Role); err != nil {
		svc.logger.Errorf("role store update err:%v", err)
		return err
	}
	return nil
}

type UpdateRoleMenuRequest struct {
	RoleId  int   `json:"role_id"`
	MenuIds []int `json:"menu_ids"`
}

func (svc *RoleService) UpdateMenu(ctx context.Context, param *UpdateRoleMenuRequest) error {
	svc.logger.Debugf("update role menu request:%+v", param)
	role, err := svc.store.Roles().Get(ctx, param.RoleId)
	if err != nil {
		svc.logger.Errorf("query role err:%v", err)
		return err
	}
	if role == nil {
		return errors.New("role not found")
	}
	return svc.store.Tx(ctx, func(ctx context.Context, factory store.Factory) error {
		if err := factory.RoleMenus().DeleteByRole(ctx, param.RoleId); err != nil {
			svc.logger.Errorf("role delete menu err:%v", err)
			return err
		}
		for _, id := range param.MenuIds {
			if err := factory.RoleMenus().Create(ctx, param.RoleId, id); err != nil {
				svc.logger.Errorf("role add menu err:%v", err)
				return err
			}
		}
		return nil
	})
}

type UpdateRoleRestRequest struct {
	RoleId  int   `json:"role_id"`
	RestIds []int `json:"rest_ids"`
}

func (svc *RoleService) UpdateRest(ctx context.Context, param *UpdateRoleRestRequest) error {
	svc.logger.Debugf("update role rest request:%+v", param)
	role, err := svc.store.Roles().Get(ctx, param.RoleId)
	if err != nil {
		svc.logger.Errorf("query role err:%v", err)
		return err
	}
	if role == nil {
		return errors.New("role not found")
	}
	return svc.store.Tx(ctx, func(ctx context.Context, factory store.Factory) error {
		if err := factory.RoleRests().DeleteByRole(ctx, param.RoleId); err != nil {
			svc.logger.Errorf("role delete rest err:%v", err)
			return err
		}
		for _, id := range param.RestIds {
			if err := factory.RoleRests().Create(ctx, param.RoleId, id); err != nil {
				svc.logger.Errorf("role add rest err:%v", err)
				return err
			}
		}
		return nil
	})
}

type DeleteRoleRequest struct {
	entity.OneOption
}

func (svc *RoleService) Delete(ctx context.Context, param *DeleteRoleRequest) error {
	svc.logger.Debugf("role delete request:%+v", param)
	if err := svc.store.Roles().Delete(ctx, param.Id); err != nil {
		svc.logger.Errorf("role store delete err:%v", err)
		return err
	}
	return nil
}
