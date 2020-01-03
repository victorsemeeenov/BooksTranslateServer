package database

import (
	"github.com/jinzhu/gorm"
)

type Sentence struct {
	gorm.Model
	Value 		string
	OrderNumber int 	 `gorm:"name:order_number"`
	ChapterID   int 	 `gorm:"name:chapter_id"`
	LanguageID  int 	 `gorm:"name:language_id"`
	Chapter		Chapter  `gorm:"associated_foreignkey:ChapterID;"`
	Language	Language `gorm:"associated_foreignkey:LanguageID;"`
	Words       []*Word  `gorm:"many2many:words_sentences;"`
}