package response

type TranslateSentence struct {
	Code 		 uint	  `json:"code"`
	Lang 		 string	  `json:"lang"`
	Translations []string `json:"text"`
}