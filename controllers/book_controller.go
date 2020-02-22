package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	request2 "github.com/BooksTranslateServer/models/request"
	book2 "github.com/BooksTranslateServer/models/response/book"
	translation2 "github.com/BooksTranslateServer/models/response/translation"
	"github.com/BooksTranslateServer/services/book"
	"github.com/BooksTranslateServer/services/translation"
	"github.com/BooksTranslateServer/utils/error/types"
	"github.com/gin-gonic/gin"
	"github.com/sarulabs/di"
	"net/http"
	"strconv"
)

type BookController struct {}

func (b *BookController) GetBookList(container di.Container) func (*gin.Context) {
	return func(c *gin.Context) {
		_, ok := CheckAuth(container, c)
		if !ok {
			return
		}
		bookService, ok := container.Get("book").(book.BookService)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errors.New(types.INTERNAL_SERVER_ERROR).Error()})
			return
		}
		dbBooks, err := bookService.GetBookList()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		bookList := book2.CreateBookListResponse(dbBooks)
		c.JSON(http.StatusOK, bookList)
	}
}

func (b *BookController) GetAllSentences(container di.Container) func (ctx *gin.Context) {
	return func(c *gin.Context) {
		_, ok := CheckAuth(container, c)
		if !ok {
			return
		}
		bookService, ok := container.Get("book").(book.BookService)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errors.New(types.INTERNAL_SERVER_ERROR).Error()})
			return
		}
		bookParameter := c.Param("book_id")
		bookID, err := strconv.Atoi(bookParameter)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "wrong parameter: must be number"})
		}
		allSentences, err := bookService.GetAllSentence(bookID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		response := book2.MakeAllSentenceResponse(allSentences)
		bodyBytes := new(bytes.Buffer)
		err = json.NewEncoder(bodyBytes).Encode(response)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		contentDisposition := fmt.Sprintf(`attachment; filename="%d.json"`, bookID)
		extraHeaders := map[string]string{
			"Content-Disposition": contentDisposition,
		}
		c.DataFromReader(http.StatusOK, int64(bodyBytes.Len()), "multipart/form-data", bodyBytes, extraHeaders)
	}
}

func (b *BookController) TranslateSentence(container di.Container) func (ctx *gin.Context) {
	return func(c *gin.Context) {
		_, ok := CheckAuth(container, c)
		if !ok {
			return
		}
		translateService, ok := container.Get("translation").(translation.TranslationService)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errors.New(types.INTERNAL_SERVER_ERROR).Error()})
			return
		}
		var request request2.TranslateSentenceRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		sentence, err := translateService.TranslateSentence(request.SentenceID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		response := translation2.MakeSentenceTranslationResponse(sentence.Translations)
		c.JSON(http.StatusOK, response)
	}
}

func (b *BookController) TranslateWord(container di.Container) func (ctx *gin.Context) {
	return func(c *gin.Context) {
		_, ok := CheckAuth(container, c)
		if !ok {
			return
		}
		translateService, ok := container.Get("translation").(translation.TranslationService)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errors.New(types.INTERNAL_SERVER_ERROR).Error()})
			return
		}
		var request request2.TranslateWordRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		words, err := translateService.TranslateWord(request.Text, request.Lang)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		response := translation2.MakeTranslateWordResponse(words)
		c.JSON(http.StatusOK, response)
	}
}