package translate_storage

import (
	"github.com/BooksTranslateServer/models/database"
	"github.com/BooksTranslateServer/data"
	"errors"
	"github.com/BooksTranslateServer/utils/error/types"
	"github.com/BooksTranslateServer/models/third_api/response"
)

type TranslateDB struct {}

func (t TranslateDB) GetWordTranslation(word string, language string, callback func (dbWord *database.Word, dbLanguage *database.Language, errs []error)) {
	go func() { 
		var errs []error
		dbLanguage, err := findLanguage(language)
		if err != nil {
			errs = append(errs, err)
			callback(nil, nil, errs)
			return
		}
		var dbWord *database.Word
		if data.Db.Where("value = ?", word).
		Where("language_id = ?", dbLanguage.ID). 
		First(dbWord).
		RecordNotFound() {
			err := errors.New(types.CANT_FIND_WORD_IN_DB)
			errs = append(errs, err)
			callback(dbWord, dbLanguage, errs)
			return
		}
		if data.Db.Model(dbWord). 
		Related(dbWord.Translations). 
		RecordNotFound() {
			err := errors.New(types.CANT_FIND_WORD_TRANSLATION_IN_DB)
			errs = append(errs, err)
			callback(dbWord, dbLanguage, errs)
			return
		}
		if data.Db.Model(dbWord).
		Related(dbWord.Synonims).
		RecordNotFound() {
			err := errors.New(types.CANT_FIND_SYNONIMS)
			errs = append(errs, err)
		}
		callback(dbWord, dbLanguage, errs)
		return
	}()
}

func (t TranslateDB) GetTextTranslation(text string, language string, callback func (*database.Sentence, *database.Language, []error)) {
	var errs []error
	dbLanguage, err := findLanguage(language)
	if err != nil {
		errs = append(errs, err)
		callback(nil, nil, errs)
		return
	}
	var dbSentence database.Sentence
	if data.Db.Where("value = ?", text). 
	Where("language_id = ?", dbLanguage.ID). 
	First(&dbSentence). 
	RecordNotFound() {
		err := errors.New(types.CANT_FIND_SENTENCE_IN_DB)
		errs = append(errs, err)
		callback(nil, nil, errs)
		return
	}
	if data.Db.Model(dbSentence). 
	Related(&dbSentence). 
	RecordNotFound() {
		err := errors.New(types.CANT_FIND_SENTENCE_TRANSLATION_IN_DB)
		errs = append(errs, err)
		callback(nil, nil, errs)
	}
}

func (t TranslateDB) SaveWordTranslation(res response.TranslateWord, lang database.Language, callback func (*database.Word, []error)) {
	go func(){
		var errs []error
		var dbWord *database.Word
		if data.Db. 
		Where("value = ?", res.Text).
		First(&database.Word{}).
		RecordNotFound() {
			dbWord := &database.Word{
				Value: res.Text,
				Transcription: res.Transcription,
				PartOfSpeech: res.PartOfSpeech,
				LanguageID: int(lang.ID),
			}
			err := data.Db.Create(&dbWord).Error
			if err != nil {
				errs = append(errs, err)
			}
		}
		var dbTranslation database.Translation
		if data.Db.
		Where("value = ?", res.Translation.Text). 
		Find(&dbTranslation).
		RecordNotFound() {
			dbTranslation = database.Translation {Value: res.Translation.Text,
												 PartOfSpeech: res.Translation.PartOfSpeech,
												 Gender: res.Translation.Gender,
												}
			err := data.Db.Create(&dbTranslation).Error
			if err != nil {
				errs = append(errs, err)
			}
		}
		data.Db.Create(&database.WordTranslation{
			WordID: int(dbWord.ID),
			TranslationID: int(dbTranslation.ID),
		})
		for _, syn := range res.Synonims {
			var dbSyn database.Synonim
			if data.Db. 
			Where("value = ?", syn.Text). 
			First(&dbSyn).
			RecordNotFound() {
				dbSyn = database.Synonim {
					WordID: int(dbWord.ID),
					Value: syn.Text,
					PartOfSpeech: syn.PartOfSpeech,
					Gender: syn.Gender,
				}
				data.Db.Create(&dbSyn)
			}
			data.Db.Create(&database.WordSynonim{
				WordID: int(dbWord.ID),
				SynonimID: int(dbSyn.ID),
			})
		}
		var retWord database.Word 
		data.Db.Find(&retWord, dbWord.ID)
		callback(&retWord, errs)
	}()
}

func (t TranslateDB) SaveSentenceTranslation(res response.TranslateSentence, lang database.Language, callback func(*database.Sentence, []error)) {
	
}

func findLanguage(value string) (*database.Language, error) {
	var dbLanguage database.Language
	if data.Db.Where("value = ?", value). 
	First(&dbLanguage). 
	RecordNotFound() {
		err := errors.New(types.CANT_FIND_LANGUAGE_IN_DB)
		return nil, err
	}
	return &dbLanguage, nil
}

