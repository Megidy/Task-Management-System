package models

import (
	"database/sql"
	"log"

	"github.com/Megidy/TaskManagmentSystem/pkj/types"
)

func CreateTask(task types.Task) error {
	log.Println(task.Id)
	_, err := db.Exec("insert into tasks(title,description,priority,to_done,user_id) values(?,?,?,?,?)",
		task.Title, task.Description, task.Priority, task.ToDone.Format("2006-01-02 15:04:05"), task.UserId)
	if err != nil {
		return err
	}
	var dependency types.Dependency
	row := db.QueryRow("select id from tasks where user_id=? order by id desc limit 1 ", task.UserId)
	err = row.Scan(&dependency.TaskId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}
	log.Println(task.Id)
	dependency.UserId = task.UserId

	dependency.DependentTaskId = task.Dependency
	log.Println(dependency)
	if task.Dependency != 0 {
		err := AddDependency(dependency)
		if err != nil {
			return err

		}
	}
	return nil

}

func GetAllTasks(userId int) ([]types.Response, error) {
	var response []types.Response
	rows, err := db.Query("select title,description,priority,status,created,to_done from tasks where user_id=?",
		userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}
	for rows.Next() {
		var res types.Response
		err := rows.Scan(&res.Title, &res.Description, &res.Priority, &res.Status, &res.Created, &res.ToDone)
		if err != nil {
			return nil, err
		}
		response = append(response, res)
	}
	return response, nil

}
func GetSingleTask(userId, taskId int) (types.Response, error) {
	var response types.Response
	row := db.QueryRow("select title,description,priority,status,created,to_done from tasks where user_id=? and id=?",
		userId, taskId)

	err := row.Scan(&response.Title, &response.Description, &response.Priority, &response.Status, &response.Created, &response.ToDone)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.Response{}, nil
		} else {
			return types.Response{}, err
		}
	}
	return response, nil
}

func DeleteTask(userId, taskId int) error {
	_, err := db.Exec("delete from tasks where user_id =? and id=?",
		userId, taskId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}
	return nil
}

func UpdateTask(task types.TaskUpdateRequest, userId, taskId int) error {
	log.Println("about to update ")
	log.Println("task in function :", task)
	_, err := db.Exec("update tasks set title=?,description=?,priority=?,status=?,to_done=? where id=? and user_id=?",
		task.Title, task.Description, task.Priority, task.Status, task.ToDone.Format("2006-01-02 15:04:05"), taskId, userId)

	if err != nil {
		return err
	}
	log.Println("updated ")
	return nil
}
