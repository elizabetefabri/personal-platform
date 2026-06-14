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

const studyResourcesCollection = "study_resources"

type StudyResourceRepository struct {
	collection *mongo.Collection
}

func NewStudyResourceRepository(db *mongo.Database) *StudyResourceRepository {
	return &StudyResourceRepository{
		collection: db.Collection(studyResourcesCollection),
	}
}

func (r *StudyResourceRepository) Create(ctx context.Context, resource *entity.StudyResource) (*entity.StudyResource, error) {
	resource.ID = primitive.NewObjectID()
	resource.CreatedAt = time.Now().UTC()
	resource.UpdatedAt = time.Now().UTC()

	_, err := r.collection.InsertOne(ctx, resource)
	if err != nil {
		return nil, err
	}
	return resource, nil
}

func (r *StudyResourceRepository) ListByItemID(ctx context.Context, studyItemID string) ([]*entity.StudyResource, error) {
	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})
	cursor, err := r.collection.Find(ctx, bson.M{"study_item_id": studyItemID}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var resources []*entity.StudyResource
	if err := cursor.All(ctx, &resources); err != nil {
		return nil, err
	}
	return resources, nil
}

func (r *StudyResourceRepository) Update(ctx context.Context, id string, resource *entity.StudyResource) (*entity.StudyResource, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("id inválido")
	}

	update := bson.M{
		"$set": bson.M{
			"title":       resource.Title,
			"url":         resource.URL,
			"type":        resource.Type,
			"description": resource.Description,
			"updated_at":  time.Now().UTC(),
		},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updated entity.StudyResource
	err = r.collection.FindOneAndUpdate(ctx, bson.M{"_id": oid}, update, opts).Decode(&updated)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("recurso não encontrado")
		}
		return nil, err
	}
	return &updated, nil
}

func (r *StudyResourceRepository) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("id inválido")
	}

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("recurso não encontrado")
	}
	return nil
}
