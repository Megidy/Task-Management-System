package types

import "time"

type Task struct {
	Id          int       `json:"id"`          //unique id
	Title       string    `json:"title"`       //title
	Description string    `json:"description"` //description
	Priority    string    `json:"priority"`    //low ,middle,high
	Status      string    `json:"status"`      //pending , done , outstanding
	Dependency  int       `json:"dependency"`  // some other tasks
	Created     time.Time `json:"created"`     //when was created
	ToDone      time.Time `json:"to_done"`     // to submit until this date
	UserId      int       `json:"user_id"`     //id of the user that created that task
}

type TaskUpdateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    string `json:"priority"`

	Dependency int       `json:"dependency"`
	ToDone     time.Time `json:"to_done"` //"2024-09-22T15:00:00Z"
}

type Response struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
	Status      string `json:"status"`
	Created     string `json:"created"`
	ToDone      string `json:"to_done"`
}
type TaskRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Priority    string    `json:"priority"`
	Dependency  int       `json:"dependency"`
	ToDone      time.Time `json:"to_done"` //"2024-09-22T15:00:00Z"
}

type ChangeStatus struct {
	TaskId int    `json:"task_id"`
	Status string `json:"status"`
	UserId int    `json:"userId"`
}
