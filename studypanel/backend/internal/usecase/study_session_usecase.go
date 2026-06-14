package usecase

import (
	"context"
	"errors"
	"time"

	"studypanel-backend/internal/domain/entity"
	"studypanel-backend/internal/domain/repository"
)

// --- Create ---

type CreateStudySessionInput struct {
	StartedAt           time.Time `json:"startedAt"`
	EndedAt             time.Time `json:"endedAt"`
	DurationSeconds     int       `json:"durationSeconds"`
	Topic               string    `json:"topic"`
	Notes               string    `json:"notes"`
	Rating              int       `json:"rating"`
	MilestonesCompleted int       `json:"milestonesCompleted"`
}

type CreateStudySessionUseCase struct {
	repo repository.StudySessionRepository
}

func NewCreateStudySessionUseCase(repo repository.StudySessionRepository) *CreateStudySessionUseCase {
	return &CreateStudySessionUseCase{repo: repo}
}

func (uc *CreateStudySessionUseCase) Execute(ctx context.Context, itemID string, input CreateStudySessionInput) (*entity.StudySession, error) {
	if itemID == "" {
		return nil, errors.New("studyItemId é obrigatório")
	}
	if input.DurationSeconds < 0 {
		return nil, errors.New("durationSeconds não pode ser negativo")
	}
	if input.Rating < 1 || input.Rating > 5 {
		return nil, errors.New("rating deve ser entre 1 e 5")
	}

	session := &entity.StudySession{
		StudyItemID:         itemID,
		StartedAt:           input.StartedAt,
		EndedAt:             input.EndedAt,
		DurationSeconds:     input.DurationSeconds,
		Topic:               input.Topic,
		Notes:               input.Notes,
		Rating:              input.Rating,
		MilestonesCompleted: input.MilestonesCompleted,
		CreatedAt:           time.Now().UTC(),
	}

	return uc.repo.Create(ctx, session)
}

// --- List ---

type ListStudySessionsUseCase struct {
	repo repository.StudySessionRepository
}

func NewListStudySessionsUseCase(repo repository.StudySessionRepository) *ListStudySessionsUseCase {
	return &ListStudySessionsUseCase{repo: repo}
}

func (uc *ListStudySessionsUseCase) Execute(ctx context.Context, studyItemID string) ([]*entity.StudySession, error) {
	return uc.repo.ListByItemID(ctx, studyItemID)
}

// --- Update ---

type UpdateStudySessionInput struct {
	StartedAt           time.Time `json:"startedAt"`
	EndedAt             time.Time `json:"endedAt"`
	DurationSeconds     int       `json:"durationSeconds"`
	Topic               string    `json:"topic"`
	Notes               string    `json:"notes"`
	Rating              int       `json:"rating"`
	MilestonesCompleted int       `json:"milestonesCompleted"`
}

type UpdateStudySessionUseCase struct {
	repo repository.StudySessionRepository
}

func NewUpdateStudySessionUseCase(repo repository.StudySessionRepository) *UpdateStudySessionUseCase {
	return &UpdateStudySessionUseCase{repo: repo}
}

func (uc *UpdateStudySessionUseCase) Execute(ctx context.Context, id string, input UpdateStudySessionInput) (*entity.StudySession, error) {
	if id == "" {
		return nil, errors.New("id é obrigatório")
	}
	if input.Rating < 1 || input.Rating > 5 {
		return nil, errors.New("rating deve ser entre 1 e 5")
	}

	session := &entity.StudySession{
		StartedAt:           input.StartedAt,
		EndedAt:             input.EndedAt,
		DurationSeconds:     input.DurationSeconds,
		Topic:               input.Topic,
		Notes:               input.Notes,
		Rating:              input.Rating,
		MilestonesCompleted: input.MilestonesCompleted,
	}

	return uc.repo.Update(ctx, id, session)
}

// --- Delete ---

type DeleteStudySessionUseCase struct {
	repo repository.StudySessionRepository
}

func NewDeleteStudySessionUseCase(repo repository.StudySessionRepository) *DeleteStudySessionUseCase {
	return &DeleteStudySessionUseCase{repo: repo}
}

func (uc *DeleteStudySessionUseCase) Execute(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id é obrigatório")
	}
	return uc.repo.Delete(ctx, id)
}
