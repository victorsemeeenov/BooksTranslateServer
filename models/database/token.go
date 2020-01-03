package database

import (
	"github.com/jinzhu/gorm"
	"time"
)


type AccessToken struct {
	gorm.Model
	Value     string
	UserID    uint	`gorm:"name:user_id"`
	ExpiredIn time.Time `gorm:"name:expired_in"`
	User	  User `gorm:"associated_foreignkey:UserID;"`
}

func (a *AccessToken) TableName() string {
	return "access_tokens"
}

type RefreshToken struct {
	gorm.Model
	Value     string
	UserID    uint	`gorm:"name:user_id"`
	ExpiredIn time.Time `gorm:"name:expired_in"`
	User	  User `gorm:"associated_foreignkey:UserID;"`
}

func (r *RefreshToken) TableName() string  {
	return "refresh_tokens"
}