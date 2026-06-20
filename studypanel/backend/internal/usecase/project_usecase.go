package usecase

import (
	"context"
	"errors"
	"time"

	"studypanel-backend/internal/domain/entity"
	"studypanel-backend/internal/domain/repository"
)

// ── Create ───────────────────────────────────────────────────────────────────

type CreateProjectInput struct {
	Name        string              `json:"name"`
	Type        entity.ProjectType  `json:"type"`
	Description string              `json:"description"`
	Tags        []string            `json:"tags"`
	RepoURL     string              `json:"repoUrl"`
	DeployURL   string              `json:"deployUrl"`
	Slug        string              `json:"slug"`
	BannerColor string              `json:"bannerColor"`
	ImageURL    string              `json:"imageUrl"`
	ImageAlt    string              `json:"imageAlt"`
	DetailRoute string              `json:"detailRoute"`
	Active      bool                `json:"active"`
	Order       int                 `json:"order"`
}

type CreateProjectUseCase struct{ repo repository.ProjectRepository }

func NewCreateProjectUseCase(r repository.ProjectRepository) *CreateProjectUseCase {
	return &CreateProjectUseCase{repo: r}
}

func (uc *CreateProjectUseCase) Execute(ctx context.Context, in CreateProjectInput) (*entity.Project, error) {
	if in.Name == "" {
		return nil, errors.New("name é obrigatório")
	}
	if !entity.ValidProjectTypes[in.Type] {
		return nil, errors.New("type deve ser 'pessoal' ou 'profissional'")
	}
	if in.Tags == nil {
		in.Tags = []string{}
	}
	p := &entity.Project{
		Name: in.Name, Type: in.Type, Description: in.Description,
		Tags: in.Tags, RepoURL: in.RepoURL, DeployURL: in.DeployURL,
		Slug: in.Slug, BannerColor: in.BannerColor, ImageURL: in.ImageURL,
		ImageAlt: in.ImageAlt, DetailRoute: in.DetailRoute,
		Active: in.Active, Order: in.Order,
		CreatedAt: time.Now().UTC(), UpdatedAt: time.Now().UTC(),
	}
	return uc.repo.Create(ctx, p)
}

// ── List ─────────────────────────────────────────────────────────────────────

type ListProjectsUseCase struct{ repo repository.ProjectRepository }

func NewListProjectsUseCase(r repository.ProjectRepository) *ListProjectsUseCase {
	return &ListProjectsUseCase{repo: r}
}

func (uc *ListProjectsUseCase) Execute(ctx context.Context, projectType string) ([]*entity.Project, error) {
	return uc.repo.List(ctx, projectType)
}

// ── Get ──────────────────────────────────────────────────────────────────────

type GetProjectUseCase struct{ repo repository.ProjectRepository }

func NewGetProjectUseCase(r repository.ProjectRepository) *GetProjectUseCase {
	return &GetProjectUseCase{repo: r}
}

func (uc *GetProjectUseCase) Execute(ctx context.Context, id string) (*entity.Project, error) {
	if id == "" {
		return nil, errors.New("id é obrigatório")
	}
	return uc.repo.GetByID(ctx, id)
}

// ── Update ───────────────────────────────────────────────────────────────────

type UpdateProjectInput struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	RepoURL     string   `json:"repoUrl"`
	DeployURL   string   `json:"deployUrl"`
	BannerColor string   `json:"bannerColor"`
	ImageURL    string   `json:"imageUrl"`
	ImageAlt    string   `json:"imageAlt"`
	DetailRoute string   `json:"detailRoute"`
	Active      bool     `json:"active"`
	Order       int      `json:"order"`
}

type UpdateProjectUseCase struct{ repo repository.ProjectRepository }

func NewUpdateProjectUseCase(r repository.ProjectRepository) *UpdateProjectUseCase {
	return &UpdateProjectUseCase{repo: r}
}

func (uc *UpdateProjectUseCase) Execute(ctx context.Context, id string, in UpdateProjectInput) (*entity.Project, error) {
	if id == "" {
		return nil, errors.New("id é obrigatório")
	}
	if in.Tags == nil {
		in.Tags = []string{}
	}
	p := &entity.Project{
		Name: in.Name, Description: in.Description, Tags: in.Tags,
		RepoURL: in.RepoURL, DeployURL: in.DeployURL,
		BannerColor: in.BannerColor, ImageURL: in.ImageURL, ImageAlt: in.ImageAlt,
		DetailRoute: in.DetailRoute, Active: in.Active, Order: in.Order,
	}
	return uc.repo.Update(ctx, id, p)
}

// ── Delete ───────────────────────────────────────────────────────────────────

type DeleteProjectUseCase struct{ repo repository.ProjectRepository }

func NewDeleteProjectUseCase(r repository.ProjectRepository) *DeleteProjectUseCase {
	return &DeleteProjectUseCase{repo: r}
}

func (uc *DeleteProjectUseCase) Execute(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id é obrigatório")
	}
	return uc.repo.Delete(ctx, id)
}
