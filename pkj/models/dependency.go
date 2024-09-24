package models

import "database/sql"

type Dependency struct {
	UserId          int `json:"user_id"`
	TaskId          int `json:"task_id"`
	DependentTaskId int `json:"dependent_task_id"`
}

func AddDependency(dependency Dependency) error {
	_, err := db.Exec("insert into task_dependencies(user_id,task_id,dependent_task_id) values(?,?,?)", dependency.UserId, dependency.TaskId, dependency.DependentTaskId)
	if err != nil {
		return err
	}
	return nil
}

func GetAllDependencies(userId int) ([]Dependency, error) {
	var deps []Dependency
	query, err := db.Query("select * from task_dependencies where user_id=?", userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	for query.Next() {
		var dep Dependency
		err := query.Scan(&dep.UserId, &dep.TaskId, &dep.DependentTaskId)
		if err != nil {
			return nil, err
		}
		deps = append(deps, dep)
	}
	return deps, nil
}
