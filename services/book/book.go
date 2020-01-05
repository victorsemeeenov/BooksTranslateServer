package book

import (
	"github.com/BooksTranslateServer/models/database"
	"os"
)

type Book interface {
	LoadBook(bytes []byte, extension string, name string) (*os.File, *string, error)
	CreateSentences(bookID int, fileURL string, languageID int) (error)
	GetSentence(bookID int, chapterIndex int, sentenceIndex int, callback func(*database.Sentence, error))
}