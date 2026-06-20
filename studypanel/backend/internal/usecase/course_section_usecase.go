package usecase

import (
	"context"
	"errors"
	"time"

	"studypanel-backend/internal/domain/entity"
	"studypanel-backend/internal/domain/repository"
)

// ── Create ───────────────────────────────────────────────────────────────────

type CreateCourseSectionInput struct {
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	Description string `json:"description"`
	BannerColor string `json:"bannerColor"`
	IconClass   string `json:"iconClass"`
	ImageURL    string `json:"imageUrl"`
	Active      bool   `json:"active"`
	Order       int    `json:"order"`
}

type CreateCourseSectionUseCase struct{ repo repository.CourseSectionRepository }

func NewCreateCourseSectionUseCase(r repository.CourseSectionRepository) *CreateCourseSectionUseCase {
	return &CreateCourseSectionUseCase{repo: r}
}

func (uc *CreateCourseSectionUseCase) Execute(ctx context.Context, in CreateCourseSectionInput) (*entity.CourseSection, error) {
	if in.Slug == "" {
		return nil, errors.New("slug é obrigatório")
	}
	if in.Name == "" {
		return nil, errors.New("name é obrigatório")
	}
	s := &entity.CourseSection{
		Slug: in.Slug, Name: in.Name, Description: in.Description,
		BannerColor: in.BannerColor, IconClass: in.IconClass,
		ImageURL: in.ImageURL, Active: in.Active, Order: in.Order,
		CreatedAt: time.Now().UTC(), UpdatedAt: time.Now().UTC(),
	}
	return uc.repo.Create(ctx, s)
}

// ── List ─────────────────────────────────────────────────────────────────────

type ListCourseSectionsUseCase struct{ repo repository.CourseSectionRepository }

func NewListCourseSectionsUseCase(r repository.CourseSectionRepository) *ListCourseSectionsUseCase {
	return &ListCourseSectionsUseCase{repo: r}
}

func (uc *ListCourseSectionsUseCase) Execute(ctx context.Context) ([]*entity.CourseSection, error) {
	return uc.repo.List(ctx)
}

// ── Get ──────────────────────────────────────────────────────────────────────

type GetCourseSectionUseCase struct{ repo repository.CourseSectionRepository }

func NewGetCourseSectionUseCase(r repository.CourseSectionRepository) *GetCourseSectionUseCase {
	return &GetCourseSectionUseCase{repo: r}
}

func (uc *GetCourseSectionUseCase) Execute(ctx context.Context, id string) (*entity.CourseSection, error) {
	if id == "" {
		return nil, errors.New("id é obrigatório")
	}
	return uc.repo.GetByID(ctx, id)
}

// ── Update ───────────────────────────────────────────────────────────────────

type UpdateCourseSectionInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	BannerColor string `json:"bannerColor"`
	IconClass   string `json:"iconClass"`
	ImageURL    string `json:"imageUrl"`
	Active      bool   `json:"active"`
	Order       int    `json:"order"`
}

type UpdateCourseSectionUseCase struct{ repo repository.CourseSectionRepository }

func NewUpdateCourseSectionUseCase(r repository.CourseSectionRepository) *UpdateCourseSectionUseCase {
	return &UpdateCourseSectionUseCase{repo: r}
}

func (uc *UpdateCourseSectionUseCase) Execute(ctx context.Context, id string, in UpdateCourseSectionInput) (*entity.CourseSection, error) {
	if id == "" {
		return nil, errors.New("id é obrigatório")
	}
	s := &entity.CourseSection{
		Name: in.Name, Description: in.Description,
		BannerColor: in.BannerColor, IconClass: in.IconClass,
		ImageURL: in.ImageURL, Active: in.Active, Order: in.Order,
	}
	return uc.repo.Update(ctx, id, s)
}

// ── Delete ───────────────────────────────────────────────────────────────────

type DeleteCourseSectionUseCase struct{ repo repository.CourseSectionRepository }

func NewDeleteCourseSectionUseCase(r repository.CourseSectionRepository) *DeleteCourseSectionUseCase {
	return &DeleteCourseSectionUseCase{repo: r}
}

func (uc *DeleteCourseSectionUseCase) Execute(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id é obrigatório")
	}
	return uc.repo.Delete(ctx, id)
}
