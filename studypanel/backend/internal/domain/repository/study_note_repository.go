package repository

import (
	"context"

	"studypanel-backend/internal/domain/entity"
)

type StudyNoteRepository interface {
	Create(ctx context.Context, note *entity.StudyNote) (*entity.StudyNote, error)
	ListByItemID(ctx context.Context, studyItemID string) ([]*entity.StudyNote, error)
	Update(ctx context.Context, id string, note *entity.StudyNote) (*entity.StudyNote, error)
	Delete(ctx context.Context, id string) error
}
