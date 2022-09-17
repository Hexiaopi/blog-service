package dao

import (
	"context"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/hexiaopi/blog-service/global"
	"github.com/hexiaopi/blog-service/internal/entity"
	"github.com/hexiaopi/blog-service/internal/model"
)

type TagDao struct {
	db *gorm.DB
}

func NewTagDao(db *gorm.DB) *TagDao {
	return &TagDao{db: db}
}

func (dao *TagDao) Create(ctx context.Context, param *entity.Tag) error {
	tag := model.Tag{
		Name:       param.Name,
		Desc:       param.Desc,
		State:      param.State,
		Operator:   param.Operator,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	return dao.db.Create(&tag).Error
}

func (dao *TagDao) Get(ctx context.Context, id int) (*entity.Tag, error) {
	var tag model.Tag
	if err := dao.db.First(&tag, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	result := entity.Tag{
		Id:         tag.ID,
		Name:       tag.Name,
		Desc:       tag.Desc,
		State:      tag.State,
		CreateTime: tag.CreateTime.Format(global.DefaultTimeFormat),
		UpdateTime: tag.UpdateTime.Format(global.DefaultTimeFormat),
		Operator:   tag.Operator,
	}
	return &result, nil
}

func (dao *TagDao) Update(ctx context.Context, param *entity.Tag) error {
	tag := model.Tag{
		ID:         param.Id,
		Name:       param.Name,
		Desc:       param.Desc,
		State:      param.State,
		Operator:   param.Operator,
		UpdateTime: time.Now(),
	}
	return dao.db.Model(&model.Tag{}).Update(tag).Error
}

func (dao *TagDao) Delete(ctx context.Context, id int) error {
	tag := model.Tag{ID: id}
	return dao.db.Delete(&tag).Error
}

func (dao *TagDao) List(ctx context.Context, opt *entity.ListOption) ([]entity.Tag, int64, error) {
	query := dao.db
	if opt.Page >= 0 && opt.Limit > 0 {
		query = dao.db.Offset(opt.GetPageOffset()).Limit(opt.Limit)
	}
	var tags []model.Tag
	var err error
	var total int64
	if opt.Name != "" {
		query = query.Where("name = ?", opt.Name)
	}
	if opt.Sort != "" {
		query = query.Order(opt.GetSortType())
	}
	if err = query.Where("state = ?", opt.State).Find(&tags).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	result := make([]entity.Tag, len(tags))
	for i, tag := range tags {
		result[i] = entity.Tag{
			Id:         tag.ID,
			Name:       tag.Name,
			Desc:       tag.Desc,
			State:      tag.State,
			CreateTime: tag.CreateTime.Format(global.DefaultTimeFormat),
			UpdateTime: tag.UpdateTime.Format(global.DefaultTimeFormat),
			Operator:   tag.Operator,
		}
	}
	return result, total, nil
}
