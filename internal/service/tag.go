package service

import (
	"context"

	"github.com/hexiaopi/blog-service/internal/model"
	"github.com/hexiaopi/blog-service/internal/store"
)

type TagSrv interface {
	List(ctx context.Context, param *TagListRequest) ([]model.Tag, int64, error)
	Create(ctx context.Context, param *CreateTagRequest) error
	Update(ctx context.Context, param *UpdateTagRequest) error
	Delete(ctx context.Context, param *DeleteTagRequest) error
}

type TagService struct {
	store store.Factory
}

var _ TagSrv = (*TagService)(nil)

func NewTagService(factory store.Factory) *TagService {
	return &TagService{
		store: factory,
	}
}

type TagListRequest struct {
	model.ListOption
}

func (svc *TagService) List(ctx context.Context, param *TagListRequest) ([]model.Tag, int64, error) {
	tags, total, err := svc.store.Tags().List(ctx, &param.ListOption)
	if err != nil {
		return nil, 0, err
	}
	return tags, total, nil
}

type CreateTagRequest struct {
	model.Tag
}

func (svc *TagService) Create(ctx context.Context, param *CreateTagRequest) error {
	if err := svc.store.Tags().Create(ctx, &param.Tag); err != nil {
		return err
	}
	return nil
}

type UpdateTagRequest struct {
	model.Tag
}

func (svc *TagService) Update(ctx context.Context, param *UpdateTagRequest) error {
	if err := svc.store.Tags().Update(ctx, &param.Tag); err != nil {
		return err
	}
	return nil
}

type DeleteTagRequest struct {
	model.OneOption
}

func (svc *TagService) Delete(ctx context.Context, param *DeleteTagRequest) error {
	return svc.store.Tags().Delete(ctx, param.Id)
}
