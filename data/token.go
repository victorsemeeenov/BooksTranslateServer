package data

import (
	"github.com/jinzhu/gorm"
	"time"
)

type AccessToken struct {
	gorm.Model
	Value     string
	UserID    string	`gorm:"name:user_id"`
	ExpiredIn time.Time
}

func (a *AccessToken) TableName() string {
	return "access_tokens"
}

type RefreshToken struct {
	gorm.Model
	Value     string
	UserID    string	`gorm:"name:user_id"`
	ExpiredIn time.Time
}

func (r *RefreshToken) TableName() string  {
	return "refresh_tokens"
}