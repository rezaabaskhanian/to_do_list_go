package repository

import (
	"errors"

	"github.com/rezaabaskhanian/to_do_list_go/internal/domain"
)

type TaskMemoryRepository struct {
	tasks map[int]domain.Task
}

func NewTaskMemoryRepository() *TaskMemoryRepository {
	return &TaskMemoryRepository{
		tasks: make(map[int]domain.Task),
	}
}

// دریافت همه‌ی تسک‌ها
func (r *TaskMemoryRepository) GetAll() ([]domain.Task, error) {
	var tasks []domain.Task
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// دریافت یک تسک بر اساس ID
func (r *TaskMemoryRepository) GetByID(id int) (*domain.Task, error) {
	task, exists := r.tasks[id]
	if !exists {
		return nil, errors.New("task not found")
	}
	return &task, nil
}

// ایجاد یک تسک جدید
func (r *TaskMemoryRepository) Create(task *domain.Task) error {
	r.tasks[task.ID] = *task
	return nil
}

// به‌روزرسانی تسک
func (r *TaskMemoryRepository) Update(task *domain.Task) error {
	_, exists := r.tasks[task.ID]
	if !exists {
		return errors.New("task not found")
	}
	r.tasks[task.ID] = *task
	return nil
}

// حذف تسک
func (r *TaskMemoryRepository) Delete(id int) error {
	_, exists := r.tasks[id]
	if !exists {
		return errors.New("task not found")
	}
	delete(r.tasks, id)
	return nil
}
