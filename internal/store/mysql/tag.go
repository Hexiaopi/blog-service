package mysql

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/hexiaopi/blog-service/internal/model"
)

type TagDao struct {
	db *gorm.DB
}

func NewTagDao(db *gorm.DB) *TagDao {
	return &TagDao{db: db}
}

func (dao *TagDao) Create(ctx context.Context, tag *model.Tag) error {
	tag.CreateTime = time.Now()
	tag.UpdateTime = time.Now()
	return dao.db.WithContext(ctx).Create(tag).Error
}

func (dao *TagDao) Get(ctx context.Context, id int) (*model.Tag, error) {
	var tag model.Tag
	if err := dao.db.WithContext(ctx).First(&tag, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &tag, nil
}

func (dao *TagDao) Update(ctx context.Context, tag *model.Tag) error {
	tag.UpdateTime = time.Now()
	return dao.db.WithContext(ctx).Updates(tag).Error
}

func (dao *TagDao) Delete(ctx context.Context, id int) error {
	tag := model.Tag{ID: id}
	return dao.db.WithContext(ctx).Delete(&tag).Error
}

func (dao *TagDao) List(ctx context.Context, opt *model.ListOption) ([]model.Tag, error) {
	query := dao.db.WithContext(ctx)
	if opt.Page >= 0 && opt.Limit > 0 {
		query = query.Offset(opt.GetPageOffset()).Limit(opt.Limit)
	}
	var tags []model.Tag
	var err error
	if opt.Name != "" {
		query = query.Where("name = ?", opt.Name)
	}
	if opt.Sort != "" {
		query = query.Order(opt.GetSortType())
	}
	if err = query.Select("blog_tag.* ,count(blog_article_tag.article_id) as article_total").
		Joins("left join blog_article_tag on blog_tag.id = blog_article_tag.tag_id").
		Group("id").
		Where("state = ?", opt.State).Find(&tags).Error; err != nil {
		return nil, err
	}
	// for i, tag := range tags {
	// 	count, err := dao.CountArticle(ctx, tag.ID)
	// 	if err != nil {
	// 		continue
	// 	}
	// 	tags[i].Articles = count
	// }
	return tags, nil
}

func (dao *TagDao) Count(ctx context.Context, opt *model.ListOption) (int64, error) {
	query := dao.db.WithContext(ctx)
	if opt.Page >= 0 && opt.Limit > 0 {
		query = query.Offset(opt.GetPageOffset()).Limit(opt.Limit)
	}
	var err error
	var total int64
	if opt.Name != "" {
		query = query.Where("name = ?", opt.Name)
	}
	if opt.Sort != "" {
		query = query.Order(opt.GetSortType())
	}
	if err = query.Model(&model.Tag{}).Where("state = ?", opt.State).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func (dao *TagDao) CountArticle(ctx context.Context, id int) (int64, error) {
	var count int64
	if err := dao.db.WithContext(ctx).Table("blog_article_tag").
		Where("tag_id = ?", id).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
