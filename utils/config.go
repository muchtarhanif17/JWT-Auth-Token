package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetJWTSecret() []byte {

	// Load .env (untuk local)
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found, using system env")
	}

	secret := os.Getenv("JWT_SECRET_KEY")
	if secret == "" {
		log.Fatal("JWT_SECRET_KEY is empty")
	}

	log.Println("JWT SECRET loaded") // debug
	return []byte(secret)
}
