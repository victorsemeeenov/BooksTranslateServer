package translate_storage

import (
	"github.com/BooksTranslateServer/models/database"
	"github.com/BooksTranslateServer/models/third_api/response"
)

type TranslateStorage interface {
	GetWordTranslation(word string, language string, callback func (*database.Word, *database.Language, []error))
	GetTextTranslation(text string, language string, callback func (*database.Sentence, *database.Language, []error))
	SaveWordTranslation(res response.TranslateWord, lang database.Language, callback func (*database.Word, []error))
	SaveSentenceTranslation(res response.TranslateSentence, lang database.Language, callback func(*database.Sentence, []error))
}