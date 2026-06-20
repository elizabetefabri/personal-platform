package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProjectType string

const (
	ProjectTypePessoal      ProjectType = "pessoal"
	ProjectTypeProfissional ProjectType = "profissional"
)

var ValidProjectTypes = map[ProjectType]bool{
	ProjectTypePessoal:      true,
	ProjectTypeProfissional: true,
}

type Project struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"       json:"id"`
	Name        string             `bson:"name"                json:"name"`
	Type        ProjectType        `bson:"type"                json:"type"`
	Description string             `bson:"description"         json:"description"`
	Tags        []string           `bson:"tags"                json:"tags"`
	RepoURL     string             `bson:"repo_url,omitempty"  json:"repoUrl,omitempty"`
	DeployURL   string             `bson:"deploy_url,omitempty" json:"deployUrl,omitempty"`
	Slug        string             `bson:"slug"                json:"slug"`
	BannerColor string             `bson:"banner_color"        json:"bannerColor"`
	ImageURL    string             `bson:"image_url,omitempty" json:"imageUrl,omitempty"`
	ImageAlt    string             `bson:"image_alt,omitempty" json:"imageAlt,omitempty"`
	DetailRoute string             `bson:"detail_route,omitempty" json:"detailRoute,omitempty"`
	Active      bool               `bson:"active"              json:"active"`
	Order       int                `bson:"order"               json:"order"`
	CreatedAt   time.Time          `bson:"created_at"          json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updated_at"          json:"updatedAt"`
}
