package translation

import "github.com/BooksTranslateServer/models/database"

type SentenceTranslationResponse struct {
	Translations []string `json:"translations"`
}

func MakeSentenceTranslationResponse (dbSentences []database.SentenceTranslation) SentenceTranslationResponse {
	var response SentenceTranslationResponse
	for _, sen := range dbSentences {
		response.Translations = append(response.Translations, sen.Value)
	}
	return response
}