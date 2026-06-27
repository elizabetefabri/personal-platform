package repository

import (
	"context"

	"studypanel-backend/internal/domain/entity"
)

type CulinariaRecipeRepository interface {
	Create(ctx context.Context, recipe *entity.CulinariaRecipe) (*entity.CulinariaRecipe, error)
	GetByID(ctx context.Context, id string) (*entity.CulinariaRecipe, error)
	List(ctx context.Context, category string) ([]*entity.CulinariaRecipe, error)
	Update(ctx context.Context, id string, recipe *entity.CulinariaRecipe) (*entity.CulinariaRecipe, error)
	Delete(ctx context.Context, id string) error
}
