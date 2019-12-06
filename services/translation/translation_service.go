package translation

import (
	"github.com/BooksTranslateServer/services/translate_storage"
	"github.com/BooksTranslateServer/services/network/translate_api"
	"github.com/BooksTranslateServer/models/database"
	"github.com/BooksTranslateServer/models/third_api/response"
	"github.com/BooksTranslateServer/services/logging"
	"github.com/BooksTranslateServer/services/book"
)

type TranslationService struct {
	TranslateAPI 	 translate_api.TranslateAPI
	TranslateStorage translate_storage.TranslateStorage
	BookService	     book.Book
}

func (t *TranslationService) TranslateWord(word string, language string, callback func(*database.Word, error)) {
	t.TranslateStorage.GetWordTranslation(word, language, func (dbWord *database.Word, lang *database.Language, errs []error) {
		if len(errs) > 0 {
			for _, err := range errs {
				logging.Logger.Error(err.Error())
			}
		} 
		if dbWord == nil {
			t.apiWordTranslation(word, lang, callback)
		} else {
			callback(dbWord, nil)
		}
	})
}

func (t *TranslationService) apiWordTranslation(word string, language *database.Language, callback func(*database.Word, error)) {
	t.TranslateAPI.GetWordTranslation(word, language.Value, func(rWord *response.TranslateWord, err error) {
		if err != nil {
			logging.Logger.Error(err.Error())
		}
		t.TranslateStorage.SaveWordTranslation(*rWord, *language, func(word *database.Word, errs []error){
			callback(word, err)
			for _, err := range errs {
				logging.Logger.Error(err.Error())
			}
		})
	})
}

func (t *TranslationService) TranslateSentence(sentenceIndex int, chapterIndex int, bookID int, language string, callback func(database.Sentence, error))  {
	t.BookService.GetSentence(bookID, chapterIndex, sentenceIndex, func(sen *database.Sentence, err error) {
		if err != nil {
			logging.Logger.Error(err.Error())
		}
		callback(*sen, err)
	})
}

func (t *TranslationService) apiSentenceTranslation(text string, language *database.Language, callback func(*database.Sentence, error)) {
	t.TranslateAPI.GetTextTranslation(text, language.Value, func(sen *response.TranslateSentence, err error) {
		if err != nil {
			logging.Logger.Error(err.Error())
		}
		t.TranslateStorage.SaveSentenceTranslation(*sen, *language, func(sentence *database.Sentence, errs []error) {
			for _, err := range errs {
				logging.Logger.Error(err.Error())
			}
			callback(sentence, err)
		})
	})
}