package repository

import (
	"context"

	"studypanel-backend/internal/domain/entity"
)

type StudyItemRepository interface {
	Create(ctx context.Context, item *entity.StudyItem) (*entity.StudyItem, error)
	GetByID(ctx context.Context, id string) (*entity.StudyItem, error)
	List(ctx context.Context, filter Filter) ([]*entity.StudyItem, error)
	Update(ctx context.Context, id string, item *entity.StudyItem) (*entity.StudyItem, error)
	Delete(ctx context.Context, id string) error
}

type Filter struct {
	Section string
	Topic   string
}
