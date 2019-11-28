package database_model

import (
	"github.com/jinzhu/gorm"
)

type SentenceTranslation struct {
	gorm.Model
	SentenceID int      `gorm:"name:sentence_id"`
	Value 	   string
	Sentence   Sentence `gorm:"association_foreignkey:sentence_id;"`
}

func (s *SentenceTranslation) TableName() string  {
	return "sentence_translations"
}