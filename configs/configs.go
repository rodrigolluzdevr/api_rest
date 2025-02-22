package configs

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("configs/.env")
	if err != nil {
		log.Println("archive .env not found.")
	}
}
