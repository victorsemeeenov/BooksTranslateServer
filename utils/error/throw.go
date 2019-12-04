package error

import (
	"github.com/jinzhu/gorm"
	"github.com/BooksTranslateServer/services/logging"
)

func Throw(db *gorm.DB) error {
	err := db.Error
	if err != nil {
		logging.Logger.Error(err.Error())
	}
	return err
}