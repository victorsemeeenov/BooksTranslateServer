package database

import (
	"github.com/jinzhu/gorm"
)

type Word struct {
	gorm.Model
	Value 		    string     `gorm:"name:value"`
	LanguageID	  int 	 		 `gorm:"name:language_id"`
	Translations []Translation `gorm:"many2many:words_translations"`
	Language      Language		 `gorm:"associated_foreignkey:LanguageID"`
	Sentences	    []Sentence    `gorm:"many2many:words_sentences"`
	PartOfSpeech  string 	    `gorm:"name:part_of_speech"`
	Transcription string      `gorm:"name:transcription"`
}



