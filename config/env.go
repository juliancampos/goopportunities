package config

import (
	"os"

	"github.com/joho/godotenv"
)

func InitializeEnv() error {
	err := godotenv.Load("./.env")
	if err != nil {
		panic(err)
	}
	return nil
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
