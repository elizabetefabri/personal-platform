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
	CategoryID          string                  `json:"categoryId"`
	CategorySlug        string                  `json:"categorySlug"`
	Name                string                  `json:"name"`
	Slug                string                  `json:"slug"`
	Description         string                  `json:"description"`
	PrepTimeMinutes     int                     `json:"prepTimeMinutes"`
	CookTimeMinutes     int                     `json:"cookTimeMinutes"`
	ServingsStr         string                  `json:"servingsStr"`
	Difficulty          entity.RecipeDifficulty `json:"difficulty"`
	Status              entity.RecipeStatus     `json:"status"`
	Tags                []string                `json:"tags"`
	ImageURL            string                  `json:"imageUrl"`
	YoutubeURL          string                  `json:"youtubeUrl"`
	SourceURL           string                  `json:"sourceUrl"`
	Ingredients         []string                `json:"ingredients"`
	PreparationSteps    []string                `json:"preparationSteps"`
	Utensils            string                  `json:"utensils"`
	Tips                string                  `json:"tips"`
	Substitutions       string                  `json:"substitutions"`
	StorageInstructions string                  `json:"storageInstructions"`
	EstimatedCost       float64                 `json:"estimatedCost"`
	PersonalRating      int                     `json:"personalRating"`
	Tested              bool                    `json:"tested"`
	Notes               string                  `json:"notes"`
	Active              bool                    `json:"active"`
}

type CreateCulinariaRecipeUseCase struct{ repo repository.CulinariaRecipeRepository }

func NewCreateCulinariaRecipeUseCase(r repository.CulinariaRecipeRepository) *CreateCulinariaRecipeUseCase {
	return &CreateCulinariaRecipeUseCase{repo: r}
}

func (uc *CreateCulinariaRecipeUseCase) Execute(ctx context.Context, in CreateCulinariaRecipeInput) (*entity.CulinariaRecipe, error) {
	if in.Name == "" {
		return nil, errors.New("name é obrigatório")
	}
	if in.Slug == "" {
		return nil, errors.New("slug é obrigatório")
	}
	if in.Tags == nil {
		in.Tags = []string{}
	}
	if in.Ingredients == nil {
		in.Ingredients = []string{}
	}
	if in.PreparationSteps == nil {
		in.PreparationSteps = []string{}
	}
	if in.Status == "" {
		in.Status = entity.RecipeStatusPending
	}
	recipe := &entity.CulinariaRecipe{
		CategoryID:          in.CategoryID,
		CategorySlug:        in.CategorySlug,
		Name:                in.Name,
		Slug:                in.Slug,
		Description:         in.Description,
		PrepTimeMinutes:     in.PrepTimeMinutes,
		CookTimeMinutes:     in.CookTimeMinutes,
		ServingsStr:         in.ServingsStr,
		Difficulty:          in.Difficulty,
		Status:              in.Status,
		Tags:                in.Tags,
		ImageURL:            in.ImageURL,
		YoutubeURL:          in.YoutubeURL,
		SourceURL:           in.SourceURL,
		Ingredients:         in.Ingredients,
		PreparationSteps:    in.PreparationSteps,
		Utensils:            in.Utensils,
		Tips:                in.Tips,
		Substitutions:       in.Substitutions,
		StorageInstructions: in.StorageInstructions,
		EstimatedCost:       in.EstimatedCost,
		PersonalRating:      in.PersonalRating,
		Tested:              in.Tested,
		Notes:               in.Notes,
		Active:              in.Active,
		CreatedAt:           time.Now().UTC(),
		UpdatedAt:           time.Now().UTC(),
	}
	return uc.repo.Create(ctx, recipe)
}

// ── List ─────────────────────────────────────────────────────────────────────

type ListCulinariaRecipesUseCase struct{ repo repository.CulinariaRecipeRepository }

func NewListCulinariaRecipesUseCase(r repository.CulinariaRecipeRepository) *ListCulinariaRecipesUseCase {
	return &ListCulinariaRecipesUseCase{repo: r}
}

func (uc *ListCulinariaRecipesUseCase) Execute(ctx context.Context, categorySlug string) ([]*entity.CulinariaRecipe, error) {
	return uc.repo.List(ctx, categorySlug)
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

// ── GetBySlug ─────────────────────────────────────────────────────────────────

type GetCulinariaRecipeBySlugUseCase struct{ repo repository.CulinariaRecipeRepository }

func NewGetCulinariaRecipeBySlugUseCase(r repository.CulinariaRecipeRepository) *GetCulinariaRecipeBySlugUseCase {
	return &GetCulinariaRecipeBySlugUseCase{repo: r}
}

func (uc *GetCulinariaRecipeBySlugUseCase) Execute(ctx context.Context, categorySlug, recipeSlug string) (*entity.CulinariaRecipe, error) {
	if categorySlug == "" || recipeSlug == "" {
		return nil, errors.New("categorySlug e recipeSlug são obrigatórios")
	}
	return uc.repo.GetBySlug(ctx, categorySlug, recipeSlug)
}

// ── Update ───────────────────────────────────────────────────────────────────

type UpdateCulinariaRecipeInput struct {
	CategoryID          string                  `json:"categoryId"`
	CategorySlug        string                  `json:"categorySlug"`
	Name                string                  `json:"name"`
	Description         string                  `json:"description"`
	PrepTimeMinutes     int                     `json:"prepTimeMinutes"`
	CookTimeMinutes     int                     `json:"cookTimeMinutes"`
	ServingsStr         string                  `json:"servingsStr"`
	Difficulty          entity.RecipeDifficulty `json:"difficulty"`
	Status              entity.RecipeStatus     `json:"status"`
	Tags                []string                `json:"tags"`
	ImageURL            string                  `json:"imageUrl"`
	YoutubeURL          string                  `json:"youtubeUrl"`
	SourceURL           string                  `json:"sourceUrl"`
	Ingredients         []string                `json:"ingredients"`
	PreparationSteps    []string                `json:"preparationSteps"`
	Utensils            string                  `json:"utensils"`
	Tips                string                  `json:"tips"`
	Substitutions       string                  `json:"substitutions"`
	StorageInstructions string                  `json:"storageInstructions"`
	EstimatedCost       float64                 `json:"estimatedCost"`
	PersonalRating      int                     `json:"personalRating"`
	Tested              bool                    `json:"tested"`
	Notes               string                  `json:"notes"`
	Active              bool                    `json:"active"`
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
	if in.PreparationSteps == nil {
		in.PreparationSteps = []string{}
	}
	recipe := &entity.CulinariaRecipe{
		CategoryID:          in.CategoryID,
		CategorySlug:        in.CategorySlug,
		Name:                in.Name,
		Description:         in.Description,
		PrepTimeMinutes:     in.PrepTimeMinutes,
		CookTimeMinutes:     in.CookTimeMinutes,
		ServingsStr:         in.ServingsStr,
		Difficulty:          in.Difficulty,
		Status:              in.Status,
		Tags:                in.Tags,
		ImageURL:            in.ImageURL,
		YoutubeURL:          in.YoutubeURL,
		SourceURL:           in.SourceURL,
		Ingredients:         in.Ingredients,
		PreparationSteps:    in.PreparationSteps,
		Utensils:            in.Utensils,
		Tips:                in.Tips,
		Substitutions:       in.Substitutions,
		StorageInstructions: in.StorageInstructions,
		EstimatedCost:       in.EstimatedCost,
		PersonalRating:      in.PersonalRating,
		Tested:              in.Tested,
		Notes:               in.Notes,
		Active:              in.Active,
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
