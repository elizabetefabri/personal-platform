package repository

import (
	"context"

	"studypanel-backend/internal/domain/entity"
)

type VidaCriativaRepository interface {
	Create(ctx context.Context, item *entity.VidaCriativaItem) (*entity.VidaCriativaItem, error)
	GetByID(ctx context.Context, id string) (*entity.VidaCriativaItem, error)
	List(ctx context.Context, category string) ([]*entity.VidaCriativaItem, error)
	Update(ctx context.Context, id string, item *entity.VidaCriativaItem) (*entity.VidaCriativaItem, error)
	Delete(ctx context.Context, id string) error
}
