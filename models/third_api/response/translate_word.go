package response

import (
	"encoding/json"
	"github.com/buger/jsonparser"
)

type WordMean struct {
	Text 		 string `json:"text"`
	PartOfSpeech string `json:"pos"`
	Gender		 string `json:"gen"`
}

type Translation struct {
	WordMean WordMean
	Synonims   []WordMean
	Means		  		    		  []string
	ExamplesAndTranslations map[string]string
}

type TranslateWordList struct {
	Code uint
	Message string
	Translations []TranslateWord
}

type TranslateWord struct {
	Text 	      			string
	PartOfSpeech  			string
	Transcription			string
	Translations 		[]Translation
}

func TranslateWordFromJSON(jsondata []byte) (res *TranslateWordList, errors []error) {
	res = &TranslateWordList{}
	var translations []TranslateWord
	_, err := jsonparser.ArrayEach(jsondata,
		func(bytes []byte, _ jsonparser.ValueType, _ int, err error) {
			translation, errs := parseDef(bytes)
			for _, err := range errs {
				errors = append(errors, err)
			}
			translations = append(translations, translation)
		},
		"def")
	if err != nil {
		print(err)
	}
	code, _ := jsonparser.GetInt(jsondata, "code")
	res.Code = uint(code)
	var filteredTranslations []TranslateWord
	for _, translation := range translations {
		if translation.Text != "" {
			filteredTranslations = append(filteredTranslations, translation)
		}
	}
	res.Translations = filteredTranslations
	res.Translations = filteredTranslations
	message, _ := jsonparser.GetString(jsondata, "message")
	res.Message = message
	return 
}

func parseDef(def []byte) (translateWord TranslateWord, errs []error) {
	translateWord.Text, _ = jsonparser.GetString(def, "text")
	translateWord.PartOfSpeech, _ = jsonparser.GetString(def, "pos")
	translateWord.Transcription, _ = jsonparser.GetString(def, "ts")
	var translations []Translation
	_, err := jsonparser.ArrayEach(def,
		func(bytes []byte, _ jsonparser.ValueType, _ int, err error) {
			translation, errs := parseTr(bytes)
			for _, err := range errs {
				errs = append(errs, err)
			}
			translations = append(translations, translation)
		},
		"tr")

	if err != nil {
		errs = append(errs, err)
	}
	translateWord.Translations = translations
	return
}

func parseTr(tr []byte) (translation Translation, errors []error) {
	var syns []WordMean
	translation.WordMean.Text, _ = jsonparser.GetString(tr, "text")
	translation.WordMean.PartOfSpeech, _ = jsonparser.GetString(tr, "pos")
	translation.WordMean.Gender, _ = jsonparser.GetString(tr, "gen")
	_, err := jsonparser.ArrayEach(tr,
		func(bytes []byte, _ jsonparser.ValueType, _ int, err error) {
			var syn WordMean
			syn.Text, _ = jsonparser.GetString(bytes, "text")
			syn.PartOfSpeech, _ = jsonparser.GetString(bytes, "pos")
			syn.Gender, _ = jsonparser.GetString(bytes, "gen")
			if err != nil {
				errors = append(errors, err)
			}
			err = json.Unmarshal(bytes, &syn)
			if err == nil {
				syns = append(syns, syn)
			}
		},
		"syn")
	translation.Synonims = syns
	if err != nil {
		errors = append(errors, err)
	}
	err = nil
	var means []string
	_, err = jsonparser.ArrayEach(tr,
		func(bytes []byte, _ jsonparser.ValueType, _ int, err error) {
			mean, err := jsonparser.GetString(bytes)
			if err == nil {
				means = append(means, mean)
			}
		},
		"mean",
	)
	err = nil
	translation.Means = means
	if err != nil {
		errors = append(errors, err)
	}
	err = nil
	examplesAndTranslations := make(map[string]string)
	_, err = jsonparser.ArrayEach(tr,
		func(bytes []byte, _ jsonparser.ValueType, _ int, err error) {
			example, err := jsonparser.GetString(bytes, "text")
			translation, err := jsonparser.GetString(bytes, "tr", "text")
			if err == nil {
				examplesAndTranslations[example] = translation
			}
		},
		"ex",
	)
	translation.ExamplesAndTranslations = examplesAndTranslations
	if err != nil {
		errors = append(errors, err)
	}
	return
}