package types

import "time"

type Project struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Tasks       []Task    `json:"tasks"`
	CreatedAt   time.Time `json:"created_at"`
}

type Task struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Project     string    `json:"project"`
	Status      string    `json:"status"`
	Priority    int       `json:"priority"`
	Labels      []string  `json:"labels"`
	CreatedAt   time.Time `json:"created_at"`
	DueAt       time.Time `json:"due_at"`
	CompletedAt time.Time `json:"completed_at"`
}
