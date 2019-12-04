package database

import (
	"github.com/jinzhu/gorm"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//User
type User struct {
	gorm.Model
	Name  	      string
	Email 	      string
	Password      string
	AccessTokens  []*AccessToken  `gorm:"foreignkey:UserID"`
	RefreshTokens []*RefreshToken `gorm:"foreignkey:UserID"`
}

type UserJWT struct {
	*jwt.StandardClaims
	ID     		uint		`json:"user_id"`
	Name   		string		`json:"username"`
	CreatedAt   time.Time	`json:"created_at"`
	ExpiresAt   time.Time   `json:"expires_at"`
}

func (u *UserJWT) TokenString() (tokenString string, err error)  {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = &UserJWT {
		StandardClaims: &jwt.StandardClaims {
			ExpiresAt: u.ExpiresAt.Unix(),
		},
		ID: u.ID,
		Name: u.Name,
		CreatedAt: u.CreatedAt,
		ExpiresAt: u.ExpiresAt,
	}
	tokenString, err = token.SignedString([]byte("secret"))
	return
}