package database_model

import (
	"github.com/jinzhu/gorm"
)

type Translation struct {
	gorm.Model
	Value string
	Words []*Word `gorm:many2many:words_translations`
}

