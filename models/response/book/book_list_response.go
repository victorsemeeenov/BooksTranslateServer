package book

import "github.com/BooksTranslateServer/models/database"

type BookResponse struct {
	ID 						int			 `json:"id"`
	Title 				string   `json:"title"`
	NumberOfPages int	     `json:"numberOfPages"`
	Authors 			[]string `json:"author"`
	Year					int   `json:"year"`
	Language      string   `json:"language"`
}

func CreateBookListResponse(books []database.Book) []BookResponse {
	var list []BookResponse
	for _, book := range books {
		var authors []string
		for _, a := range book.Authors {
			authors = append(authors, a.Name)
		}
		response := BookResponse {
			ID:            int(book.ID),
			Title:         book.Name,
			NumberOfPages: book.NumberOfPages,
			Authors:       authors,
			Year:          book.Year,
			Language: 		 book.Language.Value,
		}
		list = append(list, response)
	}
	return list
}