package auth

import (
	"github.com/jinzhu/gorm"
	"crypto/sha1"
	. "github.com/BooksTranslateServer/models/database"
	. "github.com/BooksTranslateServer/models/request"
	. "github.com/BooksTranslateServer/data"
	"errors"
	"github.com/BooksTranslateServer/services/logging"
	"github.com/BooksTranslateServer/utils/error/types"
	."github.com/BooksTranslateServer/utils/error"
	"time"
	"encoding/base64"
)

type AuthService struct {}

func passwordHash(password string) string {
	hasher := sha1.New()
	var bytes []byte
	hasher.Write(bytes)
	return base64.StdEncoding.EncodeToString(hasher.Sum(nil))
}

func (a *AuthService) RegisterUser(userRequest RegisterUser) (refreshToken RefreshToken, accessToken AccessToken, err error)  {
	passwordHash := passwordHash(userRequest.Password)
	newUser := User{
		Name: userRequest.Username,
		Email: userRequest.Email,
		Password: string(passwordHash)}
	var user User
	err = Throw(Db.Where("name = ?", newUser.Name).First(&user))
	if err == nil {
		err = errors.New(types.USER_EXIST_WITH_NAME)
		return
	}
	err = Throw(Db.Where("email = ?", newUser.Email).First(&user))
	if err == nil {
		err = errors.New(types.USER_EXIST_WITH_EMAIL)
		return
	}
	tx := Db.Begin()
	defer func () {
		if err != nil {
			logging.Logger.Error(err.Error())
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

func (a *AuthService) LoginUser(userRequest LoginUser) (refreshToken RefreshToken, accessToken AccessToken, err error) {
	passwordHash := passwordHash(userRequest.Password)
	var user User
	err = Throw(Db.Where("password = ?", passwordHash).Where("name = ?", userRequest.Username).First(&user))
	Db.Model(&user).Related(&user.RefreshTokens)
	Db.Model(&user).Related(&user.AccessTokens)
	if err != nil {
		logging.Logger.Error(err.Error())
		err = errors.New(types.USER_NOT_FOUND)
		return
	} 
	
	if len(user.RefreshTokens) == 0 {
		err = errors.New(types.REFRESH_TOKEN_NOT_FINDED)
		return
	}
	rt := user.RefreshTokens[0]
	
	if len(user.AccessTokens) == 0 {
		err = errors.New(types.ACCESS_TOKEN_NOT_FINDED)
		return
	}
	at := user.AccessTokens[0]
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
	refreshToken, accessToken, err =  createNewTokens(user, tx)
	if tx.Error != nil {
		err = tx.Error
	}
	return
}

func (a *AuthService) CheckToken(token string) (accessToken AccessToken, err error)  {
	Db.Where("value = ?", token).First(&accessToken)
	if &accessToken == nil {
		err = errors.New(types.ACCESS_TOKEN_NOT_FINDED)
		return
	}
	if accessToken.ExpiredIn.Sub(time.Now()) < 0 {
		err = errors.New(types.ACCESS_TOKEN_EXPIRED_IN)
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
		err = errors.New(types.USER_NOT_FOUND)
	}
	return
}

func generateToken(days int, u User) (tokenString string, err error, expiredIn time.Time) {
	day := time.Hour * 24
	createTime := time.Now()
	expiredIn = createTime.Add(day * time.Duration(days))
	tokenJwt := UserJWT {
		ID: u.ID,
		Name: u.Name,
		CreatedAt: createTime,
		ExpiresAt: expiredIn,
	}
	tokenString, err = tokenJwt.TokenString()
	return 
}

func createNewTokens(u User, tx *gorm.DB) (refreshToken RefreshToken, accessToken AccessToken, err error) {
	refreshTokenString, err, expiredIn := generateToken(30, u)
	if err != nil {
		return
	}
	refreshToken = RefreshToken {
		UserID: u.ID,
		Value: refreshTokenString,
		ExpiredIn: expiredIn,
	}
	tx.Create(&refreshToken)
 	accessTokenString, err, expiredIn := generateToken(15, u)
	if err != nil {
		return 
	}
	accessToken = AccessToken {
		UserID: u.ID,
		Value: accessTokenString,
		ExpiredIn: expiredIn,
	}
	tx.Create(&accessToken)
	return
}