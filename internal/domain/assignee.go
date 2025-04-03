package domain

// Assignee تعریف مدل مسئول یک تسک
// هر چیزی که در دنیای واقعی قرار دارد  برای ان مدل مینویسیم

type Assignee struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// AssigneeRepository اینترفیس برای عملیات روی مسئولین
// کسانی که مسوول اجرای تسک هستند به انها مسوول میگویم
type AssigneeRepository interface {
	GetAll() ([]Assignee, error)
	GetByID(id int) (*Assignee, error)
	Create(assignee *Assignee) error
	Update(assignee *Assignee) error
	Delete(id int) error
}
