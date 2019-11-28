package data

import (
	"github.com/jinzhu/gorm"
)

type Language struct {
	gorm.Model
	Value string
}