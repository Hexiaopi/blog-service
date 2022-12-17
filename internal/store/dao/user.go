package dao

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/hexiaopi/blog-service/internal/model"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

func (d *UserDao) Get(ctx context.Context, name string) (*model.User, error) {
	user := model.User{Name: name}
	return user.Get(d.db)
}
