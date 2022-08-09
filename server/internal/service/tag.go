package service

import (
	"context"

	"github.com/hexiaopi/blog-service/internal/app"
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
	Name  string
	State uint8
}

func (svc *TagService) List(ctx context.Context, param *TagListRequest, page *app.Page) ([]*entity.Tag, int64, error) {
	opt := entity.ListOption{
		State: param.State,
		Page:  page,
	}
	tags, total, err := svc.store.Tags().List(ctx, &opt)
	if err != nil {
		return nil, 0, err
	}
	result := make([]*entity.Tag, len(tags))
	for i, tag := range tags {
		result[i] = &entity.Tag{
			Id:        tag.ID,
			Name:      tag.Name,
			State:     tag.State,
			CreatedBy: tag.CreatedBy,
		}
	}
	return result, total, nil
}

type CreateTagRequest struct {
	Name      string
	CreatedBy string
	State     uint8
}

func (svc *TagService) Create(ctx context.Context, param *CreateTagRequest) error {
	tag := entity.Tag{
		Name:      param.Name,
		State:     param.State,
		CreatedBy: param.CreatedBy,
	}
	if err := svc.store.Tags().Create(ctx, &tag); err != nil {
		return err
	}
	return nil
}

type UpdateTagRequest struct {
	ID         uint32
	Name       string
	State      uint8
	ModifiedBy string
}

func (svc *TagService) Update(ctx context.Context, param *UpdateTagRequest) error {
	tag := entity.Tag{
		Name:       param.Name,
		State:      param.State,
		ModifiedBy: param.ModifiedBy,
	}
	if err := svc.store.Tags().Update(ctx, &tag); err != nil {
		return err
	}
	return nil
}

type DeleteTagRequest struct {
	ID int
}

func (svc *TagService) Delete(ctx context.Context, param *DeleteTagRequest) error {
	return svc.store.Tags().Delete(ctx, param.ID)
}
