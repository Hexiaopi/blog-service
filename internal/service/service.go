package service

import (
	"github.com/hexiaopi/blog-service/internal/cache"
	"github.com/hexiaopi/blog-service/internal/store"
)

type Service interface {
	Articles() ArticleSrv
	Tags() TagSrv
	Users() UserSrv
}

type service struct {
	store store.Factory
	cache cache.Factory
}

func (s *service) Articles() ArticleSrv {
	return NewArticleService(s.store, s.cache)
}

func (s *service) Tags() TagSrv {
	return NewTagService(s.store)
}

func (s *service) Users() UserSrv {
	return NewUserService(s.store)
}

func NewService(store store.Factory, cache cache.Factory) Service {
	return &service{
		store: store,
		cache: cache,
	}
}
