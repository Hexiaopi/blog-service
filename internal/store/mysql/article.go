package mysql

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/hexiaopi/blog-service/internal/model"
)

type ArticleDao struct {
	db *gorm.DB
}

func NewArticleDao(db *gorm.DB) *ArticleDao {
	return &ArticleDao{db: db}
}

func (dao *ArticleDao) Create(ctx context.Context, article *model.Article) error {
	article.CreateTime = time.Now()
	article.UpdateTime = time.Now()
	return dao.db.WithContext(ctx).Create(article).Error
}

func (dao *ArticleDao) Get(ctx context.Context, id int) (*model.Article, error) {
	var article model.Article
	if err := dao.db.WithContext(ctx).Preload("Tags").First(&article, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &article, nil
}

func (dao *ArticleDao) Update(ctx context.Context, article *model.Article) error {
	article.UpdateTime = time.Now()
	return dao.db.WithContext(ctx).Updates(article).Error
}

func (dao *ArticleDao) Delete(ctx context.Context, id int) error {
	article := model.Article{ID: id}
	return dao.db.WithContext(ctx).Delete(&article).Error
}

func (dao *ArticleDao) List(ctx context.Context, opt *model.ListOption) ([]model.Article, error) {
	query := dao.db.WithContext(ctx)
	if opt.Page >= 0 && opt.Limit > 0 {
		query = query.Offset(opt.GetPageOffset()).Limit(opt.Limit)
	}
	if opt.Name != "" {
		query = query.Where("name = ?", opt.Name)
	}
	if opt.Sort != "" {
		query = query.Order(opt.GetSortType())
	}
	articles := make([]model.Article, 0)
	if err := query.Model(&model.Article{}).
		Where("state = ?", opt.State).
		Preload("Tags").
		Find(&articles).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
	}
	return articles, nil
}

func (dao *ArticleDao) Count(ctx context.Context, opt *model.ListOption) (int64, error) {
	query := dao.db.WithContext(ctx)
	var count int64
	if opt.Name != "" {
		query = query.Where("name = ?", opt.Name)
	}
	if opt.Sort != "" {
		query = query.Order(opt.GetSortType())
	}
	if err := query.Model(&model.Article{}).
		Where("state = ?", opt.State).
		Preload("Tags").
		Count(&count).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, nil
		}
	}
	return count, nil
}
