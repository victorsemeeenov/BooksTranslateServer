package translate_api

import (
	"encoding/json"
	"github.com/BooksTranslateServer/config"
	"github.com/BooksTranslateServer/models/third_api/response"
	"github.com/BooksTranslateServer/services/network"
	"strings"
)

//MARK: Base urls
const (
	YANDEX_DICT_BASE_URL 	  = "https://dictionary.yandex.net/api/v1/dicservice.json"
	YANDEX_TRANSLATE_BASE_URL = "https://translate.yandex.net/api/v1.5/tr.json"
)

//MARK: Paths
const (
	TRANSLATE_TEXT = YANDEX_TRANSLATE_BASE_URL + "/translate"
	TRANSLATE_WORD = YANDEX_DICT_BASE_URL + "/lookup"
) 

type YandexService struct {
	YandexTranslateAPIKey string
	YandexDictApiKey 	  string
}

func GetYandexService() YandexService {
	return YandexService {
		YandexTranslateAPIKey: config.GetYandexTranslateAPIKey(),
		YandexDictApiKey:	   config.GetYandexDictAPIKey(),
	}
}

func (y YandexService) GetWordTranslation(word string,
										   lang string) (*response.TranslateWordList, error) {
	res, err := network.GetRequest(TRANSLATE_WORD, nil, map[string]string{
		"key": y.YandexDictApiKey,
		"lang": lang + "-" + "ru",
		"text": word,
		"ui": "ru",
		"flags": "2",
	})
	if err != nil {
		return nil, err
	}
	translateWord, errs := response.TranslateWordFromJSON(res.Body())
	if len(errs) > 0 {
		return translateWord, errs[0]
	}
	return translateWord, nil
} 

func (y YandexService) GetTextTranslation(text string,
										   lang string) (*response.TranslateSentence, error) {
	res, err := network.GetRequest(TRANSLATE_TEXT, nil, map[string]string{
		"key": y.YandexTranslateAPIKey,
		"text": strings.ReplaceAll(text, "\n", ""),
		"lang": "ru",
	})
	if err != nil {
		return nil, err
	}
	var translateSentence response.TranslateSentence
	err = json.Unmarshal(res.Body(), &translateSentence)
	if err != nil {
		return nil, err
	}
	return &translateSentence, nil
}

