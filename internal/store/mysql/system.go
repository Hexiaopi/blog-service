package mysql

import (
	"context"

	"gorm.io/gorm"

	"github.com/hexiaopi/blog-service/internal/model"
)

type SystemConfigDao struct {
	db *gorm.DB
}

func NewSystemConfigDao(db *gorm.DB) *SystemConfigDao {
	return &SystemConfigDao{db: db}
}

func (dao *SystemConfigDao) Get(ctx context.Context, name string) (*model.SystemConfig, error) {
	var config model.SystemConfig
	if err := dao.db.WithContext(ctx).Where("name = ?", name).First(&config).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &config, nil
}
