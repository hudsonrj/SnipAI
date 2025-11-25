package checklist

import "time"

type Checklist struct {
	ID          int       `json:"id"`
	TaskID      *int      `json:"task_id,omitempty"` // nil se for checklist independente
	ProjectID   *int      `json:"project_id,omitempty"` // nil se for checklist independente
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Items       []ChecklistItem `json:"items"`
}

type ChecklistItem struct {
	ID          int       `json:"id"`
	ChecklistID int      `json:"checklist_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	Order       int       `json:"order"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewChecklist(title, description string) *Checklist {
	now := time.Now()
	return &Checklist{
		Title:       title,
		Description: description,
		CreatedAt:   now,
		UpdatedAt:   now,
		Items:       []ChecklistItem{},
	}
}

func NewChecklistItem(checklistID int, title, description string, order int) *ChecklistItem {
	now := time.Now()
	return &ChecklistItem{
		ChecklistID: checklistID,
		Title:       title,
		Description: description,
		Completed:   false,
		Order:       order,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

