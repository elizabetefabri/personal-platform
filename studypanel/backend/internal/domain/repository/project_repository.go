package repository

import (
	"context"

	"studypanel-backend/internal/domain/entity"
)

type ProjectRepository interface {
	Create(ctx context.Context, project *entity.Project) (*entity.Project, error)
	GetByID(ctx context.Context, id string) (*entity.Project, error)
	List(ctx context.Context, projectType string) ([]*entity.Project, error)
	Update(ctx context.Context, id string, project *entity.Project) (*entity.Project, error)
	Delete(ctx context.Context, id string) error
}
