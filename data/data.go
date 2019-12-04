package data

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"io/ioutil"
	"os"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/BooksTranslateServer/services/logging"
)

var Db *gorm.DB

type DBConfig struct {
	Username string `json:"username"`
	Database string `json:"database"`
	Password string `json:"password"`
	Hostname string `json:"hostname"`
	Port 	 int	`json:"port"`
}

func init() {
	var err error
	defer logging.Logger.Sync()
	jsonFile, err := os.Open("data/database_config.json")
	if err != nil {
		logging.Logger.Fatal(err.Error())
		return
	}
	defer jsonFile.Close()

	byteValue, error := ioutil.ReadAll(jsonFile) 
	if error != nil {
		logging.Logger.Fatal(err.Error())
		return
	}

	var config DBConfig
	json.Unmarshal(byteValue, &config)
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
	Db.AutoMigrate()
}

func ThrowError(db *gorm.DB) error {
	err := db.Error
	logging.Logger.Error(err.Error())
	return err
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