package types

type Dependency struct {
	UserId          int `json:"user_id"`
	TaskId          int `json:"task_id"`
	DependentTaskId int `json:"dependent_task_id"`
}
