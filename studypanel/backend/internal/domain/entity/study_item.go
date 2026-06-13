package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StudyStatus string

const (
	StatusNotStarted StudyStatus = "Não iniciado"
	StatusInProgress StudyStatus = "Em andamento"
	StatusCompleted  StudyStatus = "Concluído"
	StatusPaused     StudyStatus = "Pausado"
)

var ValidStatuses = map[StudyStatus]bool{
	StatusNotStarted: true,
	StatusInProgress: true,
	StatusCompleted:  true,
	StatusPaused:     true,
}

type StudyItem struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"     json:"id"`
	Section     string             `bson:"section"            json:"section"`
	Topic       string             `bson:"topic"              json:"topic"`
	CourseName  string             `bson:"course_name"        json:"courseName"`
	Status      StudyStatus        `bson:"status"             json:"status"`
	Date        string             `bson:"date"               json:"date"`
	URL         string             `bson:"url"                json:"url"`
	ImageURL    string             `bson:"image_url,omitempty" json:"imageUrl,omitempty"`
	DetailRoute string             `bson:"detail_route"       json:"detailRoute"`
	CreatedAt   time.Time          `bson:"created_at"         json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updated_at"         json:"updatedAt"`
}

func (s *StudyItem) IsValid() bool {
	return s.Section != "" &&
		s.Topic != "" &&
		s.CourseName != "" &&
		ValidStatuses[s.Status]
}

func BuildDetailRoute(section, topic string) string {
	return "/" + section + "/" + topic
}
