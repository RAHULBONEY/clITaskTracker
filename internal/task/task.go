package task

import (
	"time"
)

type Task struct {
	Name        string
	IsCompleted bool
	Time        time.Time
	ID          int
	CompletedAt time.Time
}
