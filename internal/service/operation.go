package service

import (
	"context"
	"log"

	"github.com/hexiaopi/blog-service/internal/model"
	"github.com/hexiaopi/blog-service/internal/store"
)

type OperationSrv interface {
	List(ctx context.Context, param *OperationListRequest) ([]model.OperationLog, int64, error)
	Create(ctx context.Context, param *CreateOperationRequest) error
	Update(ctx context.Context, param *UpdateOperationRequest) error
	Delete(ctx context.Context, param *DeleteOperationRequest) error
}

type OperationService struct {
	store store.Factory
}

var _ OperationSrv = (*OperationService)(nil)

func NewOperationService(factory store.Factory) *OperationService {
	return &OperationService{
		store: factory,
	}
}

type OperationListRequest struct {
	model.ListOption
}

func (svc *OperationService) List(ctx context.Context, param *OperationListRequest) ([]model.OperationLog, int64, error) {
	logs, err := svc.store.Operations().List(ctx, &param.ListOption)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}
	total, err := svc.store.Operations().Count(ctx, &param.ListOption)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}
	return logs, total, nil
}

type CreateOperationRequest struct {
	model.OperationLog
}

func (svc *OperationService) Create(ctx context.Context, param *CreateOperationRequest) error {
	if err := svc.store.Operations().Create(ctx, &param.OperationLog); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

type UpdateOperationRequest struct {
	model.OperationLog
}

func (svc *OperationService) Update(ctx context.Context, param *UpdateOperationRequest) error {
	if err := svc.store.Operations().Update(ctx, &param.OperationLog); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

type DeleteOperationRequest struct {
	model.OneOption
}

func (svc *OperationService) Delete(ctx context.Context, param *DeleteOperationRequest) error {
	return svc.store.Operations().Delete(ctx, param.Id)
}
