package translate_storage

import (
	"github.com/BooksTranslateServer/models/database"
	"github.com/BooksTranslateServer/models/third_api/response"
)

type TranslateStorage interface {
	GetWordTranslation(word string, language string) ([]database.Word, error)
	GetSentenceTranslation(sentenceID int) (*database.Sentence, error)
	SaveWordTranslations(res response.TranslateWordList, language string) ([]database.Word, error)
	SaveSentenceTranslation(sentenceID int,  res response.TranslateSentence, language string) (*database.Sentence, error)
}