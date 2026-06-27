package usecase

import (
	"context"
	"errors"
	"time"

	"studypanel-backend/internal/domain/entity"
	"studypanel-backend/internal/domain/repository"
)

// ── Create ───────────────────────────────────────────────────────────────────

type CreateCulinariaRecipeInput struct {
	Title        string                   `json:"title"`
	Description  string                   `json:"description"`
	Category     string                   `json:"category"`
	Ingredients  []string                 `json:"ingredients"`
	Instructions []string                 `json:"instructions"`
	PrepTime     int                      `json:"prepTime"`
	Difficulty   entity.RecipeDifficulty  `json:"difficulty"`
	Tags         []string                 `json:"tags"`
	ImageURL     string                   `json:"imageUrl"`
	Servings     int                      `json:"servings"`
	Active       bool                     `json:"active"`
}

type CreateCulinariaRecipeUseCase struct{ repo repository.CulinariaRecipeRepository }

func NewCreateCulinariaRecipeUseCase(r repository.CulinariaRecipeRepository) *CreateCulinariaRecipeUseCase {
	return &CreateCulinariaRecipeUseCase{repo: r}
}

func (uc *CreateCulinariaRecipeUseCase) Execute(ctx context.Context, in CreateCulinariaRecipeInput) (*entity.CulinariaRecipe, error) {
	if in.Title == "" {
		return nil, errors.New("title é obrigatório")
	}
	if in.Tags == nil {
		in.Tags = []string{}
	}
	if in.Ingredients == nil {
		in.Ingredients = []string{}
	}
	if in.Instructions == nil {
		in.Instructions = []string{}
	}
	recipe := &entity.CulinariaRecipe{
		Title: in.Title, Description: in.Description,
		Category: in.Category, Ingredients: in.Ingredients,
		Instructions: in.Instructions, PrepTime: in.PrepTime,
		Difficulty: in.Difficulty, Tags: in.Tags,
		ImageURL: in.ImageURL, Servings: in.Servings, Active: in.Active,
		CreatedAt: time.Now().UTC(), UpdatedAt: time.Now().UTC(),
	}
	return uc.repo.Create(ctx, recipe)
}

// ── List ─────────────────────────────────────────────────────────────────────

type ListCulinariaRecipesUseCase struct{ repo repository.CulinariaRecipeRepository }

func NewListCulinariaRecipesUseCase(r repository.CulinariaRecipeRepository) *ListCulinariaRecipesUseCase {
	return &ListCulinariaRecipesUseCase{repo: r}
}

func (uc *ListCulinariaRecipesUseCase) Execute(ctx context.Context, category string) ([]*entity.CulinariaRecipe, error) {
	return uc.repo.List(ctx, category)
}

// ── Get ──────────────────────────────────────────────────────────────────────

type GetCulinariaRecipeUseCase struct{ repo repository.CulinariaRecipeRepository }

func NewGetCulinariaRecipeUseCase(r repository.CulinariaRecipeRepository) *GetCulinariaRecipeUseCase {
	return &GetCulinariaRecipeUseCase{repo: r}
}

func (uc *GetCulinariaRecipeUseCase) Execute(ctx context.Context, id string) (*entity.CulinariaRecipe, error) {
	if id == "" {
		return nil, errors.New("id é obrigatório")
	}
	return uc.repo.GetByID(ctx, id)
}

// ── Update ───────────────────────────────────────────────────────────────────

type UpdateCulinariaRecipeInput struct {
	Title        string                   `json:"title"`
	Description  string                   `json:"description"`
	Category     string                   `json:"category"`
	Ingredients  []string                 `json:"ingredients"`
	Instructions []string                 `json:"instructions"`
	PrepTime     int                      `json:"prepTime"`
	Difficulty   entity.RecipeDifficulty  `json:"difficulty"`
	Tags         []string                 `json:"tags"`
	ImageURL     string                   `json:"imageUrl"`
	Servings     int                      `json:"servings"`
	Active       bool                     `json:"active"`
}

type UpdateCulinariaRecipeUseCase struct{ repo repository.CulinariaRecipeRepository }

func NewUpdateCulinariaRecipeUseCase(r repository.CulinariaRecipeRepository) *UpdateCulinariaRecipeUseCase {
	return &UpdateCulinariaRecipeUseCase{repo: r}
}

func (uc *UpdateCulinariaRecipeUseCase) Execute(ctx context.Context, id string, in UpdateCulinariaRecipeInput) (*entity.CulinariaRecipe, error) {
	if id == "" {
		return nil, errors.New("id é obrigatório")
	}
	if in.Tags == nil {
		in.Tags = []string{}
	}
	if in.Ingredients == nil {
		in.Ingredients = []string{}
	}
	if in.Instructions == nil {
		in.Instructions = []string{}
	}
	recipe := &entity.CulinariaRecipe{
		Title: in.Title, Description: in.Description,
		Category: in.Category, Ingredients: in.Ingredients,
		Instructions: in.Instructions, PrepTime: in.PrepTime,
		Difficulty: in.Difficulty, Tags: in.Tags,
		ImageURL: in.ImageURL, Servings: in.Servings, Active: in.Active,
	}
	return uc.repo.Update(ctx, id, recipe)
}

// ── Delete ───────────────────────────────────────────────────────────────────

type DeleteCulinariaRecipeUseCase struct{ repo repository.CulinariaRecipeRepository }

func NewDeleteCulinariaRecipeUseCase(r repository.CulinariaRecipeRepository) *DeleteCulinariaRecipeUseCase {
	return &DeleteCulinariaRecipeUseCase{repo: r}
}

func (uc *DeleteCulinariaRecipeUseCase) Execute(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id é obrigatório")
	}
	return uc.repo.Delete(ctx, id)
}
