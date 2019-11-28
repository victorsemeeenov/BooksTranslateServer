package data

import (
	"github.com/jinzhu/gorm"
)

type Chapter struct {
	gorm.Model
	Title 		string
	OrderNumber int  `gorm:"name:order_number"`
	OrderValue  int  `gorm:"name:order_value"`
	BookID      uint `gorm:"name:book_id"`
	Book		Book `gorm:"association_foreignkey:book_id;"`
}