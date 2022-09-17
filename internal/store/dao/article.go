package dao

import (
	"context"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/hexiaopi/blog-service/global"
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
		Name:          param.Name,
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
		Name:          article.Name,
		Desc:          article.Desc,
		Content:       article.Content,
		CoverImageUrl: article.CoverImageUrl,
		State:         article.State,
		CreateTime:    article.CreateTime.Format(global.DefaultTimeFormat),
		UpdateTime:    article.UpdateTime.Format(global.DefaultTimeFormat),
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
		Name:          param.Name,
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
	if opt.Page >= 0 && opt.Limit > 0 {
		query = dao.db.Offset(opt.GetPageOffset()).Limit(opt.Limit)
	}
	var count int64
	if opt.Name != "" {
		query = query.Where("name = ?", opt.Name)
	}
	if opt.Sort != "" {
		query = query.Order(opt.GetSortType())
	}
	articles := make([]model.Article, 0)
	if err := query.Model(&model.Article{}).Where("state = ?", opt.State).Find(&articles).Count(&count).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, 0, nil
		}
	}
	result := make([]entity.Article, len(articles))
	for i, article := range articles {
		result[i] = entity.Article{
			Id:            article.ID,
			Name:          article.Name,
			Desc:          article.Desc,
			Content:       article.Content,
			CoverImageUrl: article.CoverImageUrl,
			State:         article.State,
			CreateTime:    article.CreateTime.Format(global.DefaultTimeFormat),
			UpdateTime:    article.UpdateTime.Format(global.DefaultTimeFormat),
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
			CreateTime: tag.CreateTime.Format(global.DefaultTimeFormat),
			UpdateTime: tag.UpdateTime.Format(global.DefaultTimeFormat),
			Operator:   tag.Operator,
		}
	}
	return result, nil
}
