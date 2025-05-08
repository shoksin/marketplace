package initializer

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnv() {
	if err := godotenv.Load("./internal/config/.env"); err != nil {
		log.Fatal("Error loading .env file" + err.Error())
	}
}
