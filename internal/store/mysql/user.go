package mysql

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/hexiaopi/blog-service/internal/model"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

func (dao *UserDao) Get(ctx context.Context, name string) (*model.User, error) {
	var user model.User
	err := dao.db.Where("name = ?", name).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
