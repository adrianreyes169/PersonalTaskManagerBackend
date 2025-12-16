package tasks

import (
	"time"
)

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	createdAt time.Time `json:"createdAt"`
}
