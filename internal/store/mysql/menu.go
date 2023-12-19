package mysql

import (
	"context"
	"errors"

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
