package model

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/hexiaopi/blog-service/internal/pkg/auth"
)

type User struct {
	ID           uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	Name         string `gorm:"column:name" json:"username"`
	PassWord string `gorm:"column:password" json:"password"`
	CreatedBy    string `json:"created_by"`
	ModifiedBy   string `json:"modified_by"`
	CreatedOn    uint32 `json:"created_on"`
	ModifiedOn   uint32 `json:"modified_on"`
	DeletedOn    uint32 `json:"deleted_on"`
	IsDel        uint8  `json:"is_del"`
}

func (u *User) TableName() string {
	return "blog_user"
}

func (u *User) Get(db *gorm.DB) (*User, error) {
	var user User
	db = db.Where("name = ? AND is_del = ?", u.Name, 0)
	if err := db.First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) Compare(password string) error {
	if err := auth.Compare(u.PassWord, password); err != nil {
		return fmt.Errorf("failed to compare password:%s and:%s :%v", u.PassWord, password, err)
	}
	return nil
}
