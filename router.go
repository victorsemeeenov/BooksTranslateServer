package main

import (
	"github.com/sarulabs/di"
	"github.com/gin-gonic/gin"
	"github.com/BooksTranslateServer/controllers"
	"github.com/qor/admin"
	"net/http"
)

func Route(c di.Container, admin *admin.Admin) {
	r := gin.Default()
	api := r.Group("/api")
	{
	 var authController controllers.AuthController
	 api.POST("/register", authController.Register(c))
	 api.POST("/login", authController.Login(c))

	 var bookController controllers.BookController
	 api.GET("books", bookController.GetBookList(c))
	 api.GET("books/:book_id", bookController.GetAllSentences(c))
	 api.GET("translation/sentence", bookController.TranslateSentence(c))
	 api.GET("translation/word", bookController.TranslateWord(c))
	}
	adminMux := http.NewServeMux()
	admin.MountTo("/admin", adminMux)
	r.Any("/admin/*resources", gin.WrapH(adminMux))
	r.Run("localhost:8080")
}