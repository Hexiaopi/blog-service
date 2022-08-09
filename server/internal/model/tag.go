package model

import "github.com/jinzhu/gorm"

type Tag struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	State      uint8  `json:"state"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}

func (t Tag) Create(db *gorm.DB) (*Tag, error) {
	if err := db.Create(&t).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

func (t Tag) Get(db *gorm.DB) (Tag, error) {
	var tag Tag
	err := db.Where("id = ? AND is_del = ? AND state = ?", t.ID, 0, t.State).First(&tag).Error
	if err != nil {
		return tag, err
	}
	return tag, nil
}

func (t Tag) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&t).Where("is_del = ?", 0).Update(values).Error
}

func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("is_del = ?", 0).Delete(&t).Error
}

func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}
