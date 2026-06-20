package repository

import (
	"context"

	"studypanel-backend/internal/domain/entity"
)

type CourseTopicRepository interface {
	Create(ctx context.Context, topic *entity.CourseTopic) (*entity.CourseTopic, error)
	GetByID(ctx context.Context, id string) (*entity.CourseTopic, error)
	List(ctx context.Context, sectionSlug string) ([]*entity.CourseTopic, error)
	Update(ctx context.Context, id string, topic *entity.CourseTopic) (*entity.CourseTopic, error)
	Delete(ctx context.Context, id string) error
}
