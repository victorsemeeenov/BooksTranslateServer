package data

import (
	"github.com/jinzhu/gorm"
)

type Synonim struct {
	gorm.Model
	WordID 		 int 	 `gorm:"name:word_id"`
	Value  		 string 
	PartOfSpeech string  `gorm:"name:part_of_speech"`
	Gender		 string
	Words		 []*Word `gorm:"many2many:words_synonims"`
}