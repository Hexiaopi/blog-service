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

type SysMenuDao struct {
	db *gorm.DB
}

var _ store.SysMenuStore = (*SysMenuDao)(nil)

func NewSysMenuDao(db *gorm.DB) *SysMenuDao {
	return &SysMenuDao{db: db}
}

func (dao *SysMenuDao) Create(ctx context.Context, sysRest *model.SysMenu) error {
	sysRest.CreateTime = time.Now()
	sysRest.UpdateTime = time.Now()
	return dao.db.WithContext(ctx).Create(&sysRest).Error
}

func (dao *SysMenuDao) Get(ctx context.Context, id int) (*model.SysMenu, error) {
	var sysMenu model.SysMenu
	if err := dao.db.WithContext(ctx).First(&sysMenu, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &sysMenu, nil
}

func (dao *SysMenuDao) Update(ctx context.Context, sysMenu *model.SysMenu) error {
	sysMenu.UpdateTime = time.Now()
	return dao.db.WithContext(ctx).Updates(&sysMenu).Error
}

func (dao *SysMenuDao) Delete(ctx context.Context, id int) error {
	SysMenu := model.SysMenu{ID: id}
	return dao.db.WithContext(ctx).Delete(&SysMenu).Error
}

func (dao *SysMenuDao) List(ctx context.Context, opt *entity.ListOption) ([]model.SysMenu, error) {
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
	SysMenus := make([]model.SysMenu, 0)
	if err := query.Model(&model.SysMenu{}).
		Where("parent_id = ?", opt.ParentId).
		Find(&SysMenus).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
	}
	return SysMenus, nil
}

func (dao *SysMenuDao) Count(ctx context.Context, opt *entity.ListOption) (int64, error) {
	query := dao.db.WithContext(ctx)
	var count int64
	if opt.Name != "" {
		query = query.Where("name = ?", opt.Name)
	}
	if opt.Sort != "" {
		query = query.Order(opt.GetSortType())
	}
	if err := query.Model(&model.SysMenu{}).
		Where("parent_id = ?", opt.ParentId).
		Count(&count).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, nil
		}
	}
	return count, nil
}
