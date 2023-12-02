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

type RoleDao struct {
	db *gorm.DB
}

var _ store.RoleStore = (*RoleDao)(nil)

func NewRoleDao(db *gorm.DB) *RoleDao {
	return &RoleDao{db: db}
}

func (dao *RoleDao) Create(ctx context.Context, role *entity.Role) error {
	r := role.ToModel()
	r.CreateTime = time.Now()
	r.UpdateTime = time.Now()
	return dao.db.WithContext(ctx).Create(r).Error
}

func (dao *RoleDao) Get(ctx context.Context, id int) (*entity.Role, error) {
	var role model.Role
	if err := dao.db.WithContext(ctx).First(&role, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return entity.ToEntityRole(&role), nil
}

func (dao *RoleDao) Update(ctx context.Context, role *entity.Role) error {
	r := role.ToModel()
	r.UpdateTime = time.Now()
	return dao.db.WithContext(ctx).Updates(r).Error
}

func (dao *RoleDao) Delete(ctx context.Context, id int) error {
	role := model.Role{ID: id}
	return dao.db.WithContext(ctx).Delete(&role).Error
}

func (dao *RoleDao) List(ctx context.Context, opt *entity.ListOption) ([]entity.Role, error) {
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
		Find(&roles).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
	}
	result := make([]entity.Role, 0, len(roles))
	for _, role := range roles {
		result = append(result, *entity.ToEntityRole(&role))
	}
	return result, nil
}

func (dao *RoleDao) Count(ctx context.Context, opt *entity.ListOption) (int64, error) {
	query := dao.db.WithContext(ctx)
	var count int64
	if opt.Name != "" {
		query = query.Where("name = ?", opt.Name)
	}
	if opt.Sort != "" {
		query = query.Order(opt.GetSortType())
	}
	if err := query.Model(&model.Role{}).
		Count(&count).
		Error; err != nil {
		return 0, err
	}
	return count, nil
}
