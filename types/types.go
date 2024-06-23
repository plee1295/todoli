package types

import "time"

const ()

type Project struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Tasks       []Task    `json:"tasks"`
	CreatedAt   time.Time `json:"created_at"`
}

type Task struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ProjectID   int       `json:"project_id"`
	Status      Status    `json:"status"`
	Priority    int       `json:"priority"`
	Labels      []string  `json:"labels"`
	CreatedAt   time.Time `json:"created_at"`
	DueAt       time.Time `json:"due_at"`
	CompletedAt time.Time `json:"completed_at"`
}

const (
	Open Status = iota + 1
	InProgress
	Completed
)

type Status int

func (s Status) StatusIndex() int {
	return int(s)
}
