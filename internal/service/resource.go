package service

import (
	"context"
	"encoding/base64"
	"log"

	"github.com/hexiaopi/blog-service/internal/model"
	"github.com/hexiaopi/blog-service/internal/store"
)

type ResourceSrv interface {
	Get(ctx context.Context, request *ResourceRequest) (*model.Resource, error)
	List(ctx context.Context, param *ResourceListRequest) ([]model.Resource, int64, error)
	Create(ctx context.Context, param *CreateResourceRequest) error
	Update(ctx context.Context, param *UpdateResourceRequest) error
	Delete(ctx context.Context, id int) error
}

type ResourceService struct {
	store store.Factory
}

var _ ResourceSrv = (*ResourceService)(nil)

func NewResourceService(factory store.Factory) *ResourceService {
	return &ResourceService{
		store: factory,
	}
}

type ResourceRequest struct {
	model.OneOption
}

func (svc *ResourceService) Get(ctx context.Context, request *ResourceRequest) (*model.Resource, error) {
	resource, err := svc.store.Resources().Get(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	resource.Base64 = "data:image/png;base64," + base64.StdEncoding.EncodeToString(resource.Blob)
	return resource, nil
}

type ResourceListRequest struct {
	model.ListOption
}

func (svc *ResourceService) List(ctx context.Context, param *ResourceListRequest) ([]model.Resource, int64, error) {
	resources, err := svc.store.Resources().List(ctx, &param.ListOption)
	if err != nil {
		return nil, 0, err
	}
	for i := range resources {
		resources[i].Base64 = "data:image/png;base64," + base64.RawStdEncoding.EncodeToString(resources[i].Blob)
		resources[i].Blob = nil
	}
	count, err := svc.store.Resources().Count(ctx, &param.ListOption)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}
	return resources, count, nil
}

type CreateResourceRequest struct {
	model.Resource
}

func (svc *ResourceService) Create(ctx context.Context, param *CreateResourceRequest) error {
	if err := svc.store.Resources().Create(ctx, &param.Resource); err != nil {
		return err
	}
	return nil
}

type UpdateResourceRequest struct {
	model.Resource
}

func (svc *ResourceService) Update(ctx context.Context, param *UpdateResourceRequest) error {
	err := svc.store.Resources().Update(ctx, &param.Resource)
	if err != nil {
		return err
	}
	return nil
}

func (svc *ResourceService) Delete(ctx context.Context, id int) error {
	if err := svc.store.Resources().Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
