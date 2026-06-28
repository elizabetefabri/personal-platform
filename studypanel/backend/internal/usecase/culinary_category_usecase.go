package usecase

import (
	"context"
	"errors"
	"time"

	"studypanel-backend/internal/domain/entity"
	"studypanel-backend/internal/domain/repository"
)

// ── Create ───────────────────────────────────────────────────────────────────

type CreateCulinaryCategoryInput struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Tag         string `json:"tag"`
	Color       string `json:"color"`
	Icon        string `json:"icon"`
	ImageURL    string `json:"imageUrl"`
	Order       int    `json:"order"`
	Active      bool   `json:"active"`
}

type CreateCulinaryCategoryUseCase struct{ repo repository.CulinaryCategoryRepository }

func NewCreateCulinaryCategoryUseCase(r repository.CulinaryCategoryRepository) *CreateCulinaryCategoryUseCase {
	return &CreateCulinaryCategoryUseCase{repo: r}
}

func (uc *CreateCulinaryCategoryUseCase) Execute(ctx context.Context, in CreateCulinaryCategoryInput) (*entity.CulinaryCategory, error) {
	if in.Name == "" {
		return nil, errors.New("name é obrigatório")
	}
	if in.Slug == "" {
		return nil, errors.New("slug é obrigatório")
	}
	cat := &entity.CulinaryCategory{
		Name: in.Name, Slug: in.Slug, Description: in.Description,
		Tag: in.Tag, Color: in.Color, Icon: in.Icon,
		ImageURL: in.ImageURL, Order: in.Order, Active: in.Active,
		CreatedAt: time.Now().UTC(), UpdatedAt: time.Now().UTC(),
	}
	return uc.repo.Create(ctx, cat)
}

// ── List ─────────────────────────────────────────────────────────────────────

type ListCulinaryCategoriesUseCase struct{ repo repository.CulinaryCategoryRepository }

func NewListCulinaryCategoriesUseCase(r repository.CulinaryCategoryRepository) *ListCulinaryCategoriesUseCase {
	return &ListCulinaryCategoriesUseCase{repo: r}
}

func (uc *ListCulinaryCategoriesUseCase) Execute(ctx context.Context) ([]*entity.CulinaryCategory, error) {
	return uc.repo.List(ctx)
}

// ── Get ──────────────────────────────────────────────────────────────────────

type GetCulinaryCategoryUseCase struct{ repo repository.CulinaryCategoryRepository }

func NewGetCulinaryCategoryUseCase(r repository.CulinaryCategoryRepository) *GetCulinaryCategoryUseCase {
	return &GetCulinaryCategoryUseCase{repo: r}
}

func (uc *GetCulinaryCategoryUseCase) Execute(ctx context.Context, id string) (*entity.CulinaryCategory, error) {
	if id == "" {
		return nil, errors.New("id é obrigatório")
	}
	return uc.repo.GetByID(ctx, id)
}

// ── Update ───────────────────────────────────────────────────────────────────

type UpdateCulinaryCategoryInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Tag         string `json:"tag"`
	Color       string `json:"color"`
	Icon        string `json:"icon"`
	ImageURL    string `json:"imageUrl"`
	Order       int    `json:"order"`
	Active      bool   `json:"active"`
}

type UpdateCulinaryCategoryUseCase struct{ repo repository.CulinaryCategoryRepository }

func NewUpdateCulinaryCategoryUseCase(r repository.CulinaryCategoryRepository) *UpdateCulinaryCategoryUseCase {
	return &UpdateCulinaryCategoryUseCase{repo: r}
}

func (uc *UpdateCulinaryCategoryUseCase) Execute(ctx context.Context, id string, in UpdateCulinaryCategoryInput) (*entity.CulinaryCategory, error) {
	if id == "" {
		return nil, errors.New("id é obrigatório")
	}
	cat := &entity.CulinaryCategory{
		Name: in.Name, Description: in.Description,
		Tag: in.Tag, Color: in.Color, Icon: in.Icon,
		ImageURL: in.ImageURL, Order: in.Order, Active: in.Active,
	}
	return uc.repo.Update(ctx, id, cat)
}

// ── Delete ───────────────────────────────────────────────────────────────────

type DeleteCulinaryCategoryUseCase struct{ repo repository.CulinaryCategoryRepository }

func NewDeleteCulinaryCategoryUseCase(r repository.CulinaryCategoryRepository) *DeleteCulinaryCategoryUseCase {
	return &DeleteCulinaryCategoryUseCase{repo: r}
}

func (uc *DeleteCulinaryCategoryUseCase) Execute(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id é obrigatório")
	}
	return uc.repo.Delete(ctx, id)
}
