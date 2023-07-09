package mysql

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/hexiaopi/blog-service/internal/model"
)

type RoleDao struct {
	db *gorm.DB
}

func NewRoleDao(db *gorm.DB) *RoleDao {
	return &RoleDao{db: db}
}

func (dao *RoleDao) Create(ctx context.Context, role *model.Role) error {
	role.CreateTime = time.Now()
	role.UpdateTime = time.Now()
	return dao.db.WithContext(ctx).Create(role).Error
}

func (dao *RoleDao) Get(ctx context.Context, id int) (*model.Role, error) {
	var role model.Role
	if err := dao.db.WithContext(ctx).First(&role, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &role, nil
}

func (dao *RoleDao) Update(ctx context.Context, role *model.Role) error {
	role.UpdateTime = time.Now()
	return dao.db.WithContext(ctx).Updates(role).Error
}

func (dao *RoleDao) Delete(ctx context.Context, id int) error {
	role := model.Role{ID: id}
	return dao.db.WithContext(ctx).Delete(&role).Error
}

func (dao *RoleDao) List(ctx context.Context, opt *model.ListOption) ([]model.Role, error) {
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
	roles := make([]model.Role, 0)
	if err := query.Model(&model.Role{}).
		Where("state = ?", opt.State).
		Find(&roles).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
	}
	return roles, nil
}

func (dao *RoleDao) Count(ctx context.Context, opt *model.ListOption) (int64, error) {
	query := dao.db.WithContext(ctx)
	var count int64
	if opt.Name != "" {
		query = query.Where("name = ?", opt.Name)
	}
	if opt.Sort != "" {
		query = query.Order(opt.GetSortType())
	}
	if err := query.Model(&model.Role{}).
		Where("state = ?", opt.State).
		Count(&count).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, nil
		}
	}
	return count, nil
}
