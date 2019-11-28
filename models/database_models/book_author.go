package database_model

import (
	"github.com/jinzhu/gorm"
)

type BookAuthor struct {
	gorm.Model
	bookID   int `gorm:"name:book_id"`
	authorID int `gorm:"name:author_id"`
}

func (b *BookAuthor) TableName() string {
	return "books_authors"
}