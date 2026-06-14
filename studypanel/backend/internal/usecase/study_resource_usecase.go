package usecase

import (
	"context"
	"errors"
	"time"

	"studypanel-backend/internal/domain/entity"
	"studypanel-backend/internal/domain/repository"
)

// --- Create ---

type CreateStudyResourceInput struct {
	Title       string               `json:"title"`
	URL         string               `json:"url"`
	Type        entity.ResourceType  `json:"type"`
	Description string               `json:"description"`
}

type CreateStudyResourceUseCase struct {
	repo repository.StudyResourceRepository
}

func NewCreateStudyResourceUseCase(repo repository.StudyResourceRepository) *CreateStudyResourceUseCase {
	return &CreateStudyResourceUseCase{repo: repo}
}

func (uc *CreateStudyResourceUseCase) Execute(ctx context.Context, itemID string, input CreateStudyResourceInput) (*entity.StudyResource, error) {
	if itemID == "" {
		return nil, errors.New("studyItemId é obrigatório")
	}
	if input.Title == "" {
		return nil, errors.New("title é obrigatório")
	}
	if input.URL == "" {
		return nil, errors.New("url é obrigatório")
	}
	if !entity.ValidResourceTypes[input.Type] {
		return nil, errors.New("type inválido")
	}

	now := time.Now().UTC()
	resource := &entity.StudyResource{
		StudyItemID: itemID,
		Title:       input.Title,
		URL:         input.URL,
		Type:        input.Type,
		Description: input.Description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	return uc.repo.Create(ctx, resource)
}

// --- List ---

type ListStudyResourcesUseCase struct {
	repo repository.StudyResourceRepository
}

func NewListStudyResourcesUseCase(repo repository.StudyResourceRepository) *ListStudyResourcesUseCase {
	return &ListStudyResourcesUseCase{repo: repo}
}

func (uc *ListStudyResourcesUseCase) Execute(ctx context.Context, studyItemID string) ([]*entity.StudyResource, error) {
	return uc.repo.ListByItemID(ctx, studyItemID)
}

// --- Update ---

type UpdateStudyResourceInput struct {
	Title       string              `json:"title"`
	URL         string              `json:"url"`
	Type        entity.ResourceType `json:"type"`
	Description string              `json:"description"`
}

type UpdateStudyResourceUseCase struct {
	repo repository.StudyResourceRepository
}

func NewUpdateStudyResourceUseCase(repo repository.StudyResourceRepository) *UpdateStudyResourceUseCase {
	return &UpdateStudyResourceUseCase{repo: repo}
}

func (uc *UpdateStudyResourceUseCase) Execute(ctx context.Context, id string, input UpdateStudyResourceInput) (*entity.StudyResource, error) {
	if id == "" {
		return nil, errors.New("id é obrigatório")
	}
	if input.Title == "" {
		return nil, errors.New("title é obrigatório")
	}
	if input.URL == "" {
		return nil, errors.New("url é obrigatório")
	}
	if !entity.ValidResourceTypes[input.Type] {
		return nil, errors.New("type inválido")
	}

	resource := &entity.StudyResource{
		Title:       input.Title,
		URL:         input.URL,
		Type:        input.Type,
		Description: input.Description,
	}

	return uc.repo.Update(ctx, id, resource)
}

// --- Delete ---

type DeleteStudyResourceUseCase struct {
	repo repository.StudyResourceRepository
}

func NewDeleteStudyResourceUseCase(repo repository.StudyResourceRepository) *DeleteStudyResourceUseCase {
	return &DeleteStudyResourceUseCase{repo: repo}
}

func (uc *DeleteStudyResourceUseCase) Execute(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id é obrigatório")
	}
	return uc.repo.Delete(ctx, id)
}
