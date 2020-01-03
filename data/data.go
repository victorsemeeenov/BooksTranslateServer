package data

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	cfg "github.com/BooksTranslateServer/config"
	"github.com/BooksTranslateServer/models/database"
	"github.com/BooksTranslateServer/services/logging"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/qor/qor/resource"
	"log"
)

var Db *gorm.DB

func init() {
	InitWithConfig(cfg.DEBUG)
}

func InitWithConfig(c string) {
	var err error
	defer logging.Logger.Sync()
	var config cfg.DBConfig
	switch c {
	case cfg.DEBUG:
		config.GetDebug()
		break
	case cfg.TEST:
		config.GetTest()
		break
	}
	if c == cfg.DEBUG {
		config.GetDebug()
	}
	var password string
	if config.Password != "" {
		password = "password=" + config.Password
	}
	configString := fmt.Sprintf("host=%s port=%d user=%s %s dbname=%s sslmode=disable", config.Hostname, config.Port, config.Username, password, config.Database)
	info := fmt.Sprintf("Database connection:%s", configString)
	logging.Logger.Info(info)
	Db, err = gorm.Open("postgres", configString)
	if err != nil {
		logging.Logger.Fatal(err.Error())
		return
	}
	Db.AutoMigrate(
		&database.AccessToken{},
		&database.Author{},
		&database.Book{},
		&database.BookAuthor{},
		&database.BookCategory{},
		&database.Chapter{},
		&database.Language{},
		&database.RefreshToken{},
		&database.Sentence{},
		&database.SentenceTranslation{},
		&database.Synonim{},
		&database.Translation{},
		&database.User{},
		&database.Word{},
		&database.WordSentence{},
		&database.WordSynonim{},
		&database.WordTranslation{},
	)
}

func RemoveAll() {
	Db.Delete(&database.AccessToken{})
	Db.Delete(&database.Author{})
	Db.Delete(&database.Book{})
	Db.Delete(&database.BookAuthor{})
	Db.Delete(&database.BookCategory{})
	Db.Delete(&database.Chapter{})
	Db.Delete(&database.Language{})
	Db.Delete(&database.RefreshToken{})
	Db.Delete(&database.Sentence{})
	Db.Delete(&database.SentenceTranslation{})
	Db.Delete(&database.Synonim{})
	Db.Delete(&database.Translation{})
	Db.Delete(&database.User{})
	Db.Delete(&database.Word{})
	Db.Delete(&database.WordSentence{})
	Db.Delete(&database.WordSynonim{})
	Db.Delete(&database.WordTranslation{})
}

func ThrowError(db *gorm.DB) error {
	err := db.Error
	logging.Logger.Error(err.Error())
	return err
}

func RegisterAdmin() *admin.Admin {
	db := fmt.Sprintf("%v", Db)
	logging.Logger.Debug(db)
	adm := admin.New(&admin.AdminConfig{SiteName:"Admin", DB:Db})
	adm.AddResource(&database.Author{})
	book := adm.AddResource(&database.Book{})
	adm.AddResource(&database.BookAuthor{})
	adm.AddResource(&database.BookCategory{})
	book.Meta(&admin.Meta{Name:"NumberOfPages", Type:"readonly"})
	book.Meta(&admin.Meta{Name:"URL", Type:"hidden"})
	book.Meta(&admin.Meta{Name:"BookCategoryID", Type:"readonly"})
	book.Meta(&admin.Meta{Name:"BookFile",
		                  Type:"file_picker",
		                  Valuer: func(interface{}, *qor.Context) interface{} { return "" },
		                  Setter: func(record interface{}, metaValue *resource.MetaValue, context *qor.Context) {
							_, _, err := context.Request.FormFile("BookFile")
							if err != nil {
								logging.Logger.Error("Cant load file!!!")
								return
							}
							logging.Logger.Info("File loaded!!!!")
						  }})
	return adm
}

// create a random UUID with from RFC 4122
// adapted from http://github.com/nu7hatch/gouuid
func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

// hash plaintext with SHA-1
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}
