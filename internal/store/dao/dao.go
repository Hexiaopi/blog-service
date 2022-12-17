package dao

import (
	"errors"

	"github.com/jinzhu/gorm"

	"github.com/hexiaopi/blog-service/internal/store"
)

type datastore struct {
	db *gorm.DB
}

func NewDao(db *gorm.DB) *datastore {
	return &datastore{db: db}
}

func (ds *datastore) Articles() store.ArticleStore {
	return NewArticleDao(ds.db)
}

func (ds *datastore) Tags() store.TagStore {
	return NewTagDao(ds.db)
}

func (ds *datastore) Users() store.UserStore {
	return NewUserDao(ds.db)
}

func (ds *datastore) Close() error {
	db := ds.db.DB()
	if db == nil {
		return errors.New("get gorm db instance failed")
	}

	return db.Close()
}
