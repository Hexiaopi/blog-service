package service

import (
	"context"
	"encoding/base64"

	"github.com/hexiaopi/blog-service/internal/entity"
	"github.com/hexiaopi/blog-service/internal/store"
	log "github.com/hexiaopi/blog-service/pkg/logger"
)

type ResourceSrv interface {
	Get(ctx context.Context, request *ResourceRequest) (*entity.Resource, error)
	List(ctx context.Context, param *ResourceListRequest) ([]entity.Resource, int64, error)
	Create(ctx context.Context, param *CreateResourceRequest) error
	Update(ctx context.Context, param *UpdateResourceRequest) error
	Delete(ctx context.Context, id int) error
}

type ResourceService struct {
	store  store.Factory
	logger log.Logger
}

var _ ResourceSrv = (*ResourceService)(nil)

func NewResourceService(factory store.Factory, logger log.Logger) *ResourceService {
	return &ResourceService{
		store:  factory,
		logger: logger,
	}
}

type ResourceRequest struct {
	entity.OneOption
}

func (svc *ResourceService) Get(ctx context.Context, param *ResourceRequest) (*entity.Resource, error) {
	svc.logger.Debugf("resource get request:%+v", param)
	resource, err := svc.store.Resources().Get(ctx, param.Id)
	if err != nil {
		svc.logger.Errorf("resource store get err:%v", err)
		return nil, err
	}
	resource.Base64 = "data:image/png;base64," + base64.StdEncoding.EncodeToString(resource.Blob)
	return resource, nil
}

type ResourceListRequest struct {
	entity.ListOption
}

func (svc *ResourceService) List(ctx context.Context, param *ResourceListRequest) ([]entity.Resource, int64, error) {
	svc.logger.Debugf("resource list request:%+v", param)
	resources, err := svc.store.Resources().List(ctx, &param.ListOption)
	if err != nil {
		svc.logger.Errorf("resource store list err:%v", err)
		return nil, 0, err
	}
	for i := range resources {
		resources[i].Base64 = "data:image/png;base64," + base64.RawStdEncoding.EncodeToString(resources[i].Blob)
		resources[i].Blob = nil
	}
	count, err := svc.store.Resources().Count(ctx, &param.ListOption)
	if err != nil {
		svc.logger.Errorf("resource store count err:%v", err)
		return nil, 0, err
	}
	return resources, count, nil
}

type CreateResourceRequest struct {
	entity.Resource
}

func (svc *ResourceService) Create(ctx context.Context, param *CreateResourceRequest) error {
	svc.logger.Debugf("resource create request:%+v", param)
	if err := svc.store.Resources().Create(ctx, &param.Resource); err != nil {
		svc.logger.Errorf("resource store create err:%v", err)
		return err
	}
	return nil
}

type UpdateResourceRequest struct {
	entity.Resource
}

func (svc *ResourceService) Update(ctx context.Context, param *UpdateResourceRequest) error {
	svc.logger.Debugf("resource update request:%+v", param)
	err := svc.store.Resources().Update(ctx, &param.Resource)
	if err != nil {
		svc.logger.Errorf("resource store update err:%v", err)
		return err
	}
	return nil
}

func (svc *ResourceService) Delete(ctx context.Context, id int) error {
	svc.logger.Debugf("resource delete request:%d", id)
	if err := svc.store.Resources().Delete(ctx, id); err != nil {
		svc.logger.Errorf("resource store delete err:%v", err)
		return err
	}
	return nil
}
