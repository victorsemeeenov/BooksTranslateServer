package data

import (
	"github.com/jinzhu/gorm"
)

type Author struct {
	gorm.Model
	Name   string
	BookID string  `gorm:"name:book_id"`
	Books  []*Book `gorm:"many2many:books_authors;"`
}
