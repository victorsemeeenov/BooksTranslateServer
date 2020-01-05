package database

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Name 		   		 string
	NumberOfPages  int `gorm:"name:number_of_pages"`
	Year 		   		 int
	URL 		   		 string `gorm:"name:url"`
	LanguageID 		 int `gorm:"name:language_id"`
	BookCategoryID int `gorm:"name:book_category_id"`
	BookCategory   BookCategory `gorm:"associated_foreignkey:BookCategoryID;"`
	Authors		   	 []*Author	`gorm:"many2many:books_authors;"`
	Language	     Language `gorm:"associated_foreignkey:LanguageID;"`
}