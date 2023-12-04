package mysql

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/hexiaopi/blog-service/internal/entity"
	"github.com/hexiaopi/blog-service/internal/model"
	"github.com/hexiaopi/blog-service/internal/store"
)

type ResourceDao struct {
	db *gorm.DB
}

var _ store.ResourceStore = (*ResourceDao)(nil)

func NewResourceDao(db *gorm.DB) *ResourceDao {
	return &ResourceDao{db: db}
}

func (dao *ResourceDao) Create(ctx context.Context, resource *entity.Resource) error {
	r := resource.ToModel()
	r.CreateTime = time.Now()
	r.UpdateTime = time.Now()
	return dao.db.WithContext(ctx).Create(&r).Error
}

func (dao *ResourceDao) Get(ctx context.Context, id int) (*entity.Resource, error) {
	var resource model.Resource
	if err := dao.db.WithContext(ctx).First(&resource, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	result := entity.ToEntityResource(&resource)
	return result, nil
}

func (dao *ResourceDao) Update(ctx context.Context, resource *entity.Resource) error {
	r := resource.ToModel()
	r.UpdateTime = time.Now()
	return dao.db.WithContext(ctx).Updates(&r).Error
}

func (dao *ResourceDao) Delete(ctx context.Context, id int) error {
	resource := model.Resource{ID: id}
	return dao.db.WithContext(ctx).Delete(&resource).Error
}

func (dao *ResourceDao) List(ctx context.Context, opt *entity.ListOption) ([]entity.Resource, error) {
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
	result := make([]entity.Resource, 0, len(resources))
	for _, resource := range resources {
		result = append(result, *entity.ToEntityResource(&resource))
	}
	return result, nil
}

func (dao *ResourceDao) Count(ctx context.Context, opt *entity.ListOption) (int64, error) {
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
