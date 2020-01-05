package adminpanel

import (
	"github.com/BooksTranslateServer/data"
	"github.com/BooksTranslateServer/models/database"
	"github.com/BooksTranslateServer/services/book"
	"github.com/BooksTranslateServer/services/logging"
	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/qor/qor/resource"
	"mime/multipart"
)

type Book struct {
	database.Book
}

type AdminPanel struct {
	book.Book
}

func (a AdminPanel) Register() *admin.Admin {
	bookService := a.Book
	adm := admin.New(&admin.AdminConfig{SiteName:"Admin", DB:data.Db})
	adm.AddResource(&database.Author{})
	book := adm.AddResource(Book{})
	adm.AddResource(&database.BookAuthor{})
	adm.AddResource(&database.BookCategory{})
	adm.AddResource(&database.Sentence{})
	adm.AddResource(&database.Language{})
	book.Meta(&admin.Meta{Name:"NumberOfPages", Type:"readonly"})
	book.Meta(&admin.Meta{Name:"URL", Type:"hidden"})
	book.Meta(&admin.Meta{Name:"BookCategoryID", Type:"readonly"})
	book.Meta(&admin.Meta{Name:"BookFile",
		Type:"file_picker",
		Valuer: func(interface{}, *qor.Context) interface{} { return "" },
		Setter: func(record interface{}, metaValue *resource.MetaValue, context *qor.Context) {
			valueArray, ok := metaValue.Value.([]*multipart.FileHeader)
			if !ok {
				logging.Logger.Error("Cant get header from loading file")
				return
			}
			if len(valueArray) < 1 {
				logging.Logger.Error("empty header value array")
				return
			}
			value := valueArray[0]
			filename := value.Filename
			header := value.Header
			contentType := header.Get("Content-Type")
			file, err := value.Open()
			if err != nil {
				logging.Logger.Error(err.Error())
				return
			}
			defer file.Close()
			bytes := make([]byte, value.Size)
			_, err = file.Read(bytes)
			if err != nil {
				logging.Logger.Error(err.Error())
				return
			}
			_, url, err := bookService.LoadBook(bytes, contentType, filename)
			book, ok := record.(*Book)
			if !ok {
				logging.Logger.Error("Cant get book record")
			} else {
				book.URL = *url
			}
			if err != nil {
				logging.Logger.Error(err.Error())
			}
		}})
	return adm
}

func (b *Book) AfterCreate(scope *gorm.Scope) (err error) {
	if b.URL != "" {
		book.BookService{}.CreateSentences(int(b.ID), b.URL, b.LanguageID)
	}
	return
}
