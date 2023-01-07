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
	return dao.db.Create(tag).Error
}

func (dao *TagDao) Get(ctx context.Context, id int) (*model.Tag, error) {
	var tag model.Tag
	if err := dao.db.First(&tag, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &tag, nil
}

func (dao *TagDao) Update(ctx context.Context, tag *model.Tag) error {
	tag.UpdateTime = time.Now()
	return dao.db.Updates(tag).Error
}

func (dao *TagDao) Delete(ctx context.Context, id int) error {
	tag := model.Tag{ID: id}
	return dao.db.Delete(&tag).Error
}

func (dao *TagDao) List(ctx context.Context, opt *model.ListOption) ([]model.Tag, int64, error) {
	query := dao.db
	if opt.Page >= 0 && opt.Limit > 0 {
		query = dao.db.Offset(opt.GetPageOffset()).Limit(opt.Limit)
	}
	var tags []model.Tag
	var err error
	var total int64
	if opt.Name != "" {
		query = query.Where("name = ?", opt.Name)
	}
	if opt.Sort != "" {
		query = query.Order(opt.GetSortType())
	}
	if err = query.Where("state = ?", opt.State).Find(&tags).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return tags, total, nil
}
