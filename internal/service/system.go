package service

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/hexiaopi/blog-service/internal/model"
	"github.com/hexiaopi/blog-service/internal/store"
)

type SystemSrv interface {
	Get(ctx context.Context, param *SystemGetRequest) (*model.SystemConfig, error)
}

type SystemService struct {
	store store.Factory
}

var _ SystemSrv = (*SystemService)(nil)

func NewSystemService(factory store.Factory) *SystemService {
	return &SystemService{
		store: factory,
	}
}

type SystemGetRequest struct {
	model.OneOption
}

func (srv *SystemService) Get(ctx context.Context, param *SystemGetRequest) (*model.SystemConfig, error) {
	sc, err := srv.store.Systems().Get(ctx, param.Name)
	if err != nil {
		log.Errorf("system get:%s error:%v", param.Name, err)
		return nil, err
	}
	return sc, nil
}
