package mysql

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/hexiaopi/blog-service/internal/model"
)

type OperationDao struct {
	db *gorm.DB
}

func NewOperationDao(db *gorm.DB) *OperationDao {
	return &OperationDao{db: db}
}

func (dao *OperationDao) Create(ctx context.Context, log *model.OperationLog) error {
	log.CreateTime = time.Now()
	return dao.db.WithContext(ctx).Create(log).Error
}

func (dao *OperationDao) Get(ctx context.Context, id int) (*model.OperationLog, error) {
	var log model.OperationLog
	if err := dao.db.WithContext(ctx).First(&log, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &log, nil
}

func (dao *OperationDao) Update(ctx context.Context, log *model.OperationLog) error {
	return dao.db.WithContext(ctx).Updates(log).Error
}

func (dao *OperationDao) Delete(ctx context.Context, id int) error {
	log := model.OperationLog{ID: id}
	return dao.db.WithContext(ctx).Delete(&log).Error
}

func (dao *OperationDao) List(ctx context.Context, opt *model.ListOption) ([]model.OperationLog, error) {
	query := dao.db.WithContext(ctx)
	if opt.Page >= 0 && opt.Limit > 0 {
		query = query.Offset(opt.GetPageOffset()).Limit(opt.Limit)
	}
	var logs []model.OperationLog
	if opt.Object != "" {
		query = query.Where("object = ?", opt.Object)
	}
	if opt.Action != "" {
		query = query.Where("action = ?", opt.Action)
	}
	if opt.UserId > 0 {
		query = query.Where("user_id = ?", opt.UserId)
	}
	if opt.Sort != "" {
		query = query.Order(opt.GetSortType())
	}
	if err := query.Model(&model.OperationLog{}).
		Preload("User").
		Find(&logs).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
	}
	return logs, nil
}

func (dao *OperationDao) Count(ctx context.Context, opt *model.ListOption) (int64, error) {
	query := dao.db.WithContext(ctx)
	if opt.Page >= 0 && opt.Limit > 0 {
		query = query.Offset(opt.GetPageOffset()).Limit(opt.Limit)
	}
	var total int64
	if opt.Name != "" {
		query = query.Where("name = ?", opt.Name)
	}
	if opt.Sort != "" {
		query = query.Order(opt.GetSortType())
	}
	if err := query.Model(&model.OperationLog{}).
		Count(&total).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, nil
		}
	}
	return total, nil
}
