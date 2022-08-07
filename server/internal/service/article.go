package service

import (
	"github.com/hexiaopi/blog-service/internal/app"
	"github.com/hexiaopi/blog-service/internal/dao"
	"github.com/hexiaopi/blog-service/internal/model"
)

type ArticleRequest struct {
	ID    uint32 `json:"id"`
	State uint8  `json:"state"`
}

type Article struct {
	ID            uint32     `json:"id"`
	Title         string     `json:"title"`
	Desc          string     `json:"desc"`
	Content       string     `json:"content"`
	CoverImageUrl string     `json:"cover_image_url"`
	State         uint8      `json:"state"`
	Tag           *model.Tag `json:"tag"`
}

func (svc *Service) GetArticle(param *ArticleRequest) (*Article, error) {
	article, err := svc.dao.GetArticle(param.ID, param.State)
	if err != nil {
		return nil, err
	}

	articleTag, err := svc.dao.GetArticleTagByAID(article.ID)
	if err != nil {
		return nil, err
	}

	tag, err := svc.dao.GetTag(articleTag.TagID, model.STATE_OPEN)
	if err != nil {
		return nil, err
	}

	return &Article{
		ID:            article.ID,
		Title:         article.Title,
		Desc:          article.Desc,
		Content:       article.Content,
		CoverImageUrl: article.CoverImageUrl,
		State:         article.State,
		Tag:           &tag,
	}, nil
}

type ArticleListRequest struct {
	TagID uint32 `json:"tag_id"`
	State uint8  `json:"state"`
}

func (svc *Service) ListArticle(param *ArticleListRequest, page *app.Page) ([]*Article, int, error) {
	articleTotal, err := svc.dao.CountArticles(param.State)
	if err != nil {
		return nil, 0, err
	}

	articles, err := svc.dao.ListArticles(param.State, page.PageNum, page.PageSize)
	if err != nil {
		return nil, 0, err
	}

	var articleList []*Article
	for _, article := range articles {
		articleList = append(articleList, &Article{
			ID:            article.ArticleID,
			Title:         article.ArticleTitle,
			Desc:          article.ArticleDesc,
			Content:       article.ArticleDesc,
			CoverImageUrl: article.CoverImageUrl,
			Tag:           &model.Tag{ID: article.TagID, Name: article.TagName},
		})
	}
	return articleList, articleTotal, nil
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

func (svc *Service) CreateArticle(param *CreateArticleRequest) error {
	article, err := svc.dao.CreateArticle(&dao.Article{
		TagID:         param.TagID,
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		CreatedBy:     param.CreateBy,
		State:         param.State,
	})
	if err != nil {
		return err
	}

	err = svc.dao.CreateArticleTag(article.ID, param.TagID, param.CreateBy)
	if err != nil {
		return err
	}
	return nil
}

type UpdateArticleRequest struct {
	ID            uint32 `json:"id"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	ModifiedBy    string `json:"modified_by"`
	State         uint8  `json:"state"`
	TagID         uint32 `json:"tag_id"`
}

func (svc *Service) UpdateArticle(param *UpdateArticleRequest) error {
	err := svc.dao.UpdateArticle(&dao.Article{
		ID:            param.ID,
		TagID:         param.TagID,
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		ModifiedBy:    param.ModifiedBy,
		State:         param.State,
	})
	if err != nil {
		return err
	}
	err = svc.dao.UpdateArticleTag(param.ID, param.TagID, param.ModifiedBy)
	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) DeleteArticle(articleID uint32) error {
	if err := svc.dao.DeleteArticle(articleID); err != nil {
		return err
	}
	if err := svc.dao.DeleteArticleTag(articleID); err != nil {
		return err
	}
	return nil
}
