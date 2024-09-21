package config

import (
	"database/sql"
	"log"
	"time"

	"github.com/Megidy/TaskManagmentSystem/pkj/utils"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connect() {
	dsn, err := utils.CreateDSN()
	if err != nil {
		log.Fatal(err)
	}
	database, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db.Ping()
	db = database
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}
func GetDb() *sql.DB {
	return db
}
