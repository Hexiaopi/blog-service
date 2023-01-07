package service

import (
	"context"

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
}

var _ ArticleSrv = (*ArticleService)(nil)

func NewArticleService(factory store.Factory) *ArticleService {
	return &ArticleService{
		store: factory,
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
	articles, total, err := svc.store.Articles().List(ctx, &param.ListOption)
	if err != nil {
		return nil, 0, err
	}
	return articles, total, nil
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
