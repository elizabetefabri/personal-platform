package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NoteStatus string

const (
	NoteStatusDraft    NoteStatus = "Rascunho"
	NoteStatusReviewed NoteStatus = "Revisado"
	NoteStatusFinished NoteStatus = "Finalizado"
)

var ValidNoteStatuses = map[NoteStatus]bool{
	NoteStatusDraft:    true,
	NoteStatusReviewed: true,
	NoteStatusFinished: true,
}

type NoteType string

const (
	NoteTypeAnnotation NoteType = "Anotação"
	NoteTypeSummary    NoteType = "Resumo"
	NoteTypeDoubt      NoteType = "Dúvida"
	NoteTypeWeakPoint  NoteType = "Ponto fraco"
	NoteTypeQuizError  NoteType = "Erro de quiz"
	NoteTypeRevision   NoteType = "Revisão"
)

type StudyNote struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"     json:"id"`
	StudyItemID string             `bson:"study_item_id"     json:"studyItemId"`
	Date        string             `bson:"date"              json:"date"`
	Title       string             `bson:"title"             json:"title"`
	Description string             `bson:"description"       json:"description"`
	Progress    int                `bson:"progress"          json:"progress"`
	Status      NoteStatus         `bson:"status"            json:"status"`
	Type        NoteType           `bson:"type,omitempty"    json:"type,omitempty"`
	Link        string             `bson:"link,omitempty"    json:"link,omitempty"`
	CreatedAt   time.Time          `bson:"created_at"        json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updated_at"        json:"updatedAt"`
}
