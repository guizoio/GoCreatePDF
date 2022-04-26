package configs

import (
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv(filename ...string) {
	if filename == nil {
		filename = []string{"./.env"}
	}
	if os.Getenv("APP") == "" {
		err := godotenv.Load(filename[0])
		if err != nil {
			panic("error loading .env file")
		}
	}
}
