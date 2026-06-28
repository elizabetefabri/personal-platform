package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CulinaryCategory representa uma categoria culinária (ex: Massas, Sobremesas).
type CulinaryCategory struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"        json:"id"`
	Name        string             `bson:"name"                 json:"name"`
	Slug        string             `bson:"slug"                 json:"slug"`
	Description string             `bson:"description"          json:"description"`
	Tag         string             `bson:"tag"                  json:"tag"`
	Color       string             `bson:"color"                json:"color"`
	Icon        string             `bson:"icon"                 json:"icon"`
	ImageURL    string             `bson:"image_url,omitempty"  json:"imageUrl,omitempty"`
	Order       int                `bson:"order"                json:"order"`
	Active      bool               `bson:"active"               json:"active"`
	CreatedAt   time.Time          `bson:"created_at"           json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updated_at"           json:"updatedAt"`
}
