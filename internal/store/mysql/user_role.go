package mysql

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/hexiaopi/blog-service/internal/entity"
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

func (dao *UserRoleDao) Create(ctx context.Context, userId, roleId int) error {
	userRole := &model.UserRole{
		UserId:     userId,
		RoleId:     roleId,
		CreateTime: time.Now(),
	}
	return dao.db.WithContext(ctx).Create(userRole).Error
}

func (dao *UserRoleDao) Delete(ctx context.Context, userId, roleId int) error {
	return dao.db.WithContext(ctx).
		Where("user_id = ?", userId).
		Where("role_id = ?", roleId).
		Delete(&model.UserRole{}).Error
}

func (dao *UserRoleDao) DeleteByUser(ctx context.Context, userId int) error {
	return dao.db.WithContext(ctx).
		Where("user_id = ?", userId).
		Delete(&model.UserRole{}).Error
}

func (dao *UserRoleDao) ListUserRole(ctx context.Context, userId int) ([]entity.Role, error) {
	roles := make([]model.Role, 0)
	if err := dao.db.WithContext(ctx).Model(&model.Role{}).
		Joins("inner join sys_user_role on sys_role.id = sys_user_role.role_id").
		Where("sys_user_role.user_id = ?", userId).
		Find(&roles).Error; err != nil {
		return nil, err
	}
	result := make([]entity.Role, 0, len(roles))
	for _, role := range roles {
		result = append(result, *entity.ToEntityRole(&role))
	}
	return result, nil
}
