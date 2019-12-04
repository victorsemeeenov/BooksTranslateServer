package services

import (
	"github.com/sarulabs/di"
	"github.com/BooksTranslateServer/services/auth"
)

var Services = []di.Def{
	{
		Name:  "auth",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return auth.AuthService{}, nil
		},
	},
}