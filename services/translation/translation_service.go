package translation

import (
	"github.com/BooksTranslateServer/models/database"
	"github.com/BooksTranslateServer/services/book"
	"github.com/BooksTranslateServer/services/network/translate_api"
	"github.com/BooksTranslateServer/services/translate_storage"
)

type TranslationService struct {
	TranslateAPI 	 translate_api.TranslateAPI
	TranslateStorage translate_storage.TranslateStorage
	BookService	     book.Book
}

func (t *TranslationService) TranslateWord(word string, language string) ([]database.Word, error) {
	dbWord, err := t.TranslateStorage.GetWordTranslation(word, language)
	if err == nil && dbWord != nil {
		return dbWord, nil
	}
	return t.apiWordTranslation(word, language)
}

func (t *TranslationService) apiWordTranslation(word string, language string) ([]database.Word, error) {
	res, err := t.TranslateAPI.GetWordTranslation(word, language)
	if err != nil {
		return nil, err
	}
	return t.TranslateStorage.SaveWordTranslations(*res, language)
}

func (t *TranslationService) TranslateSentence(sentenceID int) (*database.Sentence, error)  {
	sen, err := t.TranslateStorage.GetSentenceTranslation(sentenceID)
	if err == nil && sen != nil {
		return sen, nil
	}
	if sen == nil {
		return nil, err
	}
	return t.apiSentenceTranslation(sen.Value, sen.Language.Value, sentenceID)
}

func (t *TranslationService) apiSentenceTranslation(text string, language string, sentenceID int) (*database.Sentence, error) {
	sen, err := t.TranslateAPI.GetTextTranslation(text, language)
	if err != nil {
		return nil, err
	}
	return t.TranslateStorage.SaveSentenceTranslation(sentenceID, *sen, language)
}