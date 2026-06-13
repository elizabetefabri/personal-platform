package usecase

import (
	"context"
	"errors"

	"studypanel-backend/internal/domain/entity"
	"studypanel-backend/internal/domain/repository"
)

type GetStudyItemUseCase struct {
	repo repository.StudyItemRepository
}

func NewGetStudyItemUseCase(repo repository.StudyItemRepository) *GetStudyItemUseCase {
	return &GetStudyItemUseCase{repo: repo}
}

func (uc *GetStudyItemUseCase) Execute(ctx context.Context, id string) (*entity.StudyItem, error) {
	if id == "" {
		return nil, errors.New("id é obrigatório")
	}
	return uc.repo.GetByID(ctx, id)
}
