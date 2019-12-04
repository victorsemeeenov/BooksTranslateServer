package main

import (
	"github.com/sarulabs/di"
	"github.com/gin-gonic/gin"
	"github.com/BooksTranslateServer/controllers"
)

func Route(c di.Container) {
	r := gin.Default()
	api := r.Group("/api") 
	var authController controllers.AuthController
	{
		api.POST("/register", authController.Register(c))
		api.POST("/login", authController.Login(c))
	}
	r.Run(":8080")
}