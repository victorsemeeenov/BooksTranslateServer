package database_model

import (
	"github.com/jinzhu/gorm"
)

type Word struct {
	gorm.Model
	Value 		  string
	Transcription string
	PartOfSpeech  string 		 `gorm:"name:part_of_speech"`
	LanguageID	  int 	 		 `gorm:"name:language_id"`
	Translations  []*Translation `gorm:"many2many:words_translations"`
	Language      Language		 `gorm:"associated_foreignkey:language_id"`
	Sentences	  []*Sentence    `gorm:"many2many:words_sentences"`
	Synonims 	  []*Synonim	 `gorm:"many2many:words_synonims"`
}



