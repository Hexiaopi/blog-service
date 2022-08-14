package service

import (
	"context"

	"github.com/hexiaopi/blog-service/internal/app"
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
	ID    int   `json:"id"`
	State uint8 `json:"state"`
}

func (svc *ArticleService) Get(ctx context.Context, request *ArticleRequest) (*entity.Article, error) {
	article, err := svc.store.Articles().Get(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	// tag, err := svc.store.Tags().GetByArticle(ctx, request.ID)
	// if err != nil {
	// 	return nil, err
	// }

	result := &entity.Article{
		ID:            article.ID,
		Title:         article.Title,
		Desc:          article.Desc,
		Content:       article.Content,
		CoverImageUrl: article.CoverImageUrl,
		State:         article.State,
	}
	//result.Tags = append(result.Tags, entity.Tag{Id: int(tag.ID), Name: tag.Name})
	return result, nil
}

type ArticleListRequest struct {
	TagID uint32 `json:"tag_id"`
	State uint8  `json:"state"`
}

func (svc *ArticleService) List(ctx context.Context, param *ArticleListRequest, page *app.Page) ([]*entity.Article, int64, error) {
	opt := entity.ListOption{
		State: param.State,
		Page:  page,
	}
	articles, total, err := svc.store.Articles().List(ctx, &opt)
	if err != nil {
		return nil, 0, err
	}

	var articleList []*entity.Article
	for _, article := range articles {
		// tag, err := svc.store.Tags().GetByArticle(ctx, article.ID)
		// if err != nil {
		// 	return nil, 0, err
		// }
		a := &entity.Article{
			ID:            article.ID,
			Title:         article.Title,
			Desc:          article.Desc,
			Content:       article.Content,
			CoverImageUrl: article.CoverImageUrl,
		}
		//a.Tags = append(a.Tags, entity.Tag{Id: int(tag.ID), Name: tag.Name})
		articleList = append(articleList, a)
	}
	return articleList, total, nil
}

type CreateArticleRequest struct {
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	CreateBy      string `json:"create_by"`
	State         uint8  `json:"state"`
	TagID         uint32 `json:"tag_id"`
}

func (svc *ArticleService) Create(ctx context.Context, request *CreateArticleRequest) error {
	article := entity.Article{
		Title:         request.Title,
		Desc:          request.Desc,
		Content:       request.Content,
		CoverImageUrl: request.CoverImageUrl,
		CreatedBy:     request.CreateBy,
		State:         request.State,
	}
	if err := svc.store.Articles().Create(ctx, &article); err != nil {
		return err
	}
	//todo create article_tag
	return nil
}

type UpdateArticleRequest struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	ModifiedBy    string `json:"modified_by"`
	State         uint8  `json:"state"`
	TagID         int    `json:"tag_id"`
}

func (svc *ArticleService) Update(ctx context.Context, param *UpdateArticleRequest) error {
	article := entity.Article{
		ID:            param.ID,
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		ModifiedBy:    param.ModifiedBy,
		State:         param.State,
	}
	err := svc.store.Articles().Update(ctx, &article)
	if err != nil {
		return err
	}
	//todo update tag
	return nil
}

func (svc *ArticleService) Delete(ctx context.Context, articleId int) error {
	if err := svc.store.Articles().Delete(ctx, articleId); err != nil {
		return err
	}
	//todo delete tag
	return nil
}
