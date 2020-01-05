package book

import (
	"errors"
	"github.com/BooksTranslateServer/data"
	"github.com/BooksTranslateServer/models/database"
	"github.com/BooksTranslateServer/services/logging"
	"github.com/BooksTranslateServer/utils/error/types"
	"github.com/ledongthuc/pdf"
	"os"
	"strings"
)

const (
	PDF_FILE = "application/pdf"
	TEXT_FILE = "text/plain"
	REGULAR_FILE = "application/octet-stream"
)

type BookService struct {}

func (b BookService) LoadBook(bytes []byte, extension string, name string) (*os.File, *string, error) {
	switch extension {
	case PDF_FILE:
		return createFile(bytes, name, "pdf")
	case TEXT_FILE:
		return createFile(bytes, name, "txt")
	case REGULAR_FILE:
		return createFile(bytes, name, "")
	default:
		return nil, nil, errors.New(types.CANT_LOAD_BOOK_EXTENSION)
	}
}

func (b BookService) CreateSentences(bookID int, fileURL string, languageID int) (error) {
	if strings.HasSuffix(fileURL, "pdf") {
		return readPdf(fileURL, uint(bookID), languageID)
	} else {
		return readPlainText(fileURL, uint(bookID), languageID)
	}
}

func readPdf(path string, bookID uint, langID int) (error) {
	_, reader, err := pdf.Open(path)
	if err != nil {
		return err
	}
	numPage := reader.NumPage()
	outline := reader.Outline()
	for index, ol := range outline.Child {
		addChapterToDb(ol, index, "", bookID)
	}
	for index := 0; index < numPage; index++ {
		page := reader.Page(index)
		if page.V.IsNull() {
			continue
		}
		//MARK: TODO
		errs := savePdfSentences(page, 0, langID)
		for _, err := range errs {
			logging.Logger.Error(err.Error())
		}
	} 
	return nil
}

func readPlainText(path string, bookID uint, langID int) (error) {
	return errors.New(types.CANT_LOAD_BOOK_EXTENSION)
}

func savePdfSentences(page pdf.Page, chapterID int, langID int) []error {
	var errs []error
	textArray := page.Content().Text
	var stringArray []string 
	for _, text := range textArray {
		stringArray = append(stringArray, text.S)
	}
	for _, str := range stringArray {
		components := strings.Split(str, ".")
		if len(components) > 0 {
			for index, sen := range components {
				if sen != "" {
					dbSentence := database.Sentence {
						Value: sen,
						OrderNumber: index,
						ChapterID: chapterID,
						LanguageID: langID,
					}
					err := data.Db.Create(&dbSentence).Error
					if err != nil {
						errs = append(errs, err)
					}
					newErrs := addWordsToSentence(dbSentence)
					for _, err := range newErrs {
						errs = append(errs, err)
					}
				}
			}
		}
	}
	return errs
}

func addWordsToSentence(sentence database.Sentence) []error {
	var errs []error
	senWords := strings.Split(sentence.Value, " ")
	for sen := range senWords {
		var word database.Word
		first := data.Db.Where("value = ?", sen). 
		First(&word)
		errs = append(errs, first.Error)
		if !first.RecordNotFound() {
			data.Db.Create(&database.WordSentence{
				WordID: int(word.ID),
				SentenceID: int(sentence.ID),
			})
		}
	}
	return errs
}

func addChapterToDb(ol pdf.Outline, orderIndex int, prefixOrderValue string, bookID uint) []error {
	var errs []error
	newPrefixOrderValue := prefixOrderValue + string(orderIndex+1) + "."
	err := data.Db.Create(&database.Chapter{
		Title: ol.Title,
		OrderNumber: orderIndex, 
		OrderValue: newPrefixOrderValue,
		BookID: bookID,
	}).Error
	errs = append(errs, err)
	for index, outline := range ol.Child {
		newErrs := addChapterToDb(outline, index, newPrefixOrderValue, bookID)
		for _, err := range newErrs {
			errs = append(errs, err)
		}
	}
	return errs
}

func createFile(bytes []byte, name string, extension string) (*os.File, *string, error) {
	if _, err := os.Open(name); err == nil {
		return nil, nil, errors.New(types.CANT_CREATE_BOOK_WITH_THIS_NAME)
	}
	file, err := os.Create(name)
    if err != nil {
        return nil, nil, err
    }
    defer file.Close()
	file.Write(bytes)
	return file, &name, nil
}

func (b BookService) GetSentence(bookID int, chapterIndex int, sentenceIndex int, callback func(*database.Sentence, error))  {
	
}