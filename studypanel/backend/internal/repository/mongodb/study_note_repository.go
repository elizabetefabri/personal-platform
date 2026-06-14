package mongodb

import (
	"context"
	"errors"
	"time"

	"studypanel-backend/internal/domain/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const studyNotesCollection = "study_notes"

type StudyNoteRepository struct {
	collection *mongo.Collection
}

func NewStudyNoteRepository(db *mongo.Database) *StudyNoteRepository {
	return &StudyNoteRepository{
		collection: db.Collection(studyNotesCollection),
	}
}

func (r *StudyNoteRepository) Create(ctx context.Context, note *entity.StudyNote) (*entity.StudyNote, error) {
	note.ID = primitive.NewObjectID()
	note.CreatedAt = time.Now().UTC()
	note.UpdatedAt = time.Now().UTC()

	_, err := r.collection.InsertOne(ctx, note)
	if err != nil {
		return nil, err
	}
	return note, nil
}

func (r *StudyNoteRepository) ListByItemID(ctx context.Context, studyItemID string) ([]*entity.StudyNote, error) {
	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})
	cursor, err := r.collection.Find(ctx, bson.M{"study_item_id": studyItemID}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var notes []*entity.StudyNote
	if err := cursor.All(ctx, &notes); err != nil {
		return nil, err
	}
	return notes, nil
}

func (r *StudyNoteRepository) Update(ctx context.Context, id string, note *entity.StudyNote) (*entity.StudyNote, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("id inválido")
	}

	update := bson.M{
		"$set": bson.M{
			"date":        note.Date,
			"title":       note.Title,
			"description": note.Description,
			"progress":    note.Progress,
			"status":      note.Status,
			"link":        note.Link,
			"updated_at":  time.Now().UTC(),
		},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updated entity.StudyNote
	err = r.collection.FindOneAndUpdate(ctx, bson.M{"_id": oid}, update, opts).Decode(&updated)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("nota não encontrada")
		}
		return nil, err
	}
	return &updated, nil
}

func (r *StudyNoteRepository) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("id inválido")
	}

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("nota não encontrada")
	}
	return nil
}
