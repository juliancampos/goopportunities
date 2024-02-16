package config

import (
	"fmt"
	"os"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error
	db, err = InitializeSQLite()
	if err != nil {
		return fmt.Errorf("error initializing sqlite: %v", err)
	}

	err = InitializeEnv()
	if err != nil {
		return fmt.Errorf("error initializing env: %v", err)
	}

	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}

func GetEnvValue(key string) string {
	return os.Getenv(key)
}
