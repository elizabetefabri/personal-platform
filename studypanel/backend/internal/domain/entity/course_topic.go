package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CourseTopic struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"       json:"id"`
	SectionSlug string             `bson:"section_slug"        json:"sectionSlug"`
	Slug        string             `bson:"slug"                json:"slug"`
	Label       string             `bson:"label"               json:"label"`
	Description string             `bson:"description"         json:"description"`
	BannerColor string             `bson:"banner_color"        json:"bannerColor"`
	IconClass   string             `bson:"icon_class"          json:"iconClass"`
	Skill       string             `bson:"skill"               json:"skill"`
	ImageURL    string             `bson:"image_url,omitempty" json:"imageUrl,omitempty"`
	Active      bool               `bson:"active"              json:"active"`
	Order       int                `bson:"order"               json:"order"`
	CreatedAt   time.Time          `bson:"created_at"          json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updated_at"          json:"updatedAt"`
}
