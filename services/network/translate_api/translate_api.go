package translate_api

import "github.com/BooksTranslateServer/models/third_api/response"

type TranslateAPI interface {
	GetWordTranslation(word string, lang string, callback func (*response.TranslateWord, error))
	GetTextTranslation(text string, lang string, callback func (*response.TranslateSentence, error))
}