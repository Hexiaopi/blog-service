package service

import (
	"context"

	"github.com/hexiaopi/blog-service/internal/model"
	"github.com/hexiaopi/blog-service/internal/store"
	log "github.com/hexiaopi/blog-service/pkg/logger"
)

type TagSrv interface {
	List(ctx context.Context, param *TagListRequest) ([]model.Tag, int64, error)
	Create(ctx context.Context, param *CreateTagRequest) error
	Update(ctx context.Context, param *UpdateTagRequest) error
	Delete(ctx context.Context, param *DeleteTagRequest) error
}

type TagService struct {
	store  store.Factory
	logger log.Logger
}

var _ TagSrv = (*TagService)(nil)

func NewTagService(factory store.Factory, logger log.Logger) *TagService {
	return &TagService{
		store:  factory,
		logger: logger,
	}
}

type TagListRequest struct {
	model.ListOption
}

func (svc *TagService) List(ctx context.Context, param *TagListRequest) ([]model.Tag, int64, error) {
	svc.logger.Debugf("tag list request:%+v", param)
	tags, err := svc.store.Tags().List(ctx, &param.ListOption)
	if err != nil {
		svc.logger.Errorf("tag store list err:%v", err)
		return nil, 0, err
	}
	total, err := svc.store.Tags().Count(ctx, &param.ListOption)
	if err != nil {
		svc.logger.Errorf("tag store count err:%v", err)
		return nil, 0, err
	}
	return tags, total, nil
}

type CreateTagRequest struct {
	model.Tag
}

func (svc *TagService) Create(ctx context.Context, param *CreateTagRequest) error {
	svc.logger.Debugf("tag create request:%+v", param)
	if err := svc.store.Tags().Create(ctx, &param.Tag); err != nil {
		svc.logger.Errorf("tag store create err:%v", err)
		return err
	}
	return nil
}

type UpdateTagRequest struct {
	model.Tag
}

func (svc *TagService) Update(ctx context.Context, param *UpdateTagRequest) error {
	svc.logger.Debugf("tag update request:%+v", param)
	if err := svc.store.Tags().Update(ctx, &param.Tag); err != nil {
		svc.logger.Errorf("tag store update err:%v", err)
		return err
	}
	return nil
}

type DeleteTagRequest struct {
	model.OneOption
}

func (svc *TagService) Delete(ctx context.Context, param *DeleteTagRequest) error {
	svc.logger.Debugf("tag delete request:%+v", param)
	if err := svc.store.Tags().Delete(ctx, param.Id); err != nil {
		svc.logger.Errorf("tag store delete err:%v", err)
		return err
	}
	return nil
}
