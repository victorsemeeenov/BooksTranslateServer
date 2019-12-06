package database

import (
	"github.com/jinzhu/gorm"
)

type Translation struct {
	gorm.Model
	Value 		 string
	PartOfSpeech string `gorm:"key:part_of_speech"`
	Gender		 string
	Words 		 []Word `gorm:many2many:words_translations`
}

