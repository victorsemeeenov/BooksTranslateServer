package database

import (
	"github.com/jinzhu/gorm"
)

type WordSynonim struct {
	gorm.Model
	WordID 	  int `gorm:"name:word_id"`
	SynonimID int `gorm:"name:synonim_id"`
}