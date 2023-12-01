package service

import (
	"context"
	"errors"

	"github.com/hexiaopi/blog-service/internal/cache"
	"github.com/hexiaopi/blog-service/internal/model"
	"github.com/hexiaopi/blog-service/internal/store"
	log "github.com/hexiaopi/blog-service/pkg/logger"
)

type ArticleSrv interface {
	Get(ctx context.Context, request *ArticleRequest) (*model.Article, error)
	List(ctx context.Context, param *ArticleListRequest) ([]model.Article, int64, error)
	Create(ctx context.Context, param *CreateArticleRequest) error
	Update(ctx context.Context, param *UpdateArticleRequest) error
	Delete(ctx context.Context, id int) error
}

type ArticleService struct {
	store  store.Factory
	cache  cache.Factory
	logger log.Logger
}

var _ ArticleSrv = (*ArticleService)(nil)

func NewArticleService(factory store.Factory, cache cache.Factory, logger log.Logger) *ArticleService {
	return &ArticleService{
		store:  factory,
		cache:  cache,
		logger: logger,
	}
}

type ArticleRequest struct {
	model.OneOption
}

func (svc *ArticleService) Get(ctx context.Context, request *ArticleRequest) (*model.Article, error) {
	svc.logger.Debugf("article get request:%+v", request)
	article, err := svc.store.Articles().Get(ctx, request.Id)
	if err != nil {
		svc.logger.Errorf("article store get err:%v", err)
		return nil, err
	}
	return article, nil
}

type ArticleListRequest struct {
	model.ListOption
}

func (svc *ArticleService) List(ctx context.Context, param *ArticleListRequest) ([]model.Article, int64, error) {
	svc.logger.Debugf("article list request:%+v", param)
	articles, err := svc.store.Articles().List(ctx, &param.ListOption)
	if err != nil {
		svc.logger.Errorf("article store list err:%v", err)
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
				svc.logger.Errorf("cache article count get err:%v", err)
				return nil, 0, err
			}
		} else {
			return articles, count, nil
		}
	}
	count, err = svc.store.Articles().Count(ctx, &param.ListOption)
	if err != nil {
		svc.logger.Errorf("article store count err:%v", err)
		return nil, 0, err
	}
	if set {
		if err := svc.cache.Articles().SetCount(ctx, count); err != nil {
			svc.logger.Errorf("article cache set count err:%v", err)
		}
	}
	return articles, count, nil
}

type CreateArticleRequest struct {
	model.Article
}

func (svc *ArticleService) Create(ctx context.Context, param *CreateArticleRequest) error {
	svc.logger.Debugf("article create request:%+v", param)
	if err := svc.store.Articles().Create(ctx, &param.Article); err != nil {
		svc.logger.Errorf("article store create err:%v", err)
		return err
	}
	return nil
}

type UpdateArticleRequest struct {
	model.Article
}

func (svc *ArticleService) Update(ctx context.Context, param *UpdateArticleRequest) error {
	svc.logger.Debugf("article update request:%+v", param)
	err := svc.store.Articles().Update(ctx, &param.Article)
	if err != nil {
		svc.logger.Errorf("article store update err:%v", err)
		return err
	}
	//todo update tag
	return nil
}

func (svc *ArticleService) Delete(ctx context.Context, id int) error {
	svc.logger.Debugf("article delete id:%d", id)
	if err := svc.store.Articles().Delete(ctx, id); err != nil {
		svc.logger.Errorf("article store delete err:%v", err)
		return err
	}
	//todo delete tag
	return nil
}
