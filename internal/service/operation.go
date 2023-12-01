package service

import (
	"context"

	"github.com/hexiaopi/blog-service/internal/model"
	"github.com/hexiaopi/blog-service/internal/store"
	log "github.com/hexiaopi/blog-service/pkg/logger"
)

type OperationSrv interface {
	List(ctx context.Context, param *OperationListRequest) ([]model.OperationLog, int64, error)
	Create(ctx context.Context, param *CreateOperationRequest) error
	Update(ctx context.Context, param *UpdateOperationRequest) error
	Delete(ctx context.Context, param *DeleteOperationRequest) error
}

type OperationService struct {
	store  store.Factory
	logger log.Logger
}

var _ OperationSrv = (*OperationService)(nil)

func NewOperationService(factory store.Factory, logger log.Logger) *OperationService {
	return &OperationService{
		store:  factory,
		logger: logger,
	}
}

type OperationListRequest struct {
	UserName string `json:"username"`
	model.ListOption
}

func (svc *OperationService) List(ctx context.Context, param *OperationListRequest) ([]model.OperationLog, int64, error) {
	svc.logger.Debugf("operation list request:%+v", param)
	if param.UserName != "" {
		user, err := svc.store.Users().Get(ctx, param.UserName)
		if err != nil {
			svc.logger.Errorf("user store get err:%v", err)
			return nil, 0, err
		}
		param.ListOption.UserId = user.ID
	}
	logs, err := svc.store.Operations().List(ctx, &param.ListOption)
	if err != nil {
		svc.logger.Errorf("operate store list err:%v", err)
		return nil, 0, err
	}
	total, err := svc.store.Operations().Count(ctx, &param.ListOption)
	if err != nil {
		svc.logger.Errorf("operation store count err:%v", err)
		return nil, 0, err
	}
	return logs, total, nil
}

type CreateOperationRequest struct {
	model.OperationLog
}

func (svc *OperationService) Create(ctx context.Context, param *CreateOperationRequest) error {
	svc.logger.Debugf("operation create request:%+v", param)
	if err := svc.store.Operations().Create(ctx, &param.OperationLog); err != nil {
		svc.logger.Errorf("operation store create err:%v", err)
		return err
	}
	return nil
}

type UpdateOperationRequest struct {
	model.OperationLog
}

func (svc *OperationService) Update(ctx context.Context, param *UpdateOperationRequest) error {
	svc.logger.Debugf("operation update request:%+v", param)
	if err := svc.store.Operations().Update(ctx, &param.OperationLog); err != nil {
		svc.logger.Errorf("operation store update err:%v", err)
		return err
	}
	return nil
}

type DeleteOperationRequest struct {
	model.OneOption
}

func (svc *OperationService) Delete(ctx context.Context, param *DeleteOperationRequest) error {
	svc.logger.Debugf("operation delete request:%+v", param)
	if err := svc.store.Operations().Delete(ctx, param.Id); err != nil {
		svc.logger.Errorf("operation store delete err:%v", err)
		return err
	}
	return nil
}
