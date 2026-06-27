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

const vidaCriativaCollection = "vida_criativa_items"

type VidaCriativaRepository struct {
	collection *mongo.Collection
}

func NewVidaCriativaRepository(db *mongo.Database) repository.VidaCriativaRepository {
	return &VidaCriativaRepository{collection: db.Collection(vidaCriativaCollection)}
}

func (r *VidaCriativaRepository) Create(ctx context.Context, item *entity.VidaCriativaItem) (*entity.VidaCriativaItem, error) {
	item.ID = primitive.NewObjectID()
	item.CreatedAt = time.Now().UTC()
	item.UpdatedAt = time.Now().UTC()
	_, err := r.collection.InsertOne(ctx, item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (r *VidaCriativaRepository) GetByID(ctx context.Context, id string) (*entity.VidaCriativaItem, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("id inválido")
	}
	var item entity.VidaCriativaItem
	if err := r.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&item); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("item não encontrado")
		}
		return nil, err
	}
	return &item, nil
}

func (r *VidaCriativaRepository) List(ctx context.Context, category string) ([]*entity.VidaCriativaItem, error) {
	query := bson.M{}
	if category != "" {
		query["category"] = category
	}
	opts := options.Find().SetSort(bson.D{{Key: "order", Value: 1}, {Key: "created_at", Value: -1}})
	cursor, err := r.collection.Find(ctx, query, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var items []*entity.VidaCriativaItem
	if err := cursor.All(ctx, &items); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *VidaCriativaRepository) Update(ctx context.Context, id string, item *entity.VidaCriativaItem) (*entity.VidaCriativaItem, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("id inválido")
	}
	update := bson.M{"$set": bson.M{
		"title": item.Title, "description": item.Description,
		"category": item.Category, "status": item.Status,
		"tags": item.Tags, "image_url": item.ImageURL, "url": item.URL,
		"banner_color": item.BannerColor, "active": item.Active, "order": item.Order,
		"updated_at": time.Now().UTC(),
	}}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updated entity.VidaCriativaItem
	if err := r.collection.FindOneAndUpdate(ctx, bson.M{"_id": oid}, update, opts).Decode(&updated); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("item não encontrado")
		}
		return nil, err
	}
	return &updated, nil
}

func (r *VidaCriativaRepository) Delete(ctx context.Context, id string) error {
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
