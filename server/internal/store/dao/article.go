package dao

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/hexiaopi/blog-service/internal/app"
	"github.com/hexiaopi/blog-service/internal/entity"
	"github.com/hexiaopi/blog-service/internal/model"
)

type ArticleDao struct {
	db *gorm.DB
}

func NewArticleDao(db *gorm.DB) *ArticleDao {
	return &ArticleDao{db: db}
}

func (dao *ArticleDao) Create(ctx context.Context, param *entity.Article) error {
	article := model.Article{
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State:         param.State,
		CreatedBy:     param.CreatedBy,
	}
	if err := dao.db.Create(&article).Error; err != nil {
		return err
	}
	return nil
}

func (d *ArticleDao) Update(ctx context.Context, param *entity.Article) error {
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

	return article.Update(d.db, values)
}

func (d *ArticleDao) Get(ctx context.Context, id int) (*entity.Article, error) {
	article := model.Article{ID: id}
	a, err := article.Get(d.db)
	if err != nil {
		return nil, err
	}
	return &entity.Article{
		ID:      a.ID,
		Title:   a.Title,
		Desc:    a.Desc,
		Content: a.Content,
	}, nil
}

func (d *ArticleDao) Delete(ctx context.Context, id int) error {
	article := model.Article{ID: id}
	return article.Delete(d.db)
}

func (d *ArticleDao) List(ctx context.Context, opt *entity.ListOption) ([]*entity.Article, int64, error) {
	article := model.Article{State: opt.State}
	as, total, err := article.List(d.db, app.GetPageOffset(opt.PageNum, opt.PageSize), opt.PageSize)
	if err != nil {
		return nil, 0, err
	}
	result := make([]*entity.Article, len(as))
	for k, v := range as {
		result[k] = &entity.Article{
			ID:      int(v.ArticleID),
			Title:   v.ArticleTitle,
			Desc:    v.ArticleDesc,
			Content: v.Content,
		}
	}
	return result, total, nil
}
