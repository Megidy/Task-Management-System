package models

import (
	"database/sql"
	"log"

	"github.com/Megidy/TaskManagmentSystem/pkj/config"
	"github.com/Megidy/TaskManagmentSystem/pkj/types"
	"github.com/Megidy/TaskManagmentSystem/pkj/utils"
)

var db *sql.DB

func init() {
	config.Connect()
	db = config.GetDb()
	err := utils.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}
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

func IsSignedUpById(userId int) (bool, error) {
	var user types.User
	row := db.QueryRow("select id from users where id = ? ", userId)

	err := row.Scan(&user.Id)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		} else {
			return true, err
		}
	}
	if userId == user.Id {
		return true, err
	}
	return false, nil
}

func IsSignedUp(username string) (bool, error) {
	var user types.User
	row := db.QueryRow("select username from users where username = ? ", username)

	err := row.Scan(&user.Username)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		} else {
			return true, err
		}
	}
	if username == user.Username {
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

func DeleteUser(userId int) error {
	_, err := db.Exec("delete from users where id=?", userId)
	if err != nil {
		return err
	}
	return nil
}
