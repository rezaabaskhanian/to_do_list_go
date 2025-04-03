package domain

import "time"

// Task مدل یک تسک در برنامه
// هر چیزی که در دنیای واقعی قرار دارد  برای ان مدل مینویسیم
type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	AssigneeID  int       `json:"assignee_id,omitempty"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
}

// TaskRepository رابط (interface) برای تعامل با دیتابیس یا حافظه
type TaskRepository interface {
	GetAll() ([]Task, error)
	GetByID(id int) (*Task, error)
	Create(task *Task) error
	Update(task *Task) error
	Delete(id int) error
}
