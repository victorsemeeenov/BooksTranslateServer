package request

import (
	"fmt"
)

type TranslateWord struct {
	Key   string 
	Lang  string 
	Text  string 
	UI    string 
	Flags int
}

func (t TranslateWord) QueryString() string  {
	return fmt.Sprintf("key=%s&lang=%s&text=%s&ui=%s&flags=%d",
					   t.Key,
					   t.Lang,
					   t.Text,
				       t.UI,
					   t.Flags)
}