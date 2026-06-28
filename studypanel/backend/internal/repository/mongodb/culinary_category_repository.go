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

const culinaryCategoryCollection = "culinary_categories"

type CulinaryCategoryRepository struct {
	collection *mongo.Collection
}

func NewCulinaryCategoryRepository(db *mongo.Database) repository.CulinaryCategoryRepository {
	return &CulinaryCategoryRepository{collection: db.Collection(culinaryCategoryCollection)}
}

func (r *CulinaryCategoryRepository) Create(ctx context.Context, cat *entity.CulinaryCategory) (*entity.CulinaryCategory, error) {
	cat.ID = primitive.NewObjectID()
	cat.CreatedAt = time.Now().UTC()
	cat.UpdatedAt = time.Now().UTC()
	_, err := r.collection.InsertOne(ctx, cat)
	if err != nil {
		return nil, err
	}
	return cat, nil
}

func (r *CulinaryCategoryRepository) GetByID(ctx context.Context, id string) (*entity.CulinaryCategory, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("id inválido")
	}
	var cat entity.CulinaryCategory
	if err := r.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&cat); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("categoria não encontrada")
		}
		return nil, err
	}
	return &cat, nil
}

func (r *CulinaryCategoryRepository) GetBySlug(ctx context.Context, slug string) (*entity.CulinaryCategory, error) {
	var cat entity.CulinaryCategory
	if err := r.collection.FindOne(ctx, bson.M{"slug": slug}).Decode(&cat); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("categoria não encontrada")
		}
		return nil, err
	}
	return &cat, nil
}

func (r *CulinaryCategoryRepository) List(ctx context.Context) ([]*entity.CulinaryCategory, error) {
	opts := options.Find().SetSort(bson.D{{Key: "order", Value: 1}, {Key: "created_at", Value: 1}})
	cursor, err := r.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var items []*entity.CulinaryCategory
	if err := cursor.All(ctx, &items); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *CulinaryCategoryRepository) Update(ctx context.Context, id string, cat *entity.CulinaryCategory) (*entity.CulinaryCategory, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("id inválido")
	}
	update := bson.M{"$set": bson.M{
		"name":        cat.Name,
		"description": cat.Description,
		"tag":         cat.Tag,
		"color":       cat.Color,
		"icon":        cat.Icon,
		"image_url":   cat.ImageURL,
		"order":       cat.Order,
		"active":      cat.Active,
		"updated_at":  time.Now().UTC(),
	}}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updated entity.CulinaryCategory
	if err := r.collection.FindOneAndUpdate(ctx, bson.M{"_id": oid}, update, opts).Decode(&updated); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("categoria não encontrada")
		}
		return nil, err
	}
	return &updated, nil
}

func (r *CulinaryCategoryRepository) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("id inválido")
	}
	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("categoria não encontrada")
	}
	return nil
}
