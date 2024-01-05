package mysql

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/hexiaopi/blog-service/internal/model"
	"github.com/hexiaopi/blog-service/internal/store"
)

type RoleRestDao struct {
	db *gorm.DB
}

var _ store.RoleRestStore = (*RoleRestDao)(nil)

func NewRoleRestDao(db *gorm.DB) *RoleRestDao {
	return &RoleRestDao{
		db: db,
	}
}

func (dao *RoleRestDao) Create(ctx context.Context, roleId, restId int) error {
	roleMenu := &model.RoleRest{
		RoleId:     roleId,
		RestId:     restId,
		CreateTime: time.Now(),
	}
	return dao.db.WithContext(ctx).Create(roleMenu).Error
}

func (dao *RoleRestDao) DeleteByRole(ctx context.Context, roleId int) error {
	return dao.db.WithContext(ctx).Where("role_id =?", roleId).Delete(&model.RoleRest{}).Error
}

func (dao *RoleRestDao) ListByRole(ctx context.Context, roleId int) ([]int, error) {
	ids := make([]int, 0)
	if err := dao.db.Model(&model.RoleRest{}).WithContext(ctx).Select("rest_id").Where("role_id = ?", roleId).Find(&ids).Error; err != nil {
		return nil, err
	}
	return ids, nil
}
