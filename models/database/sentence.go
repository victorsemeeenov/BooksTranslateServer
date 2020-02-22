package database

import (
	"github.com/jinzhu/gorm"
)

type Sentence struct {
	gorm.Model
	Value 		    string
	OrderNumber   int 	 							  `gorm:"name:order_number"`
	BookID 		    int 	 							  `gorm:"name:associated_foreignkey:BookID;"`
	ChapterID     int 	 							  `gorm:"name:chapter_id; DEFAULT:NULL"`
	LanguageID    int 	 							  `gorm:"name:language_id"`
	Chapter		    Chapter  						  `gorm:"associated_foreignkey:ChapterID;"`
	Language	    Language 						  `gorm:"associated_foreignkey:LanguageID;"`
	Translations  []SentenceTranslation `gorm:"foreignkey:TranslationID"`
	Words         []Word  						  `gorm:"many2many:words_sentences;"`
}