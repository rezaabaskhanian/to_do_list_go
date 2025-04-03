package repository

import (
	"errors"

	"github.com/rezaabaskhanian/to_do_list_go/internal/domain"
)

type AssigneeMemoryRepository struct {
	assignees map[int]domain.Assignee
}

func NewAssigneeMemoryRepository() *AssigneeMemoryRepository {
	return &AssigneeMemoryRepository{
		assignees: make(map[int]domain.Assignee),
	}
}

func (r *AssigneeMemoryRepository) GetAll() ([]domain.Assignee, error) {
	var assignees []domain.Assignee
	for _, assignee := range r.assignees {
		assignees = append(assignees, assignee)
	}
	return assignees, nil
}

func (r *AssigneeMemoryRepository) GetByID(id int) (*domain.Assignee, error) {
	assignee, exists := r.assignees[id]
	if !exists {
		return nil, errors.New("assignee not found")
	}
	return &assignee, nil
}

func (r *AssigneeMemoryRepository) Create(assignee *domain.Assignee) error {
	r.assignees[assignee.ID] = *assignee
	return nil
}

func (r *AssigneeMemoryRepository) Update(assignee *domain.Assignee) error {
	_, exists := r.assignees[assignee.ID]
	if !exists {
		return errors.New("assignee not found")
	}
	r.assignees[assignee.ID] = *assignee
	return nil
}

func (r *AssigneeMemoryRepository) Delete(id int) error {
	_, exists := r.assignees[id]
	if !exists {
		return errors.New("assignee not found")
	}
	delete(r.assignees, id)
	return nil
}
