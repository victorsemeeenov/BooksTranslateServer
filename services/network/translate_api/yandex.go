package translate_api

import (
	"github.com/go-resty/resty"
	"github.com/BooksTranslateServer/models/third_api/response"
	"github.com/BooksTranslateServer/services/network"
	"github.com/BooksTranslateServer/models/third_api/request"
	"encoding/json"
	"github.com/BooksTranslateServer/config"
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
										   lang string,
										   callback func (*response.TranslateWord, error)) {
	body := []byte(request.TranslateWord{
		Key: y.YandexTranslateAPIKey,
		Lang: lang,
		Text: word,
		UI: "ru",
		Flags: 2,
	}.
	QueryString())
	network.PostRequest(TRANSLATE_WORD,
						map[string]string{},
						body,
						func (res *resty.Response, err error) {
							if err != nil {
								callback(nil, err)
							}
							var translateWord response.TranslateWord
							err = json.Unmarshal(res.Body(), &translateWord)
							callback(&translateWord, err)
						})
} 

func (y YandexService) GetTextTranslation(text string,
										   lang string,
							               callback func (*response.TranslateSentence, error)) {
	body := []byte(request.TranslateText {
		Key: y.YandexDictApiKey,
		Text: text,
		Lang: lang,
	}. 
	QueryString())
	network.PostRequest(y.YandexTranslateAPIKey,
						map[string]string{},
						body,
						func (res *resty.Response, err error) {
							if err != nil {
								callback (nil, err)
							}
							var translateSentence response.TranslateSentence
							err = json.Unmarshal(res.Body(), &translateSentence)
							callback(&translateSentence, err)
						})
}

