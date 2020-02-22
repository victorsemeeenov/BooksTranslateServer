package services

import (
	"github.com/BooksTranslateServer/services/adminpanel"
	"github.com/BooksTranslateServer/services/translate_storage"
	"github.com/BooksTranslateServer/services/translation"
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
		Name: "book",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return book.BookService{}, nil
		},
	},
	{
			Name: "adminpanel",
			Scope: di.App,
			Build: func(ctn di.Container) (i interface{}, err error) {
				return adminpanel.AdminPanel{ctn.Get("book").(book.BookService)}, nil
			},
	},
	{
				Name: "translation",
				Scope: di.App,
				Build: func(ctn di.Container) (i interface{}, err error) {
					return translation.TranslationService {
						TranslateAPI:     translate_api.GetYandexService(),
						TranslateStorage: translate_storage.TranslateDB{},
						BookService:      book.BookService{},
					}, nil
				},
	},
}