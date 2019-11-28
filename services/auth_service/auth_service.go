package auth

import (
	"github.com/jinzhu/gorm"
	"crypto/sha1"
	"io"
	. "github.com/BooksTranslateServer/models/database_models"
	. "github.com/BooksTranslateServer/models/request_models"
	. "github.com/BooksTranslateServer/data"
	"errors"
	"github.com/BooksTranslateServer/utils/error_types"
	"time"
)

type AuthService struct {}

func passwordHash(password string) string {
	h := sha1.New()
	io.WriteString(h, password)
	return string(h.Sum(nil))
}

func (a *AuthService) RegisterUser(userRequest RegisterUser) (refreshToken RefreshToken, accessToken AccessToken, err error)  {
	passwordHash := passwordHash(userRequest.Password)
	newUser := User{
		Name: userRequest.Username,
		Email: userRequest.Email,
		Password: string(passwordHash)}
	var user *User
	Db.Where("name = ?", newUser.Name).First(user)
	if user != nil {
		err = errors.New(error_types.USER_EXIST_WITH_NAME)
		return
	}
	Db.Where("email = ?", newUser.Email).First(user)
	if user != nil {
		err = errors.New(error_types.USER_EXIST_WITH_EMAIL)
	}
	tx := Db.Begin()
	defer func () {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	if err = tx.Error; err != nil {
		return 
	}
	tx.Create(&newUser)
	refreshToken, accessToken, err =  createNewTokens(newUser, tx)
	return 
}

func (a *AuthService) LoginUser(userRequest RegisterUser) (refreshToken RefreshToken, accessToken AccessToken, err error) {
	passwordHash := passwordHash(userRequest.Password)
	var user *User
	Db.Where("password = ?", passwordHash).First(user)
	if user == nil {
		err = errors.New(error_types.USER_NOT_FOUND)
		return
	} 
	
	if len(user.RefreshTokens) == 0 {
		err = errors.New(error_types.REFRESH_TOKEN_NOT_FINDED)
		return
	}
	rt := *user.RefreshTokens[0]
	
	if len(user.AccessTokens) == 0 {
		err = errors.New(error_types.ACCESS_TOKEN_NOT_FINDED)
		return
	}
	at := *user.AccessTokens[0]
	tx := Db.Begin()
	defer func () {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	if err = tx.Error; err != nil {
		return 
	}
	tx.First(&RefreshToken{}, rt.ID).Delete(&RefreshToken{})
	tx.First(&AccessToken{}, at.ID).Delete(&AccessToken{})
	refreshToken, accessToken, err =  createNewTokens(*user, tx)
	return
}

func (a *AuthService) CheckToken(token string) (accessToken AccessToken, err error)  {
	Db.Where("value = ?", token).First(&accessToken)
	if &accessToken == nil {
		err = errors.New(error_types.ACCESS_TOKEN_NOT_FINDED)
		return
	}
	if accessToken.ExpiredIn.Sub(time.Now()) < 0 {
		err = errors.New(error_types.ACCESS_TOKEN_EXPIRED_IN)
	}
	return
}

func (a *AuthService) AuthorizeUser(token string) (user User, err error)  {
	accessToken, err := a.CheckToken(token)
	if err != nil {
		return 
	}
	user = accessToken.User
	if &user == nil {
		err = errors.New(error_types.USER_NOT_FOUND)
	}
	return
}

func createNewTokens(u User, tx *gorm.DB) (refreshToken RefreshToken, accessToken AccessToken, err error) {
	refreshTokenJWT := UserJWT{
		ID: u.ID,
		Date: time.Now(),
		Secret: "refreshToken"}
	refreshTokenString, err := refreshTokenJWT.TokenString()
	refreshToken = RefreshToken {
		UserID: u.ID,
		Value: refreshTokenString}
	Db.Create(&refreshToken)
	accessTokenJWT := UserJWT{
		ID: u.ID,
		Date: time.Now(),
		Secret: "accessToken"}
	accessTokenString, err := accessTokenJWT.TokenString()
	accessToken = AccessToken {
		UserID: u.ID,
		Value: accessTokenString}
	return
}