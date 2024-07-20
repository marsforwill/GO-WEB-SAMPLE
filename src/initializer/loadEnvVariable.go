package initializer

import (
	"log"

	"github.com/joho/godotenv"
)

// Load will read your env file(s) and load them into ENV for later usage
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading .env file %v", err)
	}
}
