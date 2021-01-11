package utils

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func GetEnv(key string) string {
	err := godotenv.Load(".env")
	HandleErr(err)
	return os.Getenv(key)
}

func PrepareLog() {
	file, err := os.OpenFile(GetEnv("LOG_FILE"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	HandleErr(err)
	log.SetOutput(file)
}

func HandleErr(err error) {
	if err != nil {
		log.Error(err.Error())
		panic(err.Error())
	}
}

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable",
		GetEnv("DB_USER"), GetEnv("DB_PASSWORD"), GetEnv("DB_NAME"))
	db, err := gorm.Open(GetEnv("DB_DRIVER"), dsn)
	HandleErr(err)
	return db
}

func HashPassword(pass string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	HandleErr(err)
	return string(hashed)
}
