package main

import (
	"github.com/BooksTranslateServer/services/book"
	"github.com/sarulabs/di"
	"github.com/BooksTranslateServer/services"
	"github.com/BooksTranslateServer/services/logging"
	"github.com/BooksTranslateServer/data"
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
	admin := data.RegisterAdmin(container.Get("book").(book.Book))
	Route(container, admin)
}