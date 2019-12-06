package response

import (
	"github.com/buger/jsonparser"
	"encoding/json"
)

type WordMean struct {
	Text 		 string `json:"text"`
	PartOfSpeech string `json:"pos"`
	Gender		 string `json:"gen"`
}

type TranslateWord struct {
	Text 	      			string
	PartOfSpeech  			string
	Transcription 			string
	Translation   			WordMean
	Synonims	  		    []WordMean
	Means		  		    []string
	ExamplesAndTranslations map[string]string
	Code					uint
	Message					string
}

func TranslateWordFromJSON(jsondata []byte) (traslateWord TranslateWord, errors []error) {
	def, _, _, err := jsonparser.Get(jsondata, "def")
	if err != nil {
		errors = append(errors, err)
	}
	err = nil
	tr, _, _, err := jsonparser.Get(def, "tr")
	if err != nil {
		errors = append(errors, err) 
	}
	err = nil
	text, err := jsonparser.GetString(def, "text")
	traslateWord.Text = text
	if err != nil {
		errors = append(errors, err)
	}
	err = nil
	var translation WordMean
	err = json.Unmarshal(tr, &translation)
	traslateWord.Translation = translation
	if err != nil {
		errors = append(errors, err)
	}
	err = nil
	partOfSpeech, err := jsonparser.GetString(def, "pos")
	traslateWord.PartOfSpeech = partOfSpeech
	if err != nil {
		errors = append(errors, err)
	}
	err = nil
	var syns []WordMean
	transcription, err := jsonparser.GetString(def, "ts")
	traslateWord.Transcription = transcription
	_, err = jsonparser.ArrayEach(tr,
		 func(bytes []byte, _ jsonparser.ValueType, _ int, err error) {
			var syn WordMean
			err = json.Unmarshal(bytes, &syn)
			if err == nil {
				syns = append(syns, syn)
			}
	},
	"syn")
	traslateWord.Synonims = syns
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
	traslateWord.Means = means
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
	traslateWord.ExamplesAndTranslations = examplesAndTranslations
	if err != nil {
		errors = append(errors, err)
	}
	err = nil
	code, err := jsonparser.GetInt(jsondata, "code")
	traslateWord.Code = uint(code)
	if err != nil {
		errors = append(errors, err)
	}
	err = nil
	message, err := jsonparser.GetString(jsondata, "message")
	traslateWord.Message = message
	return
}