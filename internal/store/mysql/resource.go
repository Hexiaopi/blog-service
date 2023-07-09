package mysql

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/hexiaopi/blog-service/internal/model"
)

type ResourceDao struct {
	db *gorm.DB
}

func NewResourceDao(db *gorm.DB) *ResourceDao {
	return &ResourceDao{db: db}
}

func (dao *ResourceDao) Create(ctx context.Context, resource *model.Resource) error {
	now := time.Now()
	resource.CreateTime = now
	resource.UpdateTime = now
	return dao.db.WithContext(ctx).Create(resource).Error
}

func (dao *ResourceDao) Get(ctx context.Context, id int) (*model.Resource, error) {
	var resource model.Resource
	if err := dao.db.WithContext(ctx).First(&resource, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &resource, nil
}

func (dao *ResourceDao) Update(ctx context.Context, resource *model.Resource) error {
	resource.UpdateTime = time.Now()
	return dao.db.WithContext(ctx).Updates(resource).Error
}

func (dao *ResourceDao) Delete(ctx context.Context, id int) error {
	resource := model.Resource{ID: id}
	return dao.db.WithContext(ctx).Delete(&resource).Error
}

func (dao *ResourceDao) List(ctx context.Context, opt *model.ListOption) ([]model.Resource, error) {
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
	resources := make([]model.Resource, 0)
	if err := query.Model(&model.Resource{}).
		Where("state = ?", opt.State).
		Find(&resources).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
	}
	return resources, nil
}

func (dao *ResourceDao) Count(ctx context.Context, opt *model.ListOption) (int64, error) {
	query := dao.db.WithContext(ctx)
	var count int64
	if opt.Name != "" {
		query = query.Where("name = ?", opt.Name)
	}
	if opt.Sort != "" {
		query = query.Order(opt.GetSortType())
	}
	if err := query.Model(&model.Resource{}).
		Where("state = ?", opt.State).
		Count(&count).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, nil
		}
	}
	return count, nil
}
