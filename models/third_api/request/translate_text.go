package request

import "fmt"

type TranslateText struct {
	Key  string `json:"key"`
	Text string `json:"text"`
	Lang string `json:"lang"`
}

func (t TranslateText) QueryString() string {
	return fmt.Sprintf("key=%s&text=%s&lang=%s", t.Key, t.Text, t.Lang)
}