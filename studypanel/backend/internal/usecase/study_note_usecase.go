package usecase

import (
	"context"
	"errors"
	"time"

	"studypanel-backend/internal/domain/entity"
	"studypanel-backend/internal/domain/repository"
)

// --- Create ---

type CreateStudyNoteInput struct {
	StudyItemID string             `json:"studyItemId"`
	Date        string             `json:"date"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Progress    int                `json:"progress"`
	Status      entity.NoteStatus  `json:"status"`
	Link        string             `json:"link"`
}

type CreateStudyNoteUseCase struct {
	repo repository.StudyNoteRepository
}

func NewCreateStudyNoteUseCase(repo repository.StudyNoteRepository) *CreateStudyNoteUseCase {
	return &CreateStudyNoteUseCase{repo: repo}
}

func (uc *CreateStudyNoteUseCase) Execute(ctx context.Context, itemID string, input CreateStudyNoteInput) (*entity.StudyNote, error) {
	if itemID == "" {
		return nil, errors.New("studyItemId é obrigatório")
	}
	if input.Title == "" {
		return nil, errors.New("title é obrigatório")
	}
	if !entity.ValidNoteStatuses[input.Status] {
		return nil, errors.New("status inválido")
	}

	now := time.Now().UTC()
	note := &entity.StudyNote{
		StudyItemID: itemID,
		Date:        input.Date,
		Title:       input.Title,
		Description: input.Description,
		Progress:    input.Progress,
		Status:      input.Status,
		Link:        input.Link,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	return uc.repo.Create(ctx, note)
}

// --- List ---

type ListStudyNotesUseCase struct {
	repo repository.StudyNoteRepository
}

func NewListStudyNotesUseCase(repo repository.StudyNoteRepository) *ListStudyNotesUseCase {
	return &ListStudyNotesUseCase{repo: repo}
}

func (uc *ListStudyNotesUseCase) Execute(ctx context.Context, studyItemID string) ([]*entity.StudyNote, error) {
	return uc.repo.ListByItemID(ctx, studyItemID)
}

// --- Update ---

type UpdateStudyNoteInput struct {
	Date        string            `json:"date"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Progress    int               `json:"progress"`
	Status      entity.NoteStatus `json:"status"`
	Link        string            `json:"link"`
}

type UpdateStudyNoteUseCase struct {
	repo repository.StudyNoteRepository
}

func NewUpdateStudyNoteUseCase(repo repository.StudyNoteRepository) *UpdateStudyNoteUseCase {
	return &UpdateStudyNoteUseCase{repo: repo}
}

func (uc *UpdateStudyNoteUseCase) Execute(ctx context.Context, id string, input UpdateStudyNoteInput) (*entity.StudyNote, error) {
	if id == "" {
		return nil, errors.New("id é obrigatório")
	}
	if input.Title == "" {
		return nil, errors.New("title é obrigatório")
	}
	if !entity.ValidNoteStatuses[input.Status] {
		return nil, errors.New("status inválido")
	}

	note := &entity.StudyNote{
		Date:        input.Date,
		Title:       input.Title,
		Description: input.Description,
		Progress:    input.Progress,
		Status:      input.Status,
		Link:        input.Link,
	}

	return uc.repo.Update(ctx, id, note)
}

// --- Delete ---

type DeleteStudyNoteUseCase struct {
	repo repository.StudyNoteRepository
}

func NewDeleteStudyNoteUseCase(repo repository.StudyNoteRepository) *DeleteStudyNoteUseCase {
	return &DeleteStudyNoteUseCase{repo: repo}
}

func (uc *DeleteStudyNoteUseCase) Execute(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id é obrigatório")
	}
	return uc.repo.Delete(ctx, id)
}
