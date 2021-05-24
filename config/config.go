package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DBHost = ""
var DBPort = ""
var DBDatabase = ""

func Load() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")
	DBDatabase = os.Getenv("DB_DATABASE")
}
