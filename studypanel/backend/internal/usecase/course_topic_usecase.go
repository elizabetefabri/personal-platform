package usecase

import (
	"context"
	"errors"
	"time"

	"studypanel-backend/internal/domain/entity"
	"studypanel-backend/internal/domain/repository"
)

// ── Create ───────────────────────────────────────────────────────────────────

type CreateCourseTopicInput struct {
	SectionSlug string `json:"sectionSlug"`
	Slug        string `json:"slug"`
	Label       string `json:"label"`
	Description string `json:"description"`
	BannerColor string `json:"bannerColor"`
	IconClass   string `json:"iconClass"`
	Skill       string `json:"skill"`
	ImageURL    string `json:"imageUrl"`
	Active      bool   `json:"active"`
	Order       int    `json:"order"`
}

type CreateCourseTopicUseCase struct{ repo repository.CourseTopicRepository }

func NewCreateCourseTopicUseCase(r repository.CourseTopicRepository) *CreateCourseTopicUseCase {
	return &CreateCourseTopicUseCase{repo: r}
}

func (uc *CreateCourseTopicUseCase) Execute(ctx context.Context, in CreateCourseTopicInput) (*entity.CourseTopic, error) {
	if in.SectionSlug == "" {
		return nil, errors.New("sectionSlug é obrigatório")
	}
	if in.Slug == "" {
		return nil, errors.New("slug é obrigatório")
	}
	if in.Label == "" {
		return nil, errors.New("label é obrigatório")
	}
	t := &entity.CourseTopic{
		SectionSlug: in.SectionSlug, Slug: in.Slug, Label: in.Label,
		Description: in.Description, BannerColor: in.BannerColor,
		IconClass: in.IconClass, Skill: in.Skill, ImageURL: in.ImageURL,
		Active: in.Active, Order: in.Order,
		CreatedAt: time.Now().UTC(), UpdatedAt: time.Now().UTC(),
	}
	return uc.repo.Create(ctx, t)
}

// ── List ─────────────────────────────────────────────────────────────────────

type ListCourseTopicsUseCase struct{ repo repository.CourseTopicRepository }

func NewListCourseTopicsUseCase(r repository.CourseTopicRepository) *ListCourseTopicsUseCase {
	return &ListCourseTopicsUseCase{repo: r}
}

func (uc *ListCourseTopicsUseCase) Execute(ctx context.Context, sectionSlug string) ([]*entity.CourseTopic, error) {
	return uc.repo.List(ctx, sectionSlug)
}

// ── Get ──────────────────────────────────────────────────────────────────────

type GetCourseTopicUseCase struct{ repo repository.CourseTopicRepository }

func NewGetCourseTopicUseCase(r repository.CourseTopicRepository) *GetCourseTopicUseCase {
	return &GetCourseTopicUseCase{repo: r}
}

func (uc *GetCourseTopicUseCase) Execute(ctx context.Context, id string) (*entity.CourseTopic, error) {
	if id == "" {
		return nil, errors.New("id é obrigatório")
	}
	return uc.repo.GetByID(ctx, id)
}

// ── Update ───────────────────────────────────────────────────────────────────

type UpdateCourseTopicInput struct {
	Label       string `json:"label"`
	Description string `json:"description"`
	BannerColor string `json:"bannerColor"`
	IconClass   string `json:"iconClass"`
	Skill       string `json:"skill"`
	ImageURL    string `json:"imageUrl"`
	Active      bool   `json:"active"`
	Order       int    `json:"order"`
}

type UpdateCourseTopicUseCase struct{ repo repository.CourseTopicRepository }

func NewUpdateCourseTopicUseCase(r repository.CourseTopicRepository) *UpdateCourseTopicUseCase {
	return &UpdateCourseTopicUseCase{repo: r}
}

func (uc *UpdateCourseTopicUseCase) Execute(ctx context.Context, id string, in UpdateCourseTopicInput) (*entity.CourseTopic, error) {
	if id == "" {
		return nil, errors.New("id é obrigatório")
	}
	t := &entity.CourseTopic{
		Label: in.Label, Description: in.Description,
		BannerColor: in.BannerColor, IconClass: in.IconClass,
		Skill: in.Skill, ImageURL: in.ImageURL,
		Active: in.Active, Order: in.Order,
	}
	return uc.repo.Update(ctx, id, t)
}

// ── Delete ───────────────────────────────────────────────────────────────────

type DeleteCourseTopicUseCase struct{ repo repository.CourseTopicRepository }

func NewDeleteCourseTopicUseCase(r repository.CourseTopicRepository) *DeleteCourseTopicUseCase {
	return &DeleteCourseTopicUseCase{repo: r}
}

func (uc *DeleteCourseTopicUseCase) Execute(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id é obrigatório")
	}
	return uc.repo.Delete(ctx, id)
}
