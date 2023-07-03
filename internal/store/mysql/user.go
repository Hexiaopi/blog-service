package mysql

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/hexiaopi/blog-service/internal/model"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

func (dao *UserDao) Create(ctx context.Context, user *model.User) error {
	user.CreateTime = time.Now()
	user.UpdateTime = time.Now()
	return dao.db.WithContext(ctx).Create(user).Error
}

func (dao *UserDao) Update(ctx context.Context, user *model.User) error {
	user.UpdateTime = time.Now()
	return dao.db.WithContext(ctx).Updates(user).Error
}

func (dao *UserDao) Delete(ctx context.Context, id int) error {
	user := model.User{ID: id}
	return dao.db.WithContext(ctx).Delete(&user).Error
}

func (dao *UserDao) Get(ctx context.Context, name string) (*model.User, error) {
	var user model.User
	err := dao.db.WithContext(ctx).Where("name = ?", name).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (dao *UserDao) List(ctx context.Context, opt *model.ListOption) ([]model.User, error) {
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
	users := make([]model.User, 0)
	if err := query.Model(&model.User{}).
		//Where("state = ?", opt.State).
		Find(&users).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
	}
	return users, nil
}

func (dao *UserDao) Count(ctx context.Context, opt *model.ListOption) (int64, error) {
	query := dao.db.WithContext(ctx)
	if opt.Page >= 0 && opt.Limit > 0 {
		query = query.Offset(opt.GetPageOffset()).Limit(opt.Limit)
	}
	var count int64
	if opt.Name != "" {
		query = query.Where("name = ?", opt.Name)
	}
	if opt.Sort != "" {
		query = query.Order(opt.GetSortType())
	}
	if err := query.Model(&model.User{}).
		//Where("state = ?", opt.State).
		Count(&count).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, nil
		}
	}
	return count, nil
}
