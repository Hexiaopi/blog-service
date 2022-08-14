package dao

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/hexiaopi/blog-service/internal/model"
)

type AuthDao struct {
	db *gorm.DB
}

func NewAuthDao(db *gorm.DB) *AuthDao {
	return &AuthDao{db: db}
}

func (d *AuthDao) Get(ctx context.Context, key, secret string) (*model.Auth, error) {
	auth := model.Auth{AppKey: key, AppSecret: secret}
	return auth.Get(d.db)
}
