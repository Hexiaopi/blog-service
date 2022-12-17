package service

import "github.com/hexiaopi/blog-service/internal/store"

type Service interface {
	Articles() ArticleSrv
	Tags() TagSrv
	Users() UserSrv
}

type service struct {
	store store.Factory
}

func (s *service) Articles() ArticleSrv {
	return NewArticleService(s.store)
}

func (s *service) Tags() TagSrv {
	return NewTagService(s.store)
}

func (s *service) Users() UserSrv {
	return NewUserService(s.store)
}

func NewService(store store.Factory) Service {
	return &service{
		store: store,
	}
}
