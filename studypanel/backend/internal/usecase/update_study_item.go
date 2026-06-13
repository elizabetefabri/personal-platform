package usecase

import (
	"context"
	"errors"
	"time"

	"studypanel-backend/internal/domain/entity"
	"studypanel-backend/internal/domain/repository"
)

type UpdateStudyItemInput struct {
	CourseName string             `json:"courseName"`
	Status     entity.StudyStatus `json:"status"`
	Date       string             `json:"date"`
	URL        string             `json:"url"`
	ImageURL   string             `json:"imageUrl"`
}

type UpdateStudyItemUseCase struct {
	repo repository.StudyItemRepository
}

func NewUpdateStudyItemUseCase(repo repository.StudyItemRepository) *UpdateStudyItemUseCase {
	return &UpdateStudyItemUseCase{repo: repo}
}

func (uc *UpdateStudyItemUseCase) Execute(ctx context.Context, id string, input UpdateStudyItemInput) (*entity.StudyItem, error) {
	if id == "" {
		return nil, errors.New("id é obrigatório")
	}
	if input.CourseName == "" {
		return nil, errors.New("courseName é obrigatório")
	}
	if !entity.ValidStatuses[input.Status] {
		return nil, errors.New("status inválido")
	}

	existing, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	existing.CourseName = input.CourseName
	existing.Status = input.Status
	existing.Date = input.Date
	existing.URL = input.URL
	existing.ImageURL = input.ImageURL
	existing.UpdatedAt = time.Now().UTC()

	return uc.repo.Update(ctx, id, existing)
}
