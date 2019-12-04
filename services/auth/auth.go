package auth

import (
	. "github.com/BooksTranslateServer/models/database"
	. "github.com/BooksTranslateServer/models/request"
)

type Auth interface {
	RegisterUser(RegisterUser) (RefreshToken, AccessToken, error)
	LoginUser(LoginUser) (RefreshToken, AccessToken, error)
	AuthorizeUser(token string) (User, error)
	CheckToken(token string) (AccessToken, error)
}