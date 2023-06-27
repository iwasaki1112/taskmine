package entity

import "time"

type Task struct {
	ID          uint
	Title       string
	Description string
	Status      TaskStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
