package data

import (
	"github.com/jinzhu/gorm"
)

type BookCategory struct {
	gorm.Model
	Value string
}

func (b *BookCategory) TableName() string  {
	return "book_categories"
}