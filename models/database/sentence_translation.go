package database

import (
	"github.com/jinzhu/gorm"
)

type SentenceTranslation struct {
	gorm.Model
	SentenceID int      `gorm:"name:sentence_id"`
	Value 	   string
	Sentence   Sentence `gorm:"associated_foreignkey:SentenceID;"`
}

func (s *SentenceTranslation) TableName() string  {
	return "sentence_translations"
}