package main

import (
	"errors"
	"github.com/BooksTranslateServer/services"
	"github.com/BooksTranslateServer/services/adminpanel"
	"github.com/BooksTranslateServer/services/logging"
	"github.com/sarulabs/di"
)

func main() {
	defer logging.Logger.Sync()

	builder, err := di.NewBuilder()
	if err != nil {
		logging.Logger.Fatal(err.Error())
	}

	err = builder.Add(services.Services...)
	if err != nil {
		logging.Logger.Fatal(err.Error())
	}

	container := builder.Build()
	defer container.Delete()
	panel, ok := container.Get("adminpanel").(adminpanel.AdminPanel)
	if !ok {
		logging.Logger.Fatal(errors.New("cant get admin panel").Error())
	} else {
		admin := panel.Register()
		Route(container, admin)
	}
}