package dao

import (
	"context"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/hexiaopi/blog-service/internal/constant"
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
		Operator:      param.Operator,
		CreateTime:    time.Now(),
		UpdateTime:    time.Now(),
	}
	return dao.db.Create(&article).Error
}

func (dao *ArticleDao) Get(ctx context.Context, id int) (*entity.Article, error) {
	var article model.Article
	if err := dao.db.First(&article, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	result := entity.Article{
		Id:            article.ID,
		Title:         article.Title,
		Desc:          article.Desc,
		Content:       article.Content,
		CoverImageUrl: article.CoverImageUrl,
		State:         article.State,
		CreateTime:    article.CreateTime.Format(constant.DefaultTimeFormat),
		UpdateTime:    article.UpdateTime.Format(constant.DefaultTimeFormat),
		Operator:      article.Operator,
	}
	tags, err := dao.getTags(article.ID)
	if err != nil {
		return nil, err
	}
	result.Tags = tags
	return &result, nil
}

func (dao *ArticleDao) Update(ctx context.Context, param *entity.Article) error {
	article := model.Article{
		ID:            param.Id,
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State:         param.State,
		Operator:      param.Operator,
		UpdateTime:    time.Now(),
	}
	return dao.db.Model(&model.Article{}).Update(article).Error
}

func (dao *ArticleDao) Delete(ctx context.Context, id int) error {
	article := model.Article{ID: id}
	return dao.db.Delete(&article).Error
}

func (dao *ArticleDao) List(ctx context.Context, opt *entity.ListOption) ([]entity.Article, int64, error) {
	query := dao.db
	if opt.PageNum >= 0 && opt.PageSize > 0 {
		query = dao.db.Offset(opt.GetPageOffset()).Limit(opt.PageSize)
	}
	var count int64
	articles := make([]model.Article, 0)
	if err := query.Model(&model.Article{}).
		Where("state = ?", opt.State).
		Count(&count).
		Find(&articles).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, 0, nil
		}
	}
	result := make([]entity.Article, len(articles))
	for i, article := range articles {
		result[i] = entity.Article{
			Id:            article.ID,
			Title:         article.Title,
			Desc:          article.Desc,
			Content:       article.Content,
			CoverImageUrl: article.CoverImageUrl,
			State:         article.State,
			CreateTime:    article.CreateTime.Format(constant.DefaultTimeFormat),
			UpdateTime:    article.UpdateTime.Format(constant.DefaultTimeFormat),
			Operator:      article.Operator,
			Tags:          make([]entity.Tag, 0),
		}
		tags, err := dao.getTags(article.ID)
		if err != nil {
			return nil, 0, err
		}
		if tags != nil {
			result[i].Tags = tags
		}
	}
	return result, count, nil
}

func (dao *ArticleDao) getTags(id int) ([]entity.Tag, error) {
	tags := make([]model.Tag, 0)
	if err := dao.db.Model(&model.Tag{}).
		Joins("left join blog_article_tag on blog_tag.id = blog_article_tag.tag_id").
		Where("blog_article_tag.article_id = ?", id).
		Find(&tags).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return nil, err
		}
	}
	result := make([]entity.Tag, len(tags))
	for i, tag := range tags {
		result[i] = entity.Tag{
			Id:         tag.ID,
			Name:       tag.Name,
			Desc:       tag.Desc,
			State:      tag.State,
			CreateTime: tag.CreateTime.Format(constant.DefaultTimeFormat),
			UpdateTime: tag.UpdateTime.Format(constant.DefaultTimeFormat),
			Operator:   tag.Operator,
		}
	}
	return result, nil
}
