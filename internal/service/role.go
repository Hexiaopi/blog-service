package service

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/hexiaopi/blog-service/internal/model"
	"github.com/hexiaopi/blog-service/internal/store"
)

type RoleSrv interface {
	List(ctx context.Context, param *ListRoleRequest) ([]model.Role, int64, error)
	Create(ctx context.Context, param *CreateRoleRequest) error
	Update(ctx context.Context, param *UpdateRoleRequest) error
	Delete(ctx context.Context, param *DeleteRoleRequest) error
}

type RoleService struct {
	store store.Factory
}

var _ RoleSrv = (*RoleService)(nil)

func NewRoleService(factory store.Factory) *RoleService {
	return &RoleService{
		store: factory,
	}
}

type ListRoleRequest struct {
	model.ListOption
}

func (svc *RoleService) List(ctx context.Context, param *ListRoleRequest) ([]model.Role, int64, error) {
	roles, err := svc.store.Roles().List(ctx, &param.ListOption)
	if err != nil {
		log.Errorf("role store list err:%v", err)
		return nil, 0, err
	}
	total, err := svc.store.Roles().Count(ctx, &param.ListOption)
	if err != nil {
		log.Errorf("role store count err:%v", err)
		return nil, 0, err
	}
	return roles, total, nil
}

type CreateRoleRequest struct {
	model.Role
}

func (svc *RoleService) Create(ctx context.Context, param *CreateRoleRequest) error {
	if err := svc.store.Roles().Create(ctx, &param.Role); err != nil {
		log.Errorf("role store create err:%v", err)
		return err
	}
	return nil
}

type UpdateRoleRequest struct {
	model.Role
}

func (svc *RoleService) Update(ctx context.Context, param *UpdateRoleRequest) error {
	if err := svc.store.Roles().Update(ctx, &param.Role); err != nil {
		log.Errorf("role store update err:%v", err)
		return err
	}
	return nil
}

type DeleteRoleRequest struct {
	model.OneOption
}

func (svc *RoleService) Delete(ctx context.Context, param *DeleteRoleRequest) error {
	if err := svc.store.Roles().Delete(ctx, param.Id); err != nil {
		log.Errorf("role store delete err:%v", err)
		return err
	}
	return nil
}
