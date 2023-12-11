package service

import (
	"context"

	"github.com/hexiaopi/blog-service/internal/entity"
	"github.com/hexiaopi/blog-service/internal/model"
	"github.com/hexiaopi/blog-service/internal/store"
	log "github.com/hexiaopi/blog-service/pkg/logger"
)

type SysRestSrv interface {
	Get(ctx context.Context, param *SysRestRequest) (*model.SysRest, error)
	List(ctx context.Context, param *SysRestListRequest) ([]entity.SysRest, int64, error)
	Create(ctx context.Context, param *CreateSysRestRequest) error
	Update(ctx context.Context, param *UpdateSysRestRequest) error
	Delete(ctx context.Context, id int) error
}

type SysRestService struct {
	store  store.Factory
	logger log.Logger
}

var _ SysRestSrv = (*SysRestService)(nil)

func NewSysRestService(factory store.Factory, logger log.Logger) *SysRestService {
	return &SysRestService{
		store:  factory,
		logger: logger,
	}
}

type SysRestRequest struct {
	entity.OneOption
}

func (svc *SysRestService) Get(ctx context.Context, param *SysRestRequest) (*model.SysRest, error) {
	svc.logger.Debugf("SysRest get request:%+v", param)
	SysRest, err := svc.store.SysRests().Get(ctx, param.Id)
	if err != nil {
		svc.logger.Errorf("SysRest store get err:%v", err)
		return nil, err
	}
	return SysRest, nil
}

type SysRestListRequest struct {
	entity.ListOption
}

func (svc *SysRestService) List(ctx context.Context, param *SysRestListRequest) ([]entity.SysRest, int64, error) {
	svc.logger.Debugf("SysRest list request:%+v", param)
	count, err := svc.store.SysRests().Count(ctx, &param.ListOption)
	if err != nil {
		svc.logger.Errorf("SysRest store count err:%v", err)
		return nil, 0, err
	}
	if count == 0 {
		return nil, 0, nil
	}
	SysRests, err := svc.store.SysRests().List(ctx, &param.ListOption)
	if err != nil {
		svc.logger.Errorf("SysRest store list err:%v", err)
		return nil, 0, err
	}
	result := make([]entity.SysRest, 0, len(SysRests))
	for _, rest := range SysRests {
		sysRest := entity.ToEntitySysRest(&rest)
		param.ParentId = rest.ID
		childRests, err := svc.store.SysRests().List(ctx, &param.ListOption)
		if err != nil {
			svc.logger.Errorf("SysRest list children err:%v", err)
		}
		sysRest.Children = make([]entity.SysRest, 0, len(childRests))
		for _, childRest := range childRests {
			child := entity.ToEntitySysRest(&childRest)
			sysRest.Children = append(sysRest.Children, *child)
		}
		result = append(result, *sysRest)
	}
	return result, count, nil
}

type CreateSysRestRequest struct {
	model.SysRest
}

func (svc *SysRestService) Create(ctx context.Context, param *CreateSysRestRequest) error {
	svc.logger.Debugf("SysRest create request:%+v", param)
	if err := svc.store.SysRests().Create(ctx, &param.SysRest); err != nil {
		svc.logger.Errorf("SysRest store create err:%v", err)
		return err
	}
	return nil
}

type UpdateSysRestRequest struct {
	model.SysRest
}

func (svc *SysRestService) Update(ctx context.Context, param *UpdateSysRestRequest) error {
	svc.logger.Debugf("SysRest update request:%+v", param)
	err := svc.store.SysRests().Update(ctx, &param.SysRest)
	if err != nil {
		svc.logger.Errorf("SysRest store update err:%v", err)
		return err
	}
	return nil
}

func (svc *SysRestService) Delete(ctx context.Context, id int) error {
	svc.logger.Debugf("SysRest delete request:%d", id)
	if err := svc.store.SysRests().Delete(ctx, id); err != nil {
		svc.logger.Errorf("SysRest store delete err:%v", err)
		return err
	}
	return nil
}
