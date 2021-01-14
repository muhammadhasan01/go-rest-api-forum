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

// GetEnv is a function to get value on the .env flie
func GetEnv(key string) string {
	err := godotenv.Load(".env")
	HandleErr(err)
	return os.Getenv(key)
}

// PrepareLog is a method to start the log and outputs it to the LOG_FILE in the .env file
func PrepareLog() {
	file, err := os.OpenFile(GetEnv("LOG_FILE"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	HandleErr(err)
	log.SetOutput(file)
}

// HandlErr is used to log the error
func HandleErr(err error) {
	if err != nil {
		log.Error(err.Error())
	}
}

// ConnectDB is used to connect to the configured database
func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable",
		GetEnv("DB_USER"), GetEnv("DB_PASSWORD"), GetEnv("DB_NAME"))
	db, err := gorm.Open(GetEnv("DB_DRIVER"), dsn)
	HandleErr(err)
	return db
}

// HashPassword is a function used to hash a password using bcrypt
func HashPassword(pass string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	HandleErr(err)
	return string(hashed)
}
