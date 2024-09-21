package config

import (
	"database/sql"
	"log"
	"time"

	"github.com/Megidy/TaskManagmentSystem/pkj/utils"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() {
	dsn, err := utils.CreateDSN()
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}
