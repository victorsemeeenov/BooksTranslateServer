package translation

import (
	"github.com/BooksTranslateServer/models/database"
)

type Translation interface {
	TranslateWord(word string, language string) (*database.Word, error)
	TranslateSentence(sentenceID int) (*database.Sentence, error)
}