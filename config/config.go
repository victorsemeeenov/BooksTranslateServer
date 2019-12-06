package config

import (
	"github.com/joho/godotenv"
	"os"
	"github.com/BooksTranslateServer/services/logging"
	"strconv"
)

const (
	DEBUG = "Debug"
	TEST  = "Test"
	PROD  = "Prod"
)

type DBConfig struct {
	Username string 
	Database string 
	Password string 
	Hostname string 
	Port 	 int	
}

func init() {
    if err := godotenv.Load(); err != nil {
		logging.Logger.Fatal("No .env file found")
    }
}

func (d *DBConfig) GetDebug() {
	d.Username = getEnv("DB_USERNAME")
	d.Database = getEnv("DB_NAME")
	d.Password = getEnv("DB_PASSWORD")
	d.Hostname = getEnv("DB_HOST")
	d.Port 	   = getEnvAsInt("DB_PORT")
}

func (d *DBConfig) GetTest() {
	d.Username = getEnv("DB_USERNAME")
	d.Database = getEnv("TEST_DB_NAME")
	d.Password = getEnv("DB_PASSWORD")
	d.Hostname = getEnv("DB_HOST")
	d.Hostname = getEnv("DB_HOST")
	d.Port 	   = getEnvAsInt("DB_PORT")
}

func GetYandexDictAPIKey() string {
	return os.Getenv("YANDEX_DICT_API_KEY")
}

func GetYandexTranslateAPIKey() string {
	return os.Getenv("YANDEX_TRANSLATE_API_KEY")
}

func getEnv(name string) string {
	return os.Getenv(name)
}

func getEnvAsInt(name string) int {
    valueStr := getEnv(name)
	value, _ := strconv.Atoi(valueStr)
	return value
}
