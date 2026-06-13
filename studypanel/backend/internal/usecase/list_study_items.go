package usecase

import (
	"context"

	"studypanel-backend/internal/domain/entity"
	"studypanel-backend/internal/domain/repository"
)

type ListStudyItemsInput struct {
	Section string
	Topic   string
}

type ListStudyItemsUseCase struct {
	repo repository.StudyItemRepository
}

func NewListStudyItemsUseCase(repo repository.StudyItemRepository) *ListStudyItemsUseCase {
	return &ListStudyItemsUseCase{repo: repo}
}

func (uc *ListStudyItemsUseCase) Execute(ctx context.Context, input ListStudyItemsInput) ([]*entity.StudyItem, error) {
	filter := repository.Filter{
		Section: input.Section,
		Topic:   input.Topic,
	}
	return uc.repo.List(ctx, filter)
}
