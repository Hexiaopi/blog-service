package model

import "github.com/jinzhu/gorm"

type Auth struct {
	ID         uint32 `json:"id"`
	AppKey     string `json:"app_key"`
	AppSecret  string `json:"app_secret"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func (a Auth) TableName() string {
	return "blog_auth"
}

func (a Auth) Get(db *gorm.DB) (*Auth, error) {
	var auth Auth
	db = db.Where("app_key = ? AND app_secret = ? AND is_del = ?", a.AppKey, a.AppSecret, 0)
	if err := db.First(&auth).Error; err != nil {
		return nil, err
	}
	return &auth, nil
}
