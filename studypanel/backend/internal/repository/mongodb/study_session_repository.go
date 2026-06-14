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

const studySessionsCollection = "study_sessions"

type StudySessionRepository struct {
	collection *mongo.Collection
}

func NewStudySessionRepository(db *mongo.Database) *StudySessionRepository {
	return &StudySessionRepository{
		collection: db.Collection(studySessionsCollection),
	}
}

func (r *StudySessionRepository) Create(ctx context.Context, session *entity.StudySession) (*entity.StudySession, error) {
	session.ID = primitive.NewObjectID()
	session.CreatedAt = time.Now().UTC()

	_, err := r.collection.InsertOne(ctx, session)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (r *StudySessionRepository) ListByItemID(ctx context.Context, studyItemID string) ([]*entity.StudySession, error) {
	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})
	cursor, err := r.collection.Find(ctx, bson.M{"study_item_id": studyItemID}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var sessions []*entity.StudySession
	if err := cursor.All(ctx, &sessions); err != nil {
		return nil, err
	}
	return sessions, nil
}

func (r *StudySessionRepository) Update(ctx context.Context, id string, session *entity.StudySession) (*entity.StudySession, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("id inválido")
	}

	update := bson.M{
		"$set": bson.M{
			"started_at":           session.StartedAt,
			"ended_at":             session.EndedAt,
			"duration_seconds":     session.DurationSeconds,
			"topic":                session.Topic,
			"notes":                session.Notes,
			"rating":               session.Rating,
			"milestones_completed": session.MilestonesCompleted,
		},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updated entity.StudySession
	err = r.collection.FindOneAndUpdate(ctx, bson.M{"_id": oid}, update, opts).Decode(&updated)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("sessão não encontrada")
		}
		return nil, err
	}
	return &updated, nil
}

func (r *StudySessionRepository) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("id inválido")
	}

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("sessão não encontrada")
	}
	return nil
}
