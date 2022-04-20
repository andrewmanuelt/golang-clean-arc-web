package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Env(key string) string {
	err := godotenv.Load("./.env")

	if err != nil {
		panic("Unable to load env file")
	}

	return os.Getenv(key)
}
