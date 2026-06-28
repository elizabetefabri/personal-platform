package repository

import (
	"context"

	"studypanel-backend/internal/domain/entity"
)

type CulinaryCategoryRepository interface {
	Create(ctx context.Context, cat *entity.CulinaryCategory) (*entity.CulinaryCategory, error)
	GetByID(ctx context.Context, id string) (*entity.CulinaryCategory, error)
	GetBySlug(ctx context.Context, slug string) (*entity.CulinaryCategory, error)
	List(ctx context.Context) ([]*entity.CulinaryCategory, error)
	Update(ctx context.Context, id string, cat *entity.CulinaryCategory) (*entity.CulinaryCategory, error)
	Delete(ctx context.Context, id string) error
}
