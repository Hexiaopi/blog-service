package service

import (
	"context"

	"github.com/hexiaopi/blog-service/internal/entity"
	"github.com/hexiaopi/blog-service/internal/store"
	log "github.com/hexiaopi/blog-service/pkg/logger"
)

type OperationSrv interface {
	List(ctx context.Context, param *OperationListRequest) ([]entity.OperationLog, int64, error)
	Create(ctx context.Context, param *CreateOperationRequest) error
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
	entity.ListOption
}

func (svc *OperationService) List(ctx context.Context, param *OperationListRequest) ([]entity.OperationLog, int64, error) {
	svc.logger.Debugf("operation list request:%+v", param)
	if param.UserName != "" {
		user, err := svc.store.Users().GetByName(ctx, param.UserName)
		if err != nil {
			svc.logger.Errorf("user store get err:%v", err)
			return nil, 0, err
		}
		param.ListOption.UserId = user.ID
	}
	total, err := svc.store.Operations().Count(ctx, &param.ListOption)
	if err != nil {
		svc.logger.Errorf("operation store count err:%v", err)
		return nil, 0, err
	}
	if total == 0 {
		return nil, 0, nil
	}
	logs, err := svc.store.Operations().List(ctx, &param.ListOption)
	if err != nil {
		svc.logger.Errorf("operate store list err:%v", err)
		return nil, 0, err
	}
	result := make([]entity.OperationLog, 0, len(logs))
	for _, log := range logs {
		user, err := svc.store.Users().GetById(ctx, log.User.ID)
		if err != nil {
			svc.logger.Errorf("operate user query err:%v", err)
			return nil, 0, err
		}
		if user == nil {
			continue
		}
		r := entity.ToEntityOperation(&log)
		r.User = *user
		result = append(result, *r)
	}
	return result, total, nil
}

type CreateOperationRequest struct {
	entity.OperationLog
}

func (svc *OperationService) Create(ctx context.Context, param *CreateOperationRequest) error {
	svc.logger.Debugf("operation create request:%+v", param)
	if err := svc.store.Operations().Create(ctx, &param.OperationLog); err != nil {
		svc.logger.Errorf("operation store create err:%v", err)
		return err
	}
	return nil
}
