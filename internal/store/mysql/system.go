package mysql

import (
	"context"

	"gorm.io/gorm"

	"github.com/hexiaopi/blog-service/internal/model"
	"github.com/hexiaopi/blog-service/internal/store"
)

type SystemConfigDao struct {
	db *gorm.DB
}

var _ store.SystemConfigStore = (*SystemConfigDao)(nil)

func NewSystemConfigDao(db *gorm.DB) *SystemConfigDao {
	return &SystemConfigDao{db: db}
}

func (dao *SystemConfigDao) Get(ctx context.Context, name string) (*model.Config, error) {
	var config model.Config
	if err := dao.db.WithContext(ctx).Where("name = ?", name).First(&config).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &config, nil
}
