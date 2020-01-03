package database

import (
	"github.com/jinzhu/gorm"
)

type Author struct {
	gorm.Model
	Name   string
	Books  []*Book `gorm:"many2many:books_authors;"`
}
