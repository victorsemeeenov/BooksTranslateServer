package auth

import (
	. "github.com/BooksTranslateServer/models/database_models"
	. "github.com/BooksTranslateServer/models/request_models"
)

type Auth interface {
	RegisterUser(RegisterUser) (RefreshToken, AccessToken, error)
	LoginUser(RegisterUser) (RefreshToken, AccessToken, error)
	AuthorizeUser(token string) (User, error)
	CheckToken(token string) (AccessToken, error)
}