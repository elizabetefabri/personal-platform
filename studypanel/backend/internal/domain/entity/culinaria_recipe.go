package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// RecipeDifficulty representa o nível de dificuldade de uma receita.
type RecipeDifficulty string

const (
	RecipeDifficultyFacil  RecipeDifficulty = "Fácil"
	RecipeDifficultyMedio  RecipeDifficulty = "Médio"
	RecipeDifficultyDificil RecipeDifficulty = "Difícil"
)

var ValidRecipeDifficulties = map[RecipeDifficulty]bool{
	RecipeDifficultyFacil:   true,
	RecipeDifficultyMedio:   true,
	RecipeDifficultyDificil: true,
}

// CulinariaRecipe representa uma receita na seção Culinária.
type CulinariaRecipe struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"         json:"id"`
	Title        string             `bson:"title"                 json:"title"`
	Description  string             `bson:"description"           json:"description"`
	Category     string             `bson:"category"              json:"category"`
	Ingredients  []string           `bson:"ingredients"           json:"ingredients"`
	Instructions []string           `bson:"instructions"          json:"instructions"`
	PrepTime     int                `bson:"prep_time"             json:"prepTime"`
	Difficulty   RecipeDifficulty   `bson:"difficulty"            json:"difficulty"`
	Tags         []string           `bson:"tags"                  json:"tags"`
	ImageURL     string             `bson:"image_url,omitempty"   json:"imageUrl,omitempty"`
	Servings     int                `bson:"servings"              json:"servings"`
	Active       bool               `bson:"active"                json:"active"`
	CreatedAt    time.Time          `bson:"created_at"            json:"createdAt"`
	UpdatedAt    time.Time          `bson:"updated_at"            json:"updatedAt"`
}
