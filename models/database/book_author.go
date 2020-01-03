package database

import (
	"github.com/jinzhu/gorm"
)

type BookAuthor struct {
	gorm.Model
	BookID   int `gorm:"name:book_id"`
	AuthorID int `gorm:"name:author_id"`
}

func (b *BookAuthor) TableName() string {
	return "books_authors"
}