package mysql

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/hexiaopi/blog-service/internal/model"
	"github.com/hexiaopi/blog-service/internal/store"
)

type RoleMenuDao struct {
	db *gorm.DB
}

var _ store.RoleMenuStore = (*RoleMenuDao)(nil)

func NewRoleMenuDao(db *gorm.DB) *RoleMenuDao {
	return &RoleMenuDao{
		db: db,
	}
}

func (dao *RoleMenuDao) Create(ctx context.Context, roleId, menuId int) error {
	roleMenu := &model.RoleMenu{
		RoleId:     roleId,
		MenuId:     menuId,
		CreateTime: time.Now(),
	}
	return dao.db.WithContext(ctx).Create(roleMenu).Error
}

func (dao *RoleMenuDao) DeleteByRole(ctx context.Context, roleId int) error {
	return dao.db.WithContext(ctx).Where("role_id =?", roleId).Delete(&model.RoleMenu{}).Error
}

func (dao *RoleMenuDao) ListByRole(ctx context.Context, roleId int) ([]int, error) {
	ids := make([]int, 0)
	if err := dao.db.Model(&model.RoleMenu{}).WithContext(ctx).Select("menu_id").Where("role_id = ?", roleId).Find(&ids).Error; err != nil {
		return nil, err
	}
	return ids, nil
}
