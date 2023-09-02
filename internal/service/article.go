package service

import (
	"context"
	"errors"

	log "github.com/sirupsen/logrus"

	"github.com/hexiaopi/blog-service/internal/cache"
	"github.com/hexiaopi/blog-service/internal/model"
	"github.com/hexiaopi/blog-service/internal/store"
)

type ArticleSrv interface {
	Get(ctx context.Context, request *ArticleRequest) (*model.Article, error)
	List(ctx context.Context, param *ArticleListRequest) ([]model.Article, int64, error)
	Create(ctx context.Context, param *CreateArticleRequest) error
	Update(ctx context.Context, param *UpdateArticleRequest) error
	Delete(ctx context.Context, id int) error
}

type ArticleService struct {
	store store.Factory
	cache cache.Factory
}

var _ ArticleSrv = (*ArticleService)(nil)

func NewArticleService(factory store.Factory, cache cache.Factory) *ArticleService {
	return &ArticleService{
		store: factory,
		cache: cache,
	}
}

type ArticleRequest struct {
	model.OneOption
}

func (svc *ArticleService) Get(ctx context.Context, request *ArticleRequest) (*model.Article, error) {
	article, err := svc.store.Articles().Get(ctx, request.Id)
	if err != nil {
		log.Errorf("article store get err:%v", err)
		return nil, err
	}
	return article, nil
}

type ArticleListRequest struct {
	model.ListOption
}

func (svc *ArticleService) List(ctx context.Context, param *ArticleListRequest) ([]model.Article, int64, error) {
	articles, err := svc.store.Articles().List(ctx, &param.ListOption)
	if err != nil {
		log.Errorf("article store list err:%v", err)
		return nil, 0, err
	}
	var count int64
	var set bool
	if param.ListOption.Name == "" {
		count, err = svc.cache.Articles().GetCount(ctx)
		if err != nil {
			if errors.Is(err, cache.ErrNotFound) {
				set = true
			} else {
				log.Errorf("cache article count get err:%v", err)
				return nil, 0, err
			}
		} else {
			return articles, count, nil
		}
	}
	count, err = svc.store.Articles().Count(ctx, &param.ListOption)
	if err != nil {
		log.Errorf("article store count err:%v", err)
		return nil, 0, err
	}
	if set {
		if err := svc.cache.Articles().SetCount(ctx, count); err != nil {
			log.Errorf("article cache set count err:%v", err)
		}
	}
	return articles, count, nil
}

type CreateArticleRequest struct {
	model.Article
}

func (svc *ArticleService) Create(ctx context.Context, param *CreateArticleRequest) error {
	if err := svc.store.Articles().Create(ctx, &param.Article); err != nil {
		log.Errorf("article store create err:%v", err)
		return err
	}
	return nil
}

type UpdateArticleRequest struct {
	model.Article
}

func (svc *ArticleService) Update(ctx context.Context, param *UpdateArticleRequest) error {
	err := svc.store.Articles().Update(ctx, &param.Article)
	if err != nil {
		log.Errorf("article store update err:%v", err)
		return err
	}
	//todo update tag
	return nil
}

func (svc *ArticleService) Delete(ctx context.Context, id int) error {
	if err := svc.store.Articles().Delete(ctx, id); err != nil {
		log.Errorf("article store delete err:%v", err)
		return err
	}
	//todo delete tag
	return nil
}
