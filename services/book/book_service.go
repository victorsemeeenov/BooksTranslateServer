package book

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/BooksTranslateServer/data"
	"github.com/BooksTranslateServer/models/database"
	"github.com/BooksTranslateServer/services/logging"
	"github.com/BooksTranslateServer/utils/error/types"
	"github.com/cheggaaa/go-poppler"
	"github.com/ledongthuc/pdf"
	"io/ioutil"
	"os"
	"strings"
)

const (
	PDF_FILE = "application/pdf"
	TEXT_FILE = "text/plain"
	REGULAR_FILE = "application/octet-stream"
)

type BookService struct {}

type Chapters struct {
	values []string `json:"chapters"`
}

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
	doc, err := poppler.Open(path)
	if err != nil {
		return err
	}
	pageCount := doc.GetNPages()
	var errs []error
	var orderIndex int
	for index := 0; index < pageCount; index++ {
		page := doc.GetPage(index)
		text := page.Text()
		sentences := strings.Split(text, ".")
		errs = saveSentences(orderIndex, sentences, -1, langID, int(bookID))
		orderIndex = orderIndex + len(sentences)
	}
	if len(errs) > 0 {
		return  errs[0]
	} else {
		return nil
	}
}

func readPlainText(path string, bookID uint, langID int) (error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		logging.Logger.Error(err.Error())
		return err
	}
	allText := string(data)
	if textCmptns := strings.Split(allText, "==="); len(textCmptns) > 1 {
		var result map[string][]interface{}
		jsonObject := textCmptns[0]
		buffer := new(bytes.Buffer)
		if err := json.Compact(buffer, []byte(jsonObject)); err != nil {
			logging.Logger.Error(err.Error())
			return err
		}
		text := textCmptns[1]
		textComponents := strings.Split(text, "---")
		err := json.Unmarshal(buffer.Bytes(), &result)
		if err != nil {
			return err
		}
		chapters := result["chapters"]
		if len(textComponents) < len(chapters) || len(textComponents) > len(chapters) {
			err = errors.New("Invalid text chapters count")
			logging.Logger.Error(err.Error())
			return err
		}
		var errs []error
		for index, chapter := range chapters {
			chapter, _ := saveChapter(chapter.(string), index, fmt.Sprint(index+1), bookID)
			textComponent:= textComponents[index]
			sentences := strings.Split(textComponent, ".")
			errs = saveSentences(index, sentences, int(chapter.ID), langID, int(bookID))
		}
		if len(errs) > 0 {
			err = errs[0]
		}
		return err
	} else {
		err = errors.New("Doesnt have chapters!!!")
		logging.Logger.Error(err.Error())
		return err
	}
	return err
}

func saveSentences(orderIndex int, sentences []string, chapterID int, langID int, bookID int) []error {
	orderNumber := orderIndex
	var errs []error
	if len(sentences) > 0 {
		for _, sen := range sentences {
			if sen != "" {
				var dbSentence database.Sentence
				if chapterID == -1 {
					dbSentence = database.Sentence {
						Value: sen,
						BookID: bookID,
						OrderNumber: orderNumber,
						LanguageID: langID,
					}
				} else {
					dbSentence = database.Sentence {
						Value: sen,
						BookID: bookID,
						OrderNumber: orderNumber,
						ChapterID: chapterID,
						LanguageID: langID,
					}
				}
				err := data.Db.Create(&dbSentence).Error
				if err != nil {
					errs = append(errs, err)
				}
				newErrs := addWordsToSentence(dbSentence)
				for _, err := range newErrs {
					errs = append(errs, err)
				}
				orderNumber ++
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
	_, err := saveChapter(ol.Title, orderIndex, newPrefixOrderValue, bookID)
	errs = append(errs, err)
	for index, outline := range ol.Child {
		newErrs := addChapterToDb(outline, index, newPrefixOrderValue, bookID)
		for _, err := range newErrs {
			errs = append(errs, err)
		}
	}
	return errs
}

func saveChapter(title string, orderIndex int, prefixOrderValue string, bookID uint) (database.Chapter, error) {
	chapter := &database.Chapter{
		Title: title,
		OrderNumber: orderIndex,
		OrderValue: prefixOrderValue,
		BookID: bookID,
	}
	err := data.Db.Create(chapter).Error
	return *chapter, err
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

func (b BookService) GetSentence(sentenceID int) (*database.Sentence, error) {
	var sentence database.Sentence
	err := data.Db.First(&sentence, sentenceID).Error
	return &sentence, err
}

func (b BookService) GetBookList() ([]database.Book, error) {
	var books []database.Book
	err := data.Db.Find(&books).Error
	return books, err
}

func(b BookService) GetAllSentence(bookID int) ([]database.Sentence, error) {
	var sentences []database.Sentence
	err := data.Db.Where("book_id = ?", bookID).Find(&sentences).Error
	return sentences, err
}