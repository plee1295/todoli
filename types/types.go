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

type Label string

func (l Label) String() string {
	return string(l)
}

type Project struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func (p Project) String() string {
	return string(p.Name)
}

type Task struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ProjectID   string    `json:"project_id"`
	ParentID    string    `json:"parent_id"`
	Status      Status    `json:"status"`
	Priority    Priority  `json:"priority"`
	Labels      []Label   `json:"labels"`
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

const (
	Critical Priority = iota + 1
	High
	Medium
	Low
)

type Priority int
