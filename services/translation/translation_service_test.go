package translation

import (
	"github.com/BooksTranslateServer/services/translate_storage"
	"testing"
	"github.com/BooksTranslateServer/services/network/translate_api"
	"github.com/BooksTranslateServer/services/book"
	"github.com/BooksTranslateServer/config"
	"github.com/BooksTranslateServer/data"
	"os"
)

var translationService *TranslationService


func setUp() {
	data.InitWithConfig(config.TEST)
}

func tearDown() {
	data.RemoveAll()
}

func TestMain(m *testing.M) { 
	setUp()
	code := m.Run() 
	tearDown() 
	os.Exit(code)
}

func TestWordTranslation(t *testing.T) {
	translationService = &TranslationService{
		TranslateAPI: translate_api.GetYandexService(),
		TranslateStorage: translate_storage.TranslateDB{},
		BookService: book.BookService{},
	}
}