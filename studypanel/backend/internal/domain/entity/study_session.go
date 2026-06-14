package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StudySession struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty"           json:"id"`
	StudyItemID         string             `bson:"study_item_id"           json:"studyItemId"`
	StartedAt           time.Time          `bson:"started_at"              json:"startedAt"`
	EndedAt             time.Time          `bson:"ended_at"                json:"endedAt"`
	DurationSeconds     int                `bson:"duration_seconds"        json:"durationSeconds"`
	Topic               string             `bson:"topic,omitempty"         json:"topic,omitempty"`
	Notes               string             `bson:"notes,omitempty"         json:"notes,omitempty"`
	Rating              int                `bson:"rating"                  json:"rating"`
	MilestonesCompleted int                `bson:"milestones_completed"    json:"milestonesCompleted"`
	CreatedAt           time.Time          `bson:"created_at"              json:"createdAt"`
}
