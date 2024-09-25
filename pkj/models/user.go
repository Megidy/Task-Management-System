package models

import (
	"database/sql"
	"log"

	"github.com/Megidy/TaskManagmentSystem/pkj/config"
	"github.com/Megidy/TaskManagmentSystem/pkj/types"
)

var db *sql.DB

func init() {
	config.Connect()
	db = config.GetDb()
}
func FindUserById(id float64) (*types.User, error) {
	var user types.User
	row := db.QueryRow("select * from users where id =?", id)
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func GetUser(username string) (*types.User, error) {
	var user types.User
	//probably will need to rework , reduce amount of data that will return to user
	row := db.QueryRow("select * from users where username =?", username)
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
func IsSignedUp(username string) (bool, error) {
	var user types.User
	row := db.QueryRow("select username from users where username = ? ", username)
	log.Println("queried")
	err := row.Scan(&user.Username)
	log.Println("scanned")
	//rework
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		} else {
			return true, err
		}
	}
	if username == user.Username {
		log.Println("compared")
		return true, err
	}
	return false, nil
}

func CreateUser(username string, password string) error {
	_, err := db.Exec("insert into users(username,password) values (?,?) ", username, password)
	if err != nil {
		return err
	}
	return nil
}
