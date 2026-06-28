package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// RecipeDifficulty representa o nível de dificuldade de uma receita.
type RecipeDifficulty string

const (
	RecipeDifficultyFacil   RecipeDifficulty = "Fácil"
	RecipeDifficultyMedio   RecipeDifficulty = "Médio"
	RecipeDifficultyDificil RecipeDifficulty = "Difícil"
)

var ValidRecipeDifficulties = map[RecipeDifficulty]bool{
	RecipeDifficultyFacil:   true,
	RecipeDifficultyMedio:   true,
	RecipeDifficultyDificil: true,
}

// RecipeStatus representa o status de uma receita.
type RecipeStatus string

const (
	RecipeStatusPending         RecipeStatus = "pending"
	RecipeStatusTested          RecipeStatus = "tested"
	RecipeStatusFavorite        RecipeStatus = "favorite"
	RecipeStatusNeedsAdjustment RecipeStatus = "needs_adjustment"
)

// CulinariaRecipe representa uma receita na seção Culinária.
type CulinariaRecipe struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty"                  json:"id"`
	CategoryID          string             `bson:"category_id"                    json:"categoryId"`
	CategorySlug        string             `bson:"category_slug"                  json:"categorySlug"`
	Name                string             `bson:"name"                           json:"name"`
	Slug                string             `bson:"slug"                           json:"slug"`
	// Legacy fields kept for backward compatibility
	Title        string           `bson:"title"                          json:"title"`
	Category     string           `bson:"category"                       json:"category"`
	Servings     int              `bson:"servings"                       json:"servings"`
	Description         string             `bson:"description"                    json:"description"`
	PrepTime            int                `bson:"prep_time"                      json:"prepTime"`
	PrepTimeMinutes     int                `bson:"prep_time_minutes"              json:"prepTimeMinutes"`
	CookTimeMinutes     int                `bson:"cook_time_minutes"              json:"cookTimeMinutes"`
	Difficulty          RecipeDifficulty   `bson:"difficulty"                     json:"difficulty"`
	Status              RecipeStatus       `bson:"status"                         json:"status"`
	Tags                []string           `bson:"tags"                           json:"tags"`
	ImageURL            string             `bson:"image_url,omitempty"            json:"imageUrl,omitempty"`
	YoutubeURL          string             `bson:"youtube_url,omitempty"          json:"youtubeUrl,omitempty"`
	SourceURL           string             `bson:"source_url,omitempty"           json:"sourceUrl,omitempty"`
	Ingredients         []string           `bson:"ingredients"                    json:"ingredients"`
	Instructions        []string           `bson:"instructions"                   json:"instructions"`
	PreparationSteps    []string           `bson:"preparation_steps"              json:"preparationSteps"`
	ServingsStr         string             `bson:"servings_str"                   json:"servingsStr"`
	Utensils            string             `bson:"utensils,omitempty"             json:"utensils,omitempty"`
	Tips                string             `bson:"tips,omitempty"                 json:"tips,omitempty"`
	Substitutions       string             `bson:"substitutions,omitempty"        json:"substitutions,omitempty"`
	StorageInstructions string             `bson:"storage_instructions,omitempty" json:"storageInstructions,omitempty"`
	EstimatedCost       float64            `bson:"estimated_cost"                 json:"estimatedCost"`
	PersonalRating      int                `bson:"personal_rating"                json:"personalRating"`
	Tested              bool               `bson:"tested"                         json:"tested"`
	TestedAt            *time.Time         `bson:"tested_at,omitempty"            json:"testedAt,omitempty"`
	Notes               string             `bson:"notes,omitempty"                json:"notes,omitempty"`
	Active              bool               `bson:"active"                         json:"active"`
	CreatedAt           time.Time          `bson:"created_at"                     json:"createdAt"`
	UpdatedAt           time.Time          `bson:"updated_at"                     json:"updatedAt"`
}
