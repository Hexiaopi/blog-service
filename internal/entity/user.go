package entity

import (
	"fmt"
	"time"

	"github.com/hexiaopi/blog-service/internal/model"
	"github.com/hexiaopi/blog-service/internal/pkg/auth"
)

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	PassWord   string `json:"-"`
	Avatar     string `json:"avatar"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
	State      uint8  `json:"state"`
	Roles      []Role `json:"roles"`
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

func (u *User) ToModel() *model.User {
	createTime, _ := time.Parse(DefaultTimeFormat, u.CreateTime)
	updateTime, _ := time.Parse(DefaultTimeFormat, u.UpdateTime)
	return &model.User{
		ID:         u.ID,
		Name:       u.Name,
		PassWord:   u.PassWord,
		Avatar:     u.Avatar,
		CreateTime: createTime,
		UpdateTime: updateTime,
		State:      u.State,
	}
}

func ToEntityUser(user *model.User) *User {
	result := User{
		ID:         user.ID,
		Name:       user.Name,
		Avatar:     user.Avatar,
		PassWord:   user.PassWord,
		CreateTime: user.CreateTime.Format(DefaultTimeFormat),
		UpdateTime: user.UpdateTime.Format(DefaultTimeFormat),
		State:      user.State,
	}
	return &result
}
