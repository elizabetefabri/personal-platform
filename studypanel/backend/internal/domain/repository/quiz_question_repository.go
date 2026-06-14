package repository

import (
	"context"

	"studypanel-backend/internal/domain/entity"
)

type QuizQuestionRepository interface {
	Create(ctx context.Context, question *entity.QuizQuestion) (*entity.QuizQuestion, error)
	ListByTopic(ctx context.Context, section, topic string) ([]*entity.QuizQuestion, error)
	Update(ctx context.Context, id string, question *entity.QuizQuestion) (*entity.QuizQuestion, error)
	Delete(ctx context.Context, id string) error
}
