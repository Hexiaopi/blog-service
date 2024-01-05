package service

import (
	"context"

	"github.com/hexiaopi/blog-service/internal/entity"
	"github.com/hexiaopi/blog-service/internal/model"
	"github.com/hexiaopi/blog-service/internal/store"
	log "github.com/hexiaopi/blog-service/pkg/logger"
)

type SysMenuSrv interface {
	Get(ctx context.Context, param *SysMenuRequest) (*model.SysMenu, error)
	Tree(ctx context.Context, param *SysMenuListRequest) ([]entity.MenuTree, int64, error)
	List(ctx context.Context, param *SysMenuListRequest) ([]entity.SysMenu, int64, error)
	Create(ctx context.Context, param *CreateSysMenuRequest) error
	Update(ctx context.Context, param *UpdateSysMenuRequest) error
	Delete(ctx context.Context, id int) error
}

type SysMenuService struct {
	store  store.Factory
	logger log.Logger
}

var _ SysMenuSrv = (*SysMenuService)(nil)

func NewSysMenuService(factory store.Factory, logger log.Logger) *SysMenuService {
	return &SysMenuService{
		store:  factory,
		logger: logger,
	}
}

type SysMenuRequest struct {
	entity.OneOption
}

func (svc *SysMenuService) Get(ctx context.Context, param *SysMenuRequest) (*model.SysMenu, error) {
	svc.logger.Debugf("SysMenu get request:%+v", param)
	SysMenu, err := svc.store.SysMenus().Get(ctx, param.Id)
	if err != nil {
		svc.logger.Errorf("SysMenu store get err:%v", err)
		return nil, err
	}
	return SysMenu, nil
}

type SysMenuListRequest struct {
	entity.ListOption
}

func (svc *SysMenuService) Tree(ctx context.Context, param *SysMenuListRequest) ([]entity.MenuTree, int64, error) {
	svc.logger.Debugf("SysMenu list request:%+v", param)
	count, err := svc.store.SysMenus().Count(ctx, &param.ListOption)
	if err != nil {
		svc.logger.Errorf("SysMenu store count err:%v", err)
		return nil, 0, err
	}
	if count == 0 {
		return nil, 0, nil
	}
	SysMenus, err := svc.store.SysMenus().List(ctx, &param.ListOption)
	if err != nil {
		svc.logger.Errorf("SysMenu store list err:%v", err)
		return nil, 0, err
	}
	result := make([]entity.MenuTree, 0, len(SysMenus))
	for _, menu := range SysMenus {
		sysMenu := entity.ToEntityMenuTree(&menu)
		param.ParentId = menu.ID
		childMenus, err := svc.store.SysMenus().List(ctx, &param.ListOption)
		if err != nil {
			svc.logger.Errorf("SysRest list children err:%v", err)
		}
		sysMenu.Children = make([]entity.MenuTree, 0, len(childMenus))
		for _, childMenu := range childMenus {
			child := entity.ToEntityMenuTree(&childMenu)
			sysMenu.Children = append(sysMenu.Children, *child)
		}
		result = append(result, *sysMenu)
	}
	return result, count, nil
}

func (svc *SysMenuService) List(ctx context.Context, param *SysMenuListRequest) ([]entity.SysMenu, int64, error) {
	svc.logger.Debugf("SysMenu list request:%+v", param)
	count, err := svc.store.SysMenus().Count(ctx, &param.ListOption)
	if err != nil {
		svc.logger.Errorf("SysMenu store count err:%v", err)
		return nil, 0, err
	}
	if count == 0 {
		return nil, 0, nil
	}
	SysMenus, err := svc.store.SysMenus().List(ctx, &param.ListOption)
	if err != nil {
		svc.logger.Errorf("SysMenu store list err:%v", err)
		return nil, 0, err
	}
	result := make([]entity.SysMenu, 0, len(SysMenus))
	for _, menu := range SysMenus {
		sysMenu := entity.ToEntitySysMenu(&menu)
		param.ParentId = menu.ID
		childMenus, err := svc.store.SysMenus().List(ctx, &param.ListOption)
		if err != nil {
			svc.logger.Errorf("SysRest list children err:%v", err)
		}
		sysMenu.Children = make([]entity.SysMenu, 0, len(childMenus))
		for _, childMenu := range childMenus {
			child := entity.ToEntitySysMenu(&childMenu)
			sysMenu.Children = append(sysMenu.Children, *child)
		}
		result = append(result, *sysMenu)
	}
	return result, count, nil
}

type CreateSysMenuRequest struct {
	model.SysMenu
}

func (svc *SysMenuService) Create(ctx context.Context, param *CreateSysMenuRequest) error {
	svc.logger.Debugf("SysMenu create request:%+v", param)
	if err := svc.store.SysMenus().Create(ctx, &param.SysMenu); err != nil {
		svc.logger.Errorf("SysMenu store create err:%v", err)
		return err
	}
	return nil
}

type UpdateSysMenuRequest struct {
	model.SysMenu
}

func (svc *SysMenuService) Update(ctx context.Context, param *UpdateSysMenuRequest) error {
	svc.logger.Debugf("SysMenu update request:%+v", param)
	err := svc.store.SysMenus().Update(ctx, &param.SysMenu)
	if err != nil {
		svc.logger.Errorf("SysMenu store update err:%v", err)
		return err
	}
	return nil
}

func (svc *SysMenuService) Delete(ctx context.Context, id int) error {
	svc.logger.Debugf("SysMenu delete request:%d", id)
	if err := svc.store.SysMenus().Delete(ctx, id); err != nil {
		svc.logger.Errorf("SysMenu store delete err:%v", err)
		return err
	}
	return nil
}
