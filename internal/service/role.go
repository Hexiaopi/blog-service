package service

import (
	"context"

	"github.com/hexiaopi/blog-service/internal/model"
	"github.com/hexiaopi/blog-service/internal/store"
	log "github.com/hexiaopi/blog-service/pkg/logger"
)

type RoleSrv interface {
	List(ctx context.Context, param *ListRoleRequest) ([]model.Role, int64, error)
	Create(ctx context.Context, param *CreateRoleRequest) error
	Update(ctx context.Context, param *UpdateRoleRequest) error
	Delete(ctx context.Context, param *DeleteRoleRequest) error
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
	model.ListOption
}

func (svc *RoleService) List(ctx context.Context, param *ListRoleRequest) ([]model.Role, int64, error) {
	svc.logger.Debugf("role list request:%+v", param)
	roles, err := svc.store.Roles().List(ctx, &param.ListOption)
	if err != nil {
		svc.logger.Errorf("role store list err:%v", err)
		return nil, 0, err
	}
	total, err := svc.store.Roles().Count(ctx, &param.ListOption)
	if err != nil {
		svc.logger.Errorf("role store count err:%v", err)
		return nil, 0, err
	}
	return roles, total, nil
}

type CreateRoleRequest struct {
	model.Role
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
	model.Role
}

func (svc *RoleService) Update(ctx context.Context, param *UpdateRoleRequest) error {
	svc.logger.Debugf("role update request:%+v", param)
	if err := svc.store.Roles().Update(ctx, &param.Role); err != nil {
		svc.logger.Errorf("role store update err:%v", err)
		return err
	}
	return nil
}

type DeleteRoleRequest struct {
	model.OneOption
}

func (svc *RoleService) Delete(ctx context.Context, param *DeleteRoleRequest) error {
	svc.logger.Debugf("role delete request:%+v", param)
	if err := svc.store.Roles().Delete(ctx, param.Id); err != nil {
		svc.logger.Errorf("role store delete err:%v", err)
		return err
	}
	return nil
}
