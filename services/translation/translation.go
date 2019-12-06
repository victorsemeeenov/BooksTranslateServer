package translation

import (
	"github.com/BooksTranslateServer/models/database"
)

type Translation interface {
	TranslateWord(word string, language string, callback func(*database.Word, error))
	TranslateSentence(sentenceIndex int, chapterIndex int, bookID int, language string, callback func(*database.Sentence, error))
}