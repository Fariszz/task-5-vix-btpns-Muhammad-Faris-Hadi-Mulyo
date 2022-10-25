package helpers

import (
	"log"
	"os"
	"path"

	"github.com/joho/godotenv"
)

func GoDotEnv(key string) string {
	// get env location
	envFile := path.Join(".env")
	err := godotenv.Load(envFile)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}