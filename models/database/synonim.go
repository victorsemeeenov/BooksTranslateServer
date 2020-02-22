package database

import (
	"github.com/jinzhu/gorm"
)

type Synonim struct {
	gorm.Model
	TranslationID int `gorm:"name:translation_id"`
	Value  		 string 
	PartOfSpeech string  `gorm:"name:part_of_speech"`
	Gender		 string
	Translation Translation `gorm:"associated_foreignkey:TranslationID;"`
}