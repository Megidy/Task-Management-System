package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

func UnHashPassword(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
func LoadEnv() error {
	err := godotenv.Load("F:/sigmaa/TaskManagerAPI/.env")
	if err != nil {
		log.Fatal("failed to load env file ")
		return err
	}
	return nil
}
func CreateDSN() (string, error) {
	err := LoadEnv()
	if err != nil {
		return "", err
	}
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dsn := username + ":" + password + "@/" + dbname
	return dsn, nil
}
