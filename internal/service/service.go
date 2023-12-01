package service

import (
	"github.com/hexiaopi/blog-service/internal/cache"
	"github.com/hexiaopi/blog-service/internal/store"
	log "github.com/hexiaopi/blog-service/pkg/logger"
)

type Service interface {
	Articles() ArticleSrv
	Tags() TagSrv
	Users() UserSrv
	Roles() RoleSrv
	Systems() SystemSrv
	Resources() ResourceSrv
	Operations() OperationSrv
}

type service struct {
	store  store.Factory
	cache  cache.Factory
	logger log.Logger
}

func (s *service) Articles() ArticleSrv {
	return NewArticleService(s.store, s.cache, s.logger)
}

func (s *service) Tags() TagSrv {
	return NewTagService(s.store, s.logger)
}

func (s *service) Users() UserSrv {
	return NewUserService(s.store, s.logger)
}

func (s *service) Roles() RoleSrv {
	return NewRoleService(s.store, s.logger)
}

func (s *service) Systems() SystemSrv {
	return NewSystemService(s.store, s.logger)
}

func (s *service) Resources() ResourceSrv {
	return NewResourceService(s.store, s.logger)
}

func (s *service) Operations() OperationSrv {
	return NewOperationService(s.store, s.logger)
}

func NewService(store store.Factory, cache cache.Factory, logger log.Logger) Service {
	return &service{
		store:  store,
		cache:  cache,
		logger: logger,
	}
}
