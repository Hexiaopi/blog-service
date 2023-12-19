package service

import (
	"context"

	"github.com/hexiaopi/blog-service/internal/entity"
	"github.com/hexiaopi/blog-service/internal/store"
	log "github.com/hexiaopi/blog-service/pkg/logger"
)

type SysMenuSrv interface {
	List(ctx context.Context, param *SysMenuListRequest) ([]entity.SysMenu, int64, error)
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

type SysMenuListRequest struct {
	entity.ListOption
}

func (svc *SysMenuService) List(ctx context.Context, param *SysMenuListRequest) ([]entity.SysMenu, int64, error) {
	svc.logger.Debugf("SysMenu list request:%+v", param)
	count, err := svc.store.SysRests().Count(ctx, &param.ListOption)
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
