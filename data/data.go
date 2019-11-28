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
)

var Db *gorm.DB

type DBConfig struct {
	username string
	database string
	password string
	hostname string
	port 	 int
}

func init() {
	var err error
	jsonFile, err := os.Open("database_config.json")
	if err != nil {
		fmt.Println("Error opening json file")
		return
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadFile("database_config.json")
	if err != nil {
		fmt.Println("Error reading json data:", err)
		return
	}

	var config DBConfig
	json.Unmarshal(jsonData, &config)
	configString := fmt.Sprintf("dbname=%s sslmode=disable password=%s username=%s hostname=%s port=%s", config.database, config.password, config.username, config.hostname, config.port)
	Db, err = gorm.Open("postgres", configString)
	if err != nil {
		log.Fatal(err)
		return
	}
	Db.AutoMigrate()
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