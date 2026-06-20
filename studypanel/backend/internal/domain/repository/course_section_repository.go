package repository

import (
	"context"

	"studypanel-backend/internal/domain/entity"
)

type CourseSectionRepository interface {
	Create(ctx context.Context, section *entity.CourseSection) (*entity.CourseSection, error)
	GetByID(ctx context.Context, id string) (*entity.CourseSection, error)
	GetBySlug(ctx context.Context, slug string) (*entity.CourseSection, error)
	List(ctx context.Context) ([]*entity.CourseSection, error)
	Update(ctx context.Context, id string, section *entity.CourseSection) (*entity.CourseSection, error)
	Delete(ctx context.Context, id string) error
}
