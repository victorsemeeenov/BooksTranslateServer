package services

import (
	"github.com/sarulabs/di"
	"github.com/BooksTranslateServer/services/auth"
	"github.com/BooksTranslateServer/services/network/translate_api"
	"github.com/BooksTranslateServer/services/book"
)

var Services = []di.Def{
	{
		Name:  "auth",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return auth.AuthService{}, nil
		},
	},
	{
		Name: "translate_api",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return translate_api.GetYandexService(), nil
		},
	},
	{
		Name: "book",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return book.BookService{}, nil
		},
	},
}