package dao

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/hexiaopi/blog-service/internal/app"
	"github.com/hexiaopi/blog-service/internal/entity"
	"github.com/hexiaopi/blog-service/internal/model"
)

type TagDao struct {
	db *gorm.DB
}

func NewTagDao(db *gorm.DB) *TagDao {
	return &TagDao{db: db}
}

func (dao *TagDao) Get(ctx context.Context, id int) (*model.Tag, error) {
	tag := model.Tag{
		ID: id,
	}
	if err := dao.db.First(&tag).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
	}
	return &tag, nil
}

func (dao *TagDao) Create(ctx context.Context, param *entity.Tag) error {
	tag := model.Tag{
		Name:      param.Name,
		State:     param.State,
		CreatedBy: param.CreatedBy,
	}
	return dao.db.Create(&tag).Error
}

func (dao *TagDao) Update(ctx context.Context, param *entity.Tag) error {
	tag := model.Tag{
		ID:        param.Id,
		Name:      param.Name,
		CreatedBy: param.CreatedBy,
		State:     param.State,
	}
	return dao.db.Save(tag).Error
}

func (dao *TagDao) Delete(ctx context.Context, id int) error {
	tag := model.Tag{ID: id}
	return dao.db.Where("is_del = ?", 0).Delete(&tag).Error
}

func (dao *TagDao) List(ctx context.Context, opt *entity.ListOption) ([]model.Tag, int64, error) {
	var tags []model.Tag
	var err error
	var total int64
	if opt.PageNum >= 0 && opt.PageSize > 0 {
		dao.db = dao.db.Offset(app.GetPageOffset(opt.PageNum, opt.PageSize)).Limit(opt.PageSize)
	}
	if opt.Name != "" {
		dao.db = dao.db.Where("name = ?", opt.Name)
	}
	dao.db = dao.db.Where("state = ?", opt.State)
	if err = dao.db.Where("is_del = ?", 0).Find(&tags).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return tags, total, nil
}
