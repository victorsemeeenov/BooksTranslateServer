package database

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Name 		   string
	NumberOfPages  string `gorm:"name:number_of_pages"`
    Year 		   int
    URL 		   string `gorm:"name:url"`
	BookCategoryID int `gorm:"name:book_category_id"`
	BookCategory   BookCategory `gorm:"association_foreignkey:book_category_id;"`
	Authors		   []*Author	`gorm:"many2many:books_authors;"`
}