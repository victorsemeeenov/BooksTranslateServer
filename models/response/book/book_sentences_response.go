package book

import "github.com/BooksTranslateServer/models/database"

type SentenceResponse struct {
	ID 			  uint   `json:"id"`
	Text      string `json:"text"`
	ChapterID int   `json:"chapter_id"`
	BookID		int   `json:"book_id"`
	Language  string `json:"language"`
}

func MakeAllSentenceResponse(sentences []database.Sentence) []SentenceResponse {
	var list []SentenceResponse
	for _, sentence := range sentences {
		response := SentenceResponse{
			ID:        sentence.ID,
			Text:      sentence.Value,
			ChapterID: sentence.ChapterID,
			BookID:    sentence.BookID,
			Language:  sentence.Language.Value,
		}
		list = append(list, response)
	}
	return list
}
