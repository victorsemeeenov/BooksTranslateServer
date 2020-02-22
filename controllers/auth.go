package controllers

import (
	"errors"
	"github.com/BooksTranslateServer/models/database"
	"github.com/BooksTranslateServer/services/auth"
	"github.com/BooksTranslateServer/utils/error/types"
	"github.com/gin-gonic/gin"
	"github.com/sarulabs/di"
	"net/http"
	"strings"
)

func CheckAuth(container di.Container, c *gin.Context) (*database.User, bool) {
	auth, ok := container.Get("auth").(auth.AuthService)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": errors.New(types.INTERNAL_SERVER_ERROR).Error()})
		return nil, false
	}
	bearer := c.Request.Header.Get("Authorization")
	accessToken := strings.ReplaceAll(bearer, "Bearer ", "")
	user, err := auth.AuthorizeUser(accessToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return &user, false
	}
	return &user, true
}