package book

import (
	"github.com/BooksTranslateServer/models/database"
	"os"
)

type Book interface {
	LoadBook(bytes []byte, extension string, name string) (*os.File, *string, error)
	CreateSentences(bookID int, fileURL string, languageID int) (error)
	GetSentence(sentenceID int) (*database.Sentence, error)
	GetBookList() ([]database.Book, error)
	GetAllSentence(bookID int) ([]database.Sentence, error)
}