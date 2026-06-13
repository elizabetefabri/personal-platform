package usecase

import (
	"context"
	"errors"

	"studypanel-backend/internal/domain/repository"
)

type DeleteStudyItemUseCase struct {
	repo repository.StudyItemRepository
}

func NewDeleteStudyItemUseCase(repo repository.StudyItemRepository) *DeleteStudyItemUseCase {
	return &DeleteStudyItemUseCase{repo: repo}
}

func (uc *DeleteStudyItemUseCase) Execute(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id é obrigatório")
	}
	_, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	return uc.repo.Delete(ctx, id)
}
