package models

import (
	"database/sql"

	"github.com/Megidy/TaskManagmentSystem/pkj/types"
)

func CreateLog(log types.Log) error {
	_, err := db.Exec("insert into logs(user_id,task_id,action)values(?,?,?)", log.UserId, log.TaskId, log.Action)
	if err != nil {
		return err
	}
	return nil
}

func GetUsersLogs(userId int) ([]types.Log, error) {
	var logs []types.Log
	query, err := db.Query("select * from logs where user_id =?", userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	for query.Next() {
		var log types.Log
		err := query.Scan(&log.UserId, &log.TaskId, &log.Date, &log.Action)
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}
	return logs, nil
}

func GetAllLogs() ([]types.Log, error) {
	var logs []types.Log
	query, err := db.Query("select * from logs")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	for query.Next() {
		var log types.Log
		err := query.Scan(&log.UserId, &log.TaskId, &log.Date, &log.Action)
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}
	return logs, nil
}
