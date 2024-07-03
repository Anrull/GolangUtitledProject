package env

import (
	"github.com/joho/godotenv"

	"log"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func GetValue(column string) string {
	value, ok := os.LookupEnv(column)
	if ok {
		return value
	}
	log.Fatal("No value for " + column)
	return ""
}
