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

type UserDao struct {
	db *gorm.DB
}

var _ store.UserStore = (*UserDao)(nil)

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

func (dao *UserDao) Get(ctx context.Context, name string) (*entity.User, error) {
	var user model.User
	err := dao.db.WithContext(ctx).Where("name = ?", name).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return entity.ToEntityUser(&user), nil
}

func (dao *UserDao) Create(ctx context.Context, user *entity.User) (int, error) {
	u := user.ToModel()
	u.CreateTime = time.Now()
	u.UpdateTime = time.Now()
	if err := dao.db.WithContext(ctx).Create(u).Error; err != nil {
		return 0, err
	}
	return u.ID, nil
}

func (dao *UserDao) Update(ctx context.Context, user *entity.User) error {
	u := user.ToModel()
	u.UpdateTime = time.Now()
	return dao.db.WithContext(ctx).Updates(u).Error
}

func (dao *UserDao) Delete(ctx context.Context, id int) error {
	user := model.User{ID: id}
	return dao.db.WithContext(ctx).Delete(&user).Error
}

func (dao *UserDao) List(ctx context.Context, opt *entity.ListOption) ([]entity.User, error) {
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
		Where("state = ?", opt.State).
		Find(&users).
		Error; err != nil {
		return nil, err
	}
	result := make([]entity.User, 0, len(users))
	for _, user := range users {
		result = append(result, *entity.ToEntityUser(&user))
	}
	return result, nil
}

func (dao *UserDao) Count(ctx context.Context, opt *entity.ListOption) (int64, error) {
	query := dao.db.WithContext(ctx)
	var count int64
	if opt.Name != "" {
		query = query.Where("name = ?", opt.Name)
	}
	if opt.Sort != "" {
		query = query.Order(opt.GetSortType())
	}
	if err := query.Model(&model.User{}).
		Where("state = ?", opt.State).
		Count(&count).
		Error; err != nil {
		return 0, err
	}
	return count, nil
}
