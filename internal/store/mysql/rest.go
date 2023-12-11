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

type SysRestDao struct {
	db *gorm.DB
}

var _ store.SysRestStore = (*SysRestDao)(nil)

func NewSysRestDao(db *gorm.DB) *SysRestDao {
	return &SysRestDao{db: db}
}

func (dao *SysRestDao) Create(ctx context.Context, sysRest *model.SysRest) error {
	sysRest.CreateTime = time.Now()
	sysRest.UpdateTime = time.Now()
	return dao.db.WithContext(ctx).Create(&sysRest).Error
}

func (dao *SysRestDao) Get(ctx context.Context, id int) (*model.SysRest, error) {
	var sysRest model.SysRest
	if err := dao.db.WithContext(ctx).First(&sysRest, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &sysRest, nil
}

func (dao *SysRestDao) Update(ctx context.Context, sysRest *model.SysRest) error {
	sysRest.UpdateTime = time.Now()
	return dao.db.WithContext(ctx).Updates(&sysRest).Error
}

func (dao *SysRestDao) Delete(ctx context.Context, id int) error {
	SysRest := model.SysRest{ID: id}
	return dao.db.WithContext(ctx).Delete(&SysRest).Error
}

func (dao *SysRestDao) List(ctx context.Context, opt *entity.ListOption) ([]model.SysRest, error) {
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
	SysRests := make([]model.SysRest, 0)
	if err := query.Model(&model.SysRest{}).
		Where("parent_id = ?", opt.ParentId).
		Find(&SysRests).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
	}
	return SysRests, nil
}

func (dao *SysRestDao) Count(ctx context.Context, opt *entity.ListOption) (int64, error) {
	query := dao.db.WithContext(ctx)
	var count int64
	if opt.Name != "" {
		query = query.Where("name = ?", opt.Name)
	}
	if opt.Sort != "" {
		query = query.Order(opt.GetSortType())
	}
	if err := query.Model(&model.SysRest{}).
		Where("parent_id = ?", opt.ParentId).
		Count(&count).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, nil
		}
	}
	return count, nil
}
