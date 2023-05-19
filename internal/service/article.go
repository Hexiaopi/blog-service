package service

import (
	"context"
	"errors"
	"log"

	"github.com/redis/go-redis/v9"

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
		return nil, 0, err
	}
	var count int64
	var set bool
	if param.ListOption.Name == "" {
		count, err = svc.cache.Articles().GetCount(ctx)
		if err != nil {
			if errors.Is(err, redis.Nil) {
				set = true
			} else {
				log.Println(err)
				return nil, 0, err
			}
		} else {
			return articles, count, nil
		}
	}
	count, err = svc.store.Articles().Count(ctx, &param.ListOption)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}
	if set {
		if err := svc.cache.Articles().SetCount(ctx, count); err != nil {
			log.Println(err)
		}
	}
	return articles, count, nil
}

type CreateArticleRequest struct {
	model.Article
}

func (svc *ArticleService) Create(ctx context.Context, param *CreateArticleRequest) error {
	if err := svc.store.Articles().Create(ctx, &param.Article); err != nil {
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
		return err
	}
	//todo update tag
	return nil
}

func (svc *ArticleService) Delete(ctx context.Context, id int) error {
	if err := svc.store.Articles().Delete(ctx, id); err != nil {
		return err
	}
	//todo delete tag
	return nil
}
