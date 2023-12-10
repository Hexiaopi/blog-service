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

type OperationDao struct {
	db *gorm.DB
}

var _ store.OperationStore = (*OperationDao)(nil)

func NewOperationDao(db *gorm.DB) *OperationDao {
	return &OperationDao{db: db}
}

func (dao *OperationDao) Create(ctx context.Context, log *entity.OperationLog) error {
	l := log.ToModel()
	l.CreateTime = time.Now()
	return dao.db.WithContext(ctx).Create(&l).Error
}

func (dao *OperationDao) Get(ctx context.Context, id int) (*entity.OperationLog, error) {
	var log model.OperationLog
	if err := dao.db.WithContext(ctx).First(&log, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	result := entity.ToEntityOperation(&log)
	return result, nil
}

func (dao *OperationDao) Update(ctx context.Context, log *entity.OperationLog) error {
	l := log.ToModel()
	return dao.db.WithContext(ctx).Updates(l).Error
}

func (dao *OperationDao) Delete(ctx context.Context, id int) error {
	log := model.OperationLog{ID: id}
	return dao.db.WithContext(ctx).Delete(&log).Error
}

func (dao *OperationDao) List(ctx context.Context, opt *entity.ListOption) ([]model.OperationLog, error) {
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
	if opt.Result != "" {
		query = query.Where("result = ?", opt.Result)
	}
	if opt.Sort != "" {
		query = query.Order(opt.GetSortType())
	}
	if err := query.Model(&model.OperationLog{}).
		Find(&logs).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
	}
	return logs, nil
}

func (dao *OperationDao) Count(ctx context.Context, opt *entity.ListOption) (int64, error) {
	query := dao.db.WithContext(ctx)
	var total int64
	if opt.Object != "" {
		query = query.Where("object = ?", opt.Object)
	}
	if opt.Action != "" {
		query = query.Where("action = ?", opt.Action)
	}
	if opt.UserId > 0 {
		query = query.Where("user_id = ?", opt.UserId)
	}
	if opt.Result != "" {
		query = query.Where("result = ?", opt.Result)
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
