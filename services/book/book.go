package book

import (
	"github.com/BooksTranslateServer/models/database"
)

type Book interface {
	LoadBook(bytes []byte, extension string, name string, year int, bookCategoryID int, authorID int, languageID int, callback func(*database.Book, error))
	GetSentence(bookID int, chapterIndex int, sentenceIndex int, callback func(*database.Sentence, error))
}