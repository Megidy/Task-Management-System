package models

import (
	"database/sql"

	"github.com/Megidy/TaskManagmentSystem/pkj/types"
)

func AddDependency(dependency types.Dependency) error {
	_, err := db.Exec("insert into task_dependencies(user_id,task_id,dependent_task_id) values(?,?,?)", dependency.UserId, dependency.TaskId, dependency.DependentTaskId)
	if err != nil {
		return err
	}
	return nil
}

func GetAllDependencies(userId int) ([]types.Dependency, error) {
	var deps []types.Dependency
	query, err := db.Query("select * from task_dependencies where user_id=?", userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	for query.Next() {
		var dep types.Dependency
		err := query.Scan(&dep.UserId, &dep.TaskId, &dep.DependentTaskId)
		if err != nil {
			return nil, err
		}
		deps = append(deps, dep)
	}
	return deps, nil
}
