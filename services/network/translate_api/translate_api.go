package translate_api

import "github.com/BooksTranslateServer/models/third_api/response"

type TranslateAPI interface {
	GetWordTranslation(word string, lang string) (*response.TranslateWordList, error)
	GetTextTranslation(text string, lang string) (*response.TranslateSentence, error)
}