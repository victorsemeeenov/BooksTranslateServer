package database_model

import (
	"github.com/jinzhu/gorm"
)

type Sentence struct {
	gorm.Model
	Value 		string
	OrderNumber int 	 `gorm:"name:order_number"`
	ChapterID   int 	 `gorm:"name:chapter_id"`
	LanguageID  int 	 `gorm:"name:language_id"`
	Chapter		Chapter  `gorm:"association_foreignkey:chapter_id;"` 
	Language	Language `gorm:"association_foreignkey:language_id;"`
	Words       []*Word  `gorm:"many2many:words_sentences;"`
}