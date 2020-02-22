package translate_storage

import (
	"errors"
	"github.com/BooksTranslateServer/data"
	"github.com/BooksTranslateServer/models/database"
	"github.com/BooksTranslateServer/models/third_api/response"
	"github.com/BooksTranslateServer/services/logging"
	"github.com/BooksTranslateServer/utils/error/types"
)

type TranslateDB struct {}

func (t TranslateDB) GetWordTranslation(word string, language string) ([]database.Word, error) {
		dbLanguage, err := findLanguage(language)
		if err != nil {
			return nil, err
		}
		var dbWords []database.Word
		if data.Db.Where("value = ? and language_id = ?", word, dbLanguage.ID).
		Where("language_id = ?", dbLanguage.ID). 
		Find(&dbWords).
		RecordNotFound() {
			err = errors.New(types.CANT_FIND_WORD_IN_DB)
			return nil, err
		}
		var words []database.Word
		for _, dbWord := range dbWords {
			if data.Db.Model(&dbWord).Related(&dbWord.Translations, "Translations").RecordNotFound() {
				err = errors.New(types.CANT_FIND_WORD_TRANSLATION_IN_DB)
				return nil, err
			}
			var trs []database.Translation
			for _, translation := range dbWord.Translations {
				data.Db.Model(&translation).Related(&translation.Synonims, "Synonims")
				tr := translation
				trs = append(trs, tr)
			}
			dbWord.Translations = trs
			if data.Db.Model(&dbWord).Related(&dbWord.Language).RecordNotFound() {
				err = errors.New(types.CANT_FIND_LANGUAGE_IN_DB)
				return nil, err
			}
			if data.Db.Model(&dbWord).Related(&dbWord.Sentences, "Sentences").RecordNotFound() {
				err = errors.New(types.CANT_FIND_SENTENCE_IN_DB)
				return nil, err
			}
			words = append(words, dbWord)
		}
		return words, nil
}

func (t TranslateDB) GetSentenceTranslation(sentenceID int) (*database.Sentence, error) {
	var dbSentence database.Sentence
	data.Db.First(&dbSentence, sentenceID).Related(&dbSentence.Chapter)
	if data.Db.Model(&dbSentence).Related(&dbSentence.Language).
	RecordNotFound() {
		err := errors.New(types.CANT_FIND_LANGUAGE_IN_DB)
		return nil, err
	}
	data.Db.Model(&dbSentence).Related(&dbSentence.Translations)
	if len(dbSentence.Translations) == 0{
			err := errors.New(types.CANT_FIND_SENTENCE_TRANSLATION_IN_DB)
			return &dbSentence, err
	}
	if data.Db.Model(&dbSentence).Related(&dbSentence.Words).
		RecordNotFound() {
			err := errors.New(types.CANT_FIND_WORD_IN_DB)
			logging.Logger.Error(err.Error())
	}
	return &dbSentence, nil
}

func (t TranslateDB) SaveWordTranslations(res response.TranslateWordList, language string) ([]database.Word, error) {
		lang, err := findLanguage(language)
		if err != nil {
			return nil, errors.New(types.CANT_FIND_LANGUAGE_IN_DB)
		}
	 	return createTranslations(res.Translations, *lang)
}

func createTranslations(resList []response.TranslateWord,
												language database.Language) ([]database.Word, error) {
	var words []database.Word
	for _, res := range resList {
		var dbWord database.Word
		data.Db.
			Where("value = ?", res.Text).
			FirstOrCreate(&dbWord,
				database.Word{
					Value:        res.Text,
					LanguageID:   int(language.ID),
					PartOfSpeech: res.PartOfSpeech,
					Transcription: res.Transcription,
				})
		data.Db.Model(&dbWord).Related(&dbWord.Language)
		data.Db.Model(&dbWord).Related(&dbWord.Translations, "Translations")
		var translations []database.Translation
		for _, translation := range res.Translations {
			var dbTranslation database.Translation
			data.Db.
				Where("value = ?", translation.WordMean.Text).
				FirstOrCreate(&dbTranslation,
					database.Translation{
						Value:        translation.WordMean.Text,
						PartOfSpeech: translation.WordMean.PartOfSpeech,
						Gender:       translation.WordMean.Gender,
					})
			var synonims []database.Synonim
			for _, syn := range translation.Synonims {
				var dbSyn database.Synonim
				data.Db.
					Where("value = ?", syn.Text).
					FirstOrCreate(&dbSyn,
						database.Synonim{
							Value:         syn.Text,
							PartOfSpeech:  syn.PartOfSpeech,
							Gender:        syn.Gender,
							TranslationID: int(dbTranslation.ID),
						})
				data.Db.Where("value = ?", syn.Text).First(&dbSyn)
				synonims = append(synonims, dbSyn)
			}
			data.Db.Model(&dbTranslation).Related(&dbTranslation.Synonims, "Synonims")
			dbTranslation.Synonims = synonims
			data.Db.Save(&dbTranslation)
			var wordTr database.WordTranslation
			data.Db.Where("word_id = ? AND translation_id = ?",
				dbWord.ID,
				dbTranslation.ID).
				FirstOrCreate(&wordTr, database.WordTranslation{
					WordID:        int(dbWord.ID),
					TranslationID: int(dbTranslation.ID),
				})
			translations = append(translations, dbTranslation)
		}
		dbWord.Translations = translations
		data.Db.Model(&dbWord).Save(&dbWord)
		words = append(words, dbWord)
	}
	return words, nil
}

func (t TranslateDB) SaveSentenceTranslation(sentenceID int, res response.TranslateSentence, language string) (*database.Sentence, error) {
	lang, err := findLanguage(language)
	if err != nil {
		return nil, err
	}
	var dbSentence database.Sentence
	err = data.Db.Find(&dbSentence, sentenceID).Error
	if err != nil {
		return nil, err
	}
	var translations []database.SentenceTranslation
	for _, translation := range res.Translations {
		var dbTranslation database.SentenceTranslation
		err = data.Db.Where("value = ?", translation).
			FirstOrCreate(&dbTranslation, database.SentenceTranslation{
				Value: translation,
				SentenceID: sentenceID,
			}).Error
		if err != nil {
			return nil, err
		}
		translations = append(translations, dbTranslation)
	}
	dbSentence.Translations = translations
	dbSentence.Language = *lang
	data.Db.Save(&dbSentence)
	return &dbSentence, nil
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

