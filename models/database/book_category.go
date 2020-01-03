package database

import (
	"github.com/jinzhu/gorm"
)

type BookCategory struct {
	gorm.Model
	Name string
}

func (b *BookCategory) TableName() string  {
	return "book_categories"
}