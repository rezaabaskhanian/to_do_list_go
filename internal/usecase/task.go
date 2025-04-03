package usecase

import (
	"github.com/rezaabaskhanian/to_do_list_go/internal/domain"
)

// TaskUsecase منطق تجاری برای تسک‌ها

/*یعنی اینجا فقط میخاییم بگیم که چیککار کنیم

یجورایی داریم اینجا دستوراتمون را میسازیم

*/

type TaskUsecase struct {
	repo domain.TaskRepository
}

// NewTaskUsecase ساخت یک TaskUsecase جدید
func NewTaskUsecase(repo domain.TaskRepository) *TaskUsecase {
	return &TaskUsecase{repo: repo}
}

// CreateTask ایجاد یک تسک جدید
func (u *TaskUsecase) CreateTask(title, description string) (*domain.Task, error) {
	task := &domain.Task{
		Title:       title,
		Description: description,
		Done:        false,
	}
	err := u.repo.Create(task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

// GetAllTasks دریافت تمامی تسک‌ها
func (u *TaskUsecase) GetAllTasks() ([]domain.Task, error) {
	return u.repo.GetAll()
}

// UpdateTask ویرایش یک تسک
func (u *TaskUsecase) UpdateTask(task *domain.Task) error {
	return u.repo.Update(task)
}

// DeleteTask حذف یک تسک
func (u *TaskUsecase) DeleteTask(id int) error {
	return u.repo.Delete(id)
}
