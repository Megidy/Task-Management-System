package types

type Log struct {
	UserId int    `json:"user_id"`
	TaskId int    `json:"task_id"`
	Date   string `json:"date"`
	Action string `json:"action"`
}
