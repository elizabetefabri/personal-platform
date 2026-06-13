package usecase

import (
	"context"
	"errors"
	"time"

	"studypanel-backend/internal/domain/entity"
	"studypanel-backend/internal/domain/repository"
)

type CreateStudyItemInput struct {
	Section    string             `json:"section"`
	Topic      string             `json:"topic"`
	CourseName string             `json:"courseName"`
	Status     entity.StudyStatus `json:"status"`
	Date       string             `json:"date"`
	URL        string             `json:"url"`
	ImageURL   string             `json:"imageUrl"`
}

type CreateStudyItemUseCase struct {
	repo repository.StudyItemRepository
}

func NewCreateStudyItemUseCase(repo repository.StudyItemRepository) *CreateStudyItemUseCase {
	return &CreateStudyItemUseCase{repo: repo}
}

func (uc *CreateStudyItemUseCase) Execute(ctx context.Context, input CreateStudyItemInput) (*entity.StudyItem, error) {
	if input.Section == "" {
		return nil, errors.New("section é obrigatório")
	}
	if input.Topic == "" {
		return nil, errors.New("topic é obrigatório")
	}
	if input.CourseName == "" {
		return nil, errors.New("courseName é obrigatório")
	}
	if !entity.ValidStatuses[input.Status] {
		return nil, errors.New("status inválido")
	}

	now := time.Now().UTC()
	item := &entity.StudyItem{
		Section:     input.Section,
		Topic:       input.Topic,
		CourseName:  input.CourseName,
		Status:      input.Status,
		Date:        input.Date,
		URL:         input.URL,
		ImageURL:    input.ImageURL,
		DetailRoute: entity.BuildDetailRoute(input.Section, input.Topic),
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	return uc.repo.Create(ctx, item)
}
