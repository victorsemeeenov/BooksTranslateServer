package database

import (
	"github.com/jinzhu/gorm"
)

type Chapter struct {
	gorm.Model
	Title 		string
	OrderNumber int  `gorm:"name:order_number"`
	OrderValue  string  `gorm:"name:order_value"`
	BookID      uint `gorm:"name:book_id"`
	Book		Book `gorm:"associated_foreignkey:BookID;"`
}