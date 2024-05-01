package models

import "time"

type task struct {
	ID          uint      `json:"id"`
	Content     string    `json:"content"`
	DueDate     time.Time `json:"due_date"`
	IsCompleted bool      `json:"is_completed"`
}
