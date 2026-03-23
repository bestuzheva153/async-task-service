package model

import "time"

type TaskStatus string

const (
	StatusPending    TaskStatus = "pending"
	StatusProcessing TaskStatus = "processing"
	StatusDone       TaskStatus = "done"
	StatusFailed     TaskStatus = "failed"
)

type Task struct {
	ID        int64      `json:"id"`
	Type      string     `json:"type"`
	Payload   string     `json:"payload"`
	Status    TaskStatus `json:"status"`
	Result    *string    `json:"result,omitempty"`
	Error     *string    `json:"error,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
