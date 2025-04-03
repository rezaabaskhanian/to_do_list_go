package usecase

import (
	"github.com/rezaabaskhanian/to_do_list_go/internal/domain"
)

// AssigneeUsecase منطق تجاری برای مسئولین
type AssigneeUsecase struct {
	repo domain.AssigneeRepository
}

// NewAssigneeUsecase ساخت یک AssigneeUsecase جدید
func NewAssigneeUsecase(repo domain.AssigneeRepository) *AssigneeUsecase {
	return &AssigneeUsecase{repo: repo}
}

// CreateAssignee ایجاد یک مسئول جدید
func (u *AssigneeUsecase) CreateAssignee(name, email string) (*domain.Assignee, error) {
	assignee := &domain.Assignee{
		Name:  name,
		Email: email,
	}
	err := u.repo.Create(assignee)
	if err != nil {
		return nil, err
	}
	return assignee, nil
}

// GetAllAssignees دریافت تمامی مسئولین
func (u *AssigneeUsecase) GetAllAssignees() ([]domain.Assignee, error) {
	return u.repo.GetAll()
}

// UpdateAssignee ویرایش اطلاعات یک مسئول
func (u *AssigneeUsecase) UpdateAssignee(assignee *domain.Assignee) error {
	return u.repo.Update(assignee)
}

// DeleteAssignee حذف یک مسئول
func (u *AssigneeUsecase) DeleteAssignee(id int) error {
	return u.repo.Delete(id)
}
