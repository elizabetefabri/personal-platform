package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ResourceType string

const (
	ResourceTypeDoc     ResourceType = "Documentação"
	ResourceTypeVideo   ResourceType = "Vídeo"
	ResourceTypeArticle ResourceType = "Artigo"
	ResourceTypeTool    ResourceType = "Ferramenta"
	ResourceTypeBook    ResourceType = "Livro"
	ResourceTypeOther   ResourceType = "Outro"
)

var ValidResourceTypes = map[ResourceType]bool{
	ResourceTypeDoc:     true,
	ResourceTypeVideo:   true,
	ResourceTypeArticle: true,
	ResourceTypeTool:    true,
	ResourceTypeBook:    true,
	ResourceTypeOther:   true,
}

type StudyResource struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"          json:"id"`
	StudyItemID string             `bson:"study_item_id"          json:"studyItemId"`
	Title       string             `bson:"title"                  json:"title"`
	URL         string             `bson:"url"                    json:"url"`
	Type        ResourceType       `bson:"type"                   json:"type"`
	Description string             `bson:"description,omitempty"  json:"description,omitempty"`
	CreatedAt   time.Time          `bson:"created_at"             json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updated_at"             json:"updatedAt"`
}
