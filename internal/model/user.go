package model

import (
	"fmt"
	"time"

	"github.com/hexiaopi/blog-service/internal/pkg/auth"
)

type User struct {
	ID         int       `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	Name       string    `gorm:"column:name" json:"name"`
	PassWord   string    `gorm:"column:password" json:"-"`
	Avatar     string    `gorm:"column:avatar" json:"avatar"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel      uint8     `gorm:"column:is_del" json:"is_del"`
	Roles      []string  `gorm:"-" json:"roles"`
}

func (u *User) TableName() string {
	return "blog_user"
}

func (u *User) Compare(password string) error {
	if err := auth.Compare(u.PassWord, password); err != nil {
		return fmt.Errorf("failed to compare password:%s and:%s :%v", u.PassWord, password, err)
	}
	return nil
}

func (u *User) EncryptPassword() error {
	password, err := auth.Encrypt(u.PassWord)
	if err != nil {
		return err
	}
	u.PassWord = password
	return nil
}
