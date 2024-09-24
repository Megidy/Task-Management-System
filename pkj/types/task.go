package types

import "time"

type TaskUpdateRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Priority    string    `json:"priority"`
	Status      string    `json:"status"`
	Dependency  int       `json:"dependency"`
	ToDone      time.Time `json:"to_done"` //"2024-09-22T15:00:00Z"
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
