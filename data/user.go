package data

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
	AccessTokens  []*AccessToken  `gorm:"foreignkey:user_id"`
	RefreshTokens []*RefreshToken `gorm:"foreignkey:user_id"`
}

type UserJWT struct {
	ID   int
	Name string
	Date time.Time
}

func (u *UserJWT) TokenString() (tokenString string, err error)  {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})
	//MARK: ToDo secret key
	tokenString, err = token.SignedString("secret")
	return
}