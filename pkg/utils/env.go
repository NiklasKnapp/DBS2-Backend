package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVar(name string) string {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal("Could not load .env file")
	}

	return os.Getenv(name)
}
