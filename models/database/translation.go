package database

import (
	"github.com/jinzhu/gorm"
)

type Translation struct {
	gorm.Model
	Value 		   string `gorm:"key:value"`
	PartOfSpeech string `gorm:"key:part_of_speech"`
	Gender		 	 string
	Synonims     []Synonim `gorm:"foreignkey:TranslationID"`
	Words        []Word		 `gorm:"many2many:words_translations"`
}

