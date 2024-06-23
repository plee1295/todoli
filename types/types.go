package types

import (
	"time"

	"github.com/spf13/cobra"
)

type Commands struct {
	Add    *cobra.Command
	Delete *cobra.Command
	List   *cobra.Command
}

type Project struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
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
