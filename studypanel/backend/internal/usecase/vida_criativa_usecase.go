package usecase

import (
	"context"
	"errors"
	"time"

	"studypanel-backend/internal/domain/entity"
	"studypanel-backend/internal/domain/repository"
)

// ── Create ───────────────────────────────────────────────────────────────────

type CreateVidaCriativaInput struct {
	Title       string                      `json:"title"`
	Description string                      `json:"description"`
	Category    entity.VidaCriativaCategory `json:"category"`
	Status      entity.VidaCriativaStatus   `json:"status"`
	Tags        []string                    `json:"tags"`
	ImageURL    string                      `json:"imageUrl"`
	URL         string                      `json:"url"`
	BannerColor string                      `json:"bannerColor"`
	Active      bool                        `json:"active"`
	Order       int                         `json:"order"`
}

type CreateVidaCriativaUseCase struct{ repo repository.VidaCriativaRepository }

func NewCreateVidaCriativaUseCase(r repository.VidaCriativaRepository) *CreateVidaCriativaUseCase {
	return &CreateVidaCriativaUseCase{repo: r}
}

func (uc *CreateVidaCriativaUseCase) Execute(ctx context.Context, in CreateVidaCriativaInput) (*entity.VidaCriativaItem, error) {
	if in.Title == "" {
		return nil, errors.New("title é obrigatório")
	}
	if in.Tags == nil {
		in.Tags = []string{}
	}
	item := &entity.VidaCriativaItem{
		Title: in.Title, Description: in.Description,
		Category: in.Category, Status: in.Status,
		Tags: in.Tags, ImageURL: in.ImageURL, URL: in.URL,
		BannerColor: in.BannerColor, Active: in.Active, Order: in.Order,
		CreatedAt: time.Now().UTC(), UpdatedAt: time.Now().UTC(),
	}
	return uc.repo.Create(ctx, item)
}

// ── List ─────────────────────────────────────────────────────────────────────

type ListVidaCriativaUseCase struct{ repo repository.VidaCriativaRepository }

func NewListVidaCriativaUseCase(r repository.VidaCriativaRepository) *ListVidaCriativaUseCase {
	return &ListVidaCriativaUseCase{repo: r}
}

func (uc *ListVidaCriativaUseCase) Execute(ctx context.Context, category string) ([]*entity.VidaCriativaItem, error) {
	return uc.repo.List(ctx, category)
}

// ── Get ──────────────────────────────────────────────────────────────────────

type GetVidaCriativaUseCase struct{ repo repository.VidaCriativaRepository }

func NewGetVidaCriativaUseCase(r repository.VidaCriativaRepository) *GetVidaCriativaUseCase {
	return &GetVidaCriativaUseCase{repo: r}
}

func (uc *GetVidaCriativaUseCase) Execute(ctx context.Context, id string) (*entity.VidaCriativaItem, error) {
	if id == "" {
		return nil, errors.New("id é obrigatório")
	}
	return uc.repo.GetByID(ctx, id)
}

// ── Update ───────────────────────────────────────────────────────────────────

type UpdateVidaCriativaInput struct {
	Title       string                      `json:"title"`
	Description string                      `json:"description"`
	Category    entity.VidaCriativaCategory `json:"category"`
	Status      entity.VidaCriativaStatus   `json:"status"`
	Tags        []string                    `json:"tags"`
	ImageURL    string                      `json:"imageUrl"`
	URL         string                      `json:"url"`
	BannerColor string                      `json:"bannerColor"`
	Active      bool                        `json:"active"`
	Order       int                         `json:"order"`
}

type UpdateVidaCriativaUseCase struct{ repo repository.VidaCriativaRepository }

func NewUpdateVidaCriativaUseCase(r repository.VidaCriativaRepository) *UpdateVidaCriativaUseCase {
	return &UpdateVidaCriativaUseCase{repo: r}
}

func (uc *UpdateVidaCriativaUseCase) Execute(ctx context.Context, id string, in UpdateVidaCriativaInput) (*entity.VidaCriativaItem, error) {
	if id == "" {
		return nil, errors.New("id é obrigatório")
	}
	if in.Tags == nil {
		in.Tags = []string{}
	}
	item := &entity.VidaCriativaItem{
		Title: in.Title, Description: in.Description,
		Category: in.Category, Status: in.Status,
		Tags: in.Tags, ImageURL: in.ImageURL, URL: in.URL,
		BannerColor: in.BannerColor, Active: in.Active, Order: in.Order,
	}
	return uc.repo.Update(ctx, id, item)
}

// ── Delete ───────────────────────────────────────────────────────────────────

type DeleteVidaCriativaUseCase struct{ repo repository.VidaCriativaRepository }

func NewDeleteVidaCriativaUseCase(r repository.VidaCriativaRepository) *DeleteVidaCriativaUseCase {
	return &DeleteVidaCriativaUseCase{repo: r}
}

func (uc *DeleteVidaCriativaUseCase) Execute(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id é obrigatório")
	}
	return uc.repo.Delete(ctx, id)
}
