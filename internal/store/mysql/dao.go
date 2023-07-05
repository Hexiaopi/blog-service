package mysql

import (
	"errors"

	"gorm.io/gorm"

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

func (ds *datastore) Roles() store.RoleStore {
	return NewRoleDao(ds.db)
}

func (ds *datastore) Systems() store.SystemConfigStore {
	return NewSystemConfigDao(ds.db)
}

func (ds *datastore) Resources() store.ResourceStore {
	return NewResourceDao(ds.db)
}

func (ds *datastore) Operations() store.OperationStore {
	return NewOperationDao(ds.db)
}

func (ds *datastore) Close() error {
	db, err := ds.db.DB()
	if err != nil {
		return err
	}
	if db == nil {
		return errors.New("get gorm db instance failed")
	}

	return db.Close()
}
