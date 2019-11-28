package database_model

import (
	"github.com/jinzhu/gorm"
)

type WordTranslation struct {
	gorm.Model
	WordID	      int `gorm:"name:word_id"`
	TranslationID int `gorm:"name:translation_id"`
}

func (w *WordTranslation) TableName() string {
	return "words_translations"
}