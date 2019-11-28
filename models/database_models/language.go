package database_model

import (
	"github.com/jinzhu/gorm"
)

type Language struct {
	gorm.Model
	Value string
}