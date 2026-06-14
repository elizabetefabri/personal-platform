package repository

import (
	"context"

	"studypanel-backend/internal/domain/entity"
)

type StudySessionRepository interface {
	Create(ctx context.Context, session *entity.StudySession) (*entity.StudySession, error)
	ListByItemID(ctx context.Context, studyItemID string) ([]*entity.StudySession, error)
	Update(ctx context.Context, id string, session *entity.StudySession) (*entity.StudySession, error)
	Delete(ctx context.Context, id string) error
}
