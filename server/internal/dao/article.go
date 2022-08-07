package dao

import (
	"github.com/hexiaopi/blog-service/internal/app"
	"github.com/hexiaopi/blog-service/internal/model"
)

type Article struct {
	ID            uint32 `json:"id"`
	TagID         uint32 `json:"tag_id"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         uint8  `json:"state"`
}

func (d *Dao) CreateArticle(param *Article) (*model.Article, error) {
	article := model.Article{
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State:         param.State,
		CreatedBy:     param.CreatedBy,
	}
	return article.Create(d.engine)
}

func (d *Dao) UpdateArticle(param *Article) error {
	article := model.Article{ID: param.ID}
	values := map[string]interface{}{
		"modified_by": param.ModifiedBy,
		"state":       param.State,
	}
	if param.Title != "" {
		values["title"] = param.Title
	}
	if param.CoverImageUrl != "" {
		values["cover_image_url"] = param.CoverImageUrl
	}
	if param.Desc != "" {
		values["desc"] = param.Desc
	}
	if param.Content != "" {
		values["content"] = param.Content
	}

	return article.Update(d.engine, values)
}

func (d *Dao) GetArticle(id uint32, state uint8) (model.Article, error) {
	article := model.Article{ID: id, State: state}
	return article.Get(d.engine)
}

func (d *Dao) DeleteArticle(id uint32) error {
	article := model.Article{ID: id}
	return article.Delete(d.engine)
}

func (d *Dao) CountArticles(state uint8) (int, error) {
	article := model.Article{State: state}
	return article.Count(d.engine)
}

func (d *Dao) ListArticles(state uint8, pageNum, pageSize int) ([]*model.ArticleEntity, error) {
	article := model.Article{State: state}
	return article.List(d.engine, app.GetPageOffset(pageNum, pageSize), pageSize)
}

func (d *Dao) CountArticlesByTag(tagId uint32, state uint8) (int, error) {
	article := model.Article{State: state}
	return article.CountByTag(d.engine, tagId)
}

func (d *Dao) ListArticlesByTag(tagID uint32, state uint8, pageNum, pageSize int) ([]*model.ArticleEntity, error) {
	article := model.Article{State: state}
	return article.ListByTag(d.engine, tagID, app.GetPageOffset(pageNum, pageSize), pageSize)
}
