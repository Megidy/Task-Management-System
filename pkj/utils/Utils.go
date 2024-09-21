package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	err := godotenv.Load("F:/sigmaa/Real-time Notification and Task Management System/.env")
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
	username := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
	dsn := username + ":" + password + "@/" + dbname
	return dsn, nil
}
