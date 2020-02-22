package translation

import "github.com/BooksTranslateServer/models/database"

type Synonim struct {
	Value string `json:"value"`
	PartOfSpeech string `json:"pos"`
	Gender string `json:"gen"`
}

type Translation struct {
	Value string `json:"value"`
	PartOfSpeech string `json:"pos"`
	Gender string `json:"gen"`
	Synonims []Synonim `json:"synonims"`
}

type TranslateWordResponse struct {
	Value         string   `json: value`
	Transcription string `json:"transcription"`
	Language 			string   `json:"lang"`
	Translations []Translation `json:"translations"`
}

func MakeTranslateWordResponse(dbWords []database.Word) []TranslateWordResponse {
	var res []TranslateWordResponse
	for _, dbWord := range dbWords {
		word := createWordResponse(dbWord)
		res = append(res, word)
	}
	return res
}

func createWordResponse(dbWord database.Word) TranslateWordResponse {
	var res TranslateWordResponse
	var translations []Translation
	for _, tr := range dbWord.Translations {
		var syns []Synonim
		for _, s := range tr.Synonims {
			syns = append(syns, Synonim{
				Value:        s.Value,
				PartOfSpeech: s.PartOfSpeech,
				Gender:       s.Gender,
			})
		}
		translation := Translation{
			Value:        tr.Value,
			PartOfSpeech: tr.PartOfSpeech,
			Gender:       tr.Gender,
			Synonims:     syns,
		}
		translations = append(translations, translation)
	}
	res.Value = dbWord.Value
	res.Translations = translations
	res.Language = dbWord.Language.Value
	res.Transcription = dbWord.Transcription
	return res
}
