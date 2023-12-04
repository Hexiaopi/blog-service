package service

import (
	"context"

	"github.com/hexiaopi/blog-service/internal/entity"
	"github.com/hexiaopi/blog-service/internal/store"
	log "github.com/hexiaopi/blog-service/pkg/logger"
)

type SystemSrv interface {
	Get(ctx context.Context, param *SystemGetRequest) (*entity.Config, error)
}

type SystemService struct {
	store  store.Factory
	logger log.Logger
}

var _ SystemSrv = (*SystemService)(nil)

func NewSystemService(factory store.Factory, logger log.Logger) *SystemService {
	return &SystemService{
		store:  factory,
		logger: logger,
	}
}

type SystemGetRequest struct {
	entity.OneOption
}

func (svc *SystemService) Get(ctx context.Context, param *SystemGetRequest) (*entity.Config, error) {
	svc.logger.Debugf("system get request:%+v", param)
	sc, err := svc.store.Systems().Get(ctx, param.Name)
	if err != nil {
		svc.logger.Errorf("system get:%s error:%v", param.Name, err)
		return nil, err
	}
	return sc, nil
}
