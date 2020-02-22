package controllers

import (
	"errors"
	"github.com/BooksTranslateServer/models/request"
	auth2 "github.com/BooksTranslateServer/models/response/auth"
	"github.com/BooksTranslateServer/services/auth"
	"github.com/BooksTranslateServer/utils/error/types"
	"github.com/gin-gonic/gin"
	"github.com/sarulabs/di"
	"net/http"
)

type AuthController struct{}

func (a *AuthController) Register(container di.Container) func(*gin.Context) {
	return func(c *gin.Context) {
		auth, ok := container.Get("auth").(auth.AuthService)
		if !ok {
			c.JSON(http.StatusOK, gin.H{"error": errors.New(types.INTERNAL_SERVER_ERROR).Error()})
			return
		}
		var user request.RegisterUser
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		refreshToken, accessToken, error := auth.RegisterUser(user)

		if error != nil {
			c.JSON(http.StatusOK, gin.H{"error": error.Error()})
			return
		}

		response := auth2.CreateAuthResponse(accessToken.Value, refreshToken.Value, accessToken.ExpiredIn)
		c.JSON(http.StatusOK, response)
	}
}

func (a *AuthController) Login(container di.Container) func(*gin.Context) {
	return func (c* gin.Context) {
		auth, ok := container.Get("auth").(auth.AuthService)
		if !ok {
			c.JSON(http.StatusOK, gin.H{"error": errors.New(types.INTERNAL_SERVER_ERROR).Error()})
			return
		}
		var user request.LoginUser
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	
		refreshToken, accessToken, error := auth.LoginUser(user)
		if error != nil {
			c.JSON(http.StatusOK, gin.H{"error": error.Error()})
			return
		}
	
		response := auth2.CreateAuthResponse(accessToken.Value, refreshToken.Value, accessToken.ExpiredIn)
		c.JSON(http.StatusOK, response)
	}
}
