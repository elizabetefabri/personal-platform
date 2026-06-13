package mongodb

import (
	"context"
	"errors"
	"time"

	"studypanel-backend/internal/domain/entity"
	"studypanel-backend/internal/domain/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const collectionName = "study_items"

type StudyItemRepository struct {
	collection *mongo.Collection
}

func NewStudyItemRepository(db *mongo.Database) *StudyItemRepository {
	return &StudyItemRepository{
		collection: db.Collection(collectionName),
	}
}

func (r *StudyItemRepository) Create(ctx context.Context, item *entity.StudyItem) (*entity.StudyItem, error) {
	item.ID = primitive.NewObjectID()
	item.CreatedAt = time.Now().UTC()
	item.UpdatedAt = time.Now().UTC()

	_, err := r.collection.InsertOne(ctx, item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (r *StudyItemRepository) GetByID(ctx context.Context, id string) (*entity.StudyItem, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("id inválido")
	}

	var item entity.StudyItem
	err = r.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&item)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("item não encontrado")
		}
		return nil, err
	}
	return &item, nil
}

func (r *StudyItemRepository) List(ctx context.Context, filter repository.Filter) ([]*entity.StudyItem, error) {
	query := bson.M{}
	if filter.Section != "" {
		query["section"] = filter.Section
	}
	if filter.Topic != "" {
		query["topic"] = filter.Topic
	}

	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})
	cursor, err := r.collection.Find(ctx, query, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var items []*entity.StudyItem
	if err := cursor.All(ctx, &items); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *StudyItemRepository) Update(ctx context.Context, id string, item *entity.StudyItem) (*entity.StudyItem, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("id inválido")
	}

	update := bson.M{
		"$set": bson.M{
			"course_name": item.CourseName,
			"status":      item.Status,
			"date":        item.Date,
			"url":         item.URL,
			"image_url":   item.ImageURL,
			"updated_at":  time.Now().UTC(),
		},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updated entity.StudyItem
	err = r.collection.FindOneAndUpdate(ctx, bson.M{"_id": oid}, update, opts).Decode(&updated)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("item não encontrado")
		}
		return nil, err
	}
	return &updated, nil
}

func (r *StudyItemRepository) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("id inválido")
	}

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("item não encontrado")
	}
	return nil
}
