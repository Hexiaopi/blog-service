package mysql

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/hexiaopi/blog-service/internal/model"
	"github.com/hexiaopi/blog-service/internal/store"
)

type TagDao struct {
	db *gorm.DB
}

var _ store.TagStore = (*TagDao)(nil)

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
	rows, err := query.Model(&model.Tag{}).Select("blog_tag.id,name,`desc`,`state`,create_time,update_time,operator,count(blog_article_tag.article_id) as article_total").
		Joins("left join blog_article_tag on blog_tag.id = blog_article_tag.tag_id").
		Group("id").
		Where("state = ?", opt.State).
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tag model.Tag
		if err := rows.Scan(&tag.ID, &tag.Name, &tag.Desc, &tag.State, &tag.CreateTime, &tag.UpdateTime, &tag.Operator, &tag.ArticleTotal); err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

func (dao *TagDao) Count(ctx context.Context, opt *model.ListOption) (int64, error) {
	query := dao.db.WithContext(ctx)
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
