package simple_sql

import "time"

type TasksModel struct {
	ID          int
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}
