package mysql

import (
	"fmt"

	"gorm.io/gorm"
)

type Option func(*gorm.DB)

func WithId(id int) Option {
	return func(db *gorm.DB) {
		db.Where("id = ?", id)
	}
}

func WithName(name string) Option {
	return func(db *gorm.DB) {
		db.Where("name = ?", name)
	}
}

func WithOffset(offset int) Option {
	return func(db *gorm.DB) {
		db.Offset(offset)
	}
}

func WithLimit(limit int) Option {
	return func(db *gorm.DB) {
		db.Limit(limit)
	}
}

func WithOrderDesc(field string) Option {
	return func(db *gorm.DB) {
		db.Order(fmt.Sprintf("%s DESC", field))
	}
}

func WithOrderAsc(field string) Option {
	return func(db *gorm.DB) {
		db.Order(fmt.Sprintf("%s ASC", field))
	}
}
