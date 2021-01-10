package utils

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func HandleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func getEnv(key string) string {
	err := godotenv.Load(".env")
	HandleErr(err)
	return os.Getenv(key)
}

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable",
		getEnv("DB_USER"), getEnv("DB_PASSWORD"), getEnv("DB_NAME"))
	db, err := gorm.Open(getEnv("DB_DRIVER"), dsn)
	HandleErr(err)
	return db
}
