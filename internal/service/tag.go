package service

import (
	"context"

	"github.com/hexiaopi/blog-service/internal/entity"
	"github.com/hexiaopi/blog-service/internal/store"
)

type TagService struct {
	store store.Factory
}

func NewTagService(factory store.Factory) TagService {
	return TagService{
		store: factory,
	}
}

type TagListRequest struct {
	entity.ListOption
}

func (svc *TagService) List(ctx context.Context, param *TagListRequest) ([]entity.Tag, int64, error) {
	opt := entity.ListOption{
		State:    param.State,
		PageSize: param.PageSize,
		PageNum:  param.PageNum,
	}
	tags, total, err := svc.store.Tags().List(ctx, &opt)
	if err != nil {
		return nil, 0, err
	}
	return tags, total, nil
}

type CreateTagRequest struct {
	Desc string
	entity.OneOption
}

func (svc *TagService) Create(ctx context.Context, param *CreateTagRequest) error {
	tag := entity.Tag{
		Name:     param.Name,
		Desc:     param.Desc,
		State:    param.State,
		Operator: param.Operator,
	}
	if err := svc.store.Tags().Create(ctx, &tag); err != nil {
		return err
	}
	return nil
}

type UpdateTagRequest struct {
	Desc string
	entity.OneOption
}

func (svc *TagService) Update(ctx context.Context, param *UpdateTagRequest) error {
	tag := entity.Tag{
		Name:     param.Name,
		Desc:     param.Desc,
		State:    param.State,
		Operator: param.Operator,
	}
	if err := svc.store.Tags().Update(ctx, &tag); err != nil {
		return err
	}
	return nil
}

type DeleteTagRequest struct {
	entity.OneOption
}

func (svc *TagService) Delete(ctx context.Context, param *DeleteTagRequest) error {
	return svc.store.Tags().Delete(ctx, param.Id)
}
