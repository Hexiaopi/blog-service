package mysql

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/hexiaopi/blog-service/internal/model"
	"github.com/hexiaopi/blog-service/internal/store"
)

type UserRoleDao struct {
	db *gorm.DB
}

var _ store.UserRoleStore = (*UserRoleDao)(nil)

func NewUserRoleDao(db *gorm.DB) *UserRoleDao {
	return &UserRoleDao{
		db: db,
	}
}

func (dao *UserRoleDao) Create(ctx context.Context, userRole *model.UserRole) error {
	userRole.CreateTime = time.Now()
	return dao.db.WithContext(ctx).Create(userRole).Error
}

func (dao *UserRoleDao) Delete(ctx context.Context, userRole *model.UserRole) error {
	return dao.db.WithContext(ctx).
		Where("user_id = ?", userRole.UserId).
		Where("role_id = ?", userRole.RoleId).
		Delete(userRole).Error
}
