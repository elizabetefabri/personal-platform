package repository

import (
	"context"

	"studypanel-backend/internal/domain/entity"
)

type StudyResourceRepository interface {
	Create(ctx context.Context, resource *entity.StudyResource) (*entity.StudyResource, error)
	ListByItemID(ctx context.Context, studyItemID string) ([]*entity.StudyResource, error)
	Update(ctx context.Context, id string, resource *entity.StudyResource) (*entity.StudyResource, error)
	Delete(ctx context.Context, id string) error
}
