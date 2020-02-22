package request

type TranslateWordRequest struct {
	Text string `json: "text"`
	Lang string `json: "lang"`
}
