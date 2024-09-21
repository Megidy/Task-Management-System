package models

import "time"

type Task struct {
	Id          int        `json:"id"`          //unique id
	Title       string     `json:"title"`       //title
	Description string     `json:"description"` //description
	Priority    string     `json:"priority"`    //low ,middle,high
	Status      string     `json:"status"`      //pending , done , outstanding
	Dependency  []*Task    `json:"dependency"`  // some other tasks
	Created     *time.Time `json:"created"`     //when was created
	To_done     *time.Time `json:"to_done"`     // to submit until this date
}
