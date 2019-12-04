package database

import (
	"github.com/jinzhu/gorm"
)

type WordSentence struct {
	gorm.Model
	WordID 	   int `gorm:"name:word_id"`
	SentenceID int `gorm:"name:sentence_id"`
}

func (w *WordSentence) TableName() string  {
	return "words_sentences"
}