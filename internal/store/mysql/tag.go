package mysql

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/hexiaopi/blog-service/internal/entity"
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

func (dao *TagDao) Create(ctx context.Context, tag *entity.Tag) error {
	t := tag.ToModel()
	t.CreateTime = time.Now()
	t.UpdateTime = time.Now()
	return dao.db.WithContext(ctx).Create(&t).Error
}

func (dao *TagDao) Get(ctx context.Context, id int) (*entity.Tag, error) {
	var tag model.Tag
	if err := dao.db.WithContext(ctx).First(&tag, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	result := entity.ToEntityTag(&tag)
	return result, nil
}

func (dao *TagDao) Update(ctx context.Context, tag *entity.Tag) error {
	t := tag.ToModel()
	t.UpdateTime = time.Now()
	return dao.db.WithContext(ctx).Updates(&t).Error
}

func (dao *TagDao) Delete(ctx context.Context, id int) error {
	tag := model.Tag{ID: id}
	return dao.db.WithContext(ctx).Delete(&tag).Error
}

func (dao *TagDao) List(ctx context.Context, opt *entity.ListOption) ([]entity.Tag, error) {
	query := dao.db.WithContext(ctx)
	if opt.Page >= 0 && opt.Limit > 0 {
		query = query.Offset(opt.GetPageOffset()).Limit(opt.Limit)
	}
	var tags []entity.Tag
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
		var tag entity.Tag
		if err := rows.Scan(&tag.ID, &tag.Name, &tag.Desc, &tag.State, &tag.CreateTime, &tag.UpdateTime, &tag.Operator, &tag.ArticleTotal); err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

func (dao *TagDao) Count(ctx context.Context, opt *entity.ListOption) (int64, error) {
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
