package service

import (
	"context"

	"github.com/hexiaopi/blog-service/internal/entity"
	"github.com/hexiaopi/blog-service/internal/store"
)

type ArticleService struct {
	store store.Factory
}

func NewArticleService(factory store.Factory) ArticleService {
	return ArticleService{
		store: factory,
	}
}

type ArticleRequest struct {
	entity.OneOption
}

func (svc *ArticleService) Get(ctx context.Context, request *ArticleRequest) (*entity.Article, error) {
	article, err := svc.store.Articles().Get(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return article, nil
}

type ArticleListRequest struct {
	entity.ListOption
}

func (svc *ArticleService) List(ctx context.Context, param *ArticleListRequest) ([]entity.Article, int64, error) {
	articles, total, err := svc.store.Articles().List(ctx, &param.ListOption)
	if err != nil {
		return nil, 0, err
	}
	return articles, total, nil
}

type CreateArticleRequest struct {
	entity.Article
}

func (svc *ArticleService) Create(ctx context.Context, param *CreateArticleRequest) error {
	article := entity.Article{
		Name:          param.Name,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State:         param.State,
		Operator:      param.Operator,
	}
	if err := svc.store.Articles().Create(ctx, &article); err != nil {
		return err
	}
	return nil
}

type UpdateArticleRequest struct {
	entity.Article
	Tags []entity.Tag `json:"tags"`
}

func (svc *ArticleService) Update(ctx context.Context, param *UpdateArticleRequest) error {
	article := entity.Article{
		Id:            param.Id,
		Name:          param.Name,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State:         param.State,
		Operator:      param.Operator,
	}
	err := svc.store.Articles().Update(ctx, &article)
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
