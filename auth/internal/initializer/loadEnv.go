package initializer

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnv() {
	if err := godotenv.Load("./internal/config/.env"); err != nil {
		log.Println("Error loading .env file" + err.Error())
		log.Println("Continue work")
	}
}
