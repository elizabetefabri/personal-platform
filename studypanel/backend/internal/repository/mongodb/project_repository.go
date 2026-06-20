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

const projectCollection = "projects"

type ProjectRepository struct {
	collection *mongo.Collection
}

func NewProjectRepository(db *mongo.Database) repository.ProjectRepository {
	return &ProjectRepository{collection: db.Collection(projectCollection)}
}

func (r *ProjectRepository) Create(ctx context.Context, p *entity.Project) (*entity.Project, error) {
	p.ID = primitive.NewObjectID()
	p.CreatedAt = time.Now().UTC()
	p.UpdatedAt = time.Now().UTC()
	_, err := r.collection.InsertOne(ctx, p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *ProjectRepository) GetByID(ctx context.Context, id string) (*entity.Project, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("id inválido")
	}
	var p entity.Project
	if err := r.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&p); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("projeto não encontrado")
		}
		return nil, err
	}
	return &p, nil
}

func (r *ProjectRepository) List(ctx context.Context, projectType string) ([]*entity.Project, error) {
	query := bson.M{}
	if projectType != "" {
		query["type"] = projectType
	}
	opts := options.Find().SetSort(bson.D{{Key: "order", Value: 1}, {Key: "created_at", Value: -1}})
	cursor, err := r.collection.Find(ctx, query, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var items []*entity.Project
	if err := cursor.All(ctx, &items); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *ProjectRepository) Update(ctx context.Context, id string, p *entity.Project) (*entity.Project, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("id inválido")
	}
	update := bson.M{"$set": bson.M{
		"name": p.Name, "description": p.Description, "tags": p.Tags,
		"repo_url": p.RepoURL, "deploy_url": p.DeployURL,
		"banner_color": p.BannerColor, "image_url": p.ImageURL, "image_alt": p.ImageAlt,
		"detail_route": p.DetailRoute, "active": p.Active, "order": p.Order,
		"updated_at": time.Now().UTC(),
	}}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updated entity.Project
	if err := r.collection.FindOneAndUpdate(ctx, bson.M{"_id": oid}, update, opts).Decode(&updated); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("projeto não encontrado")
		}
		return nil, err
	}
	return &updated, nil
}

func (r *ProjectRepository) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("id inválido")
	}
	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("projeto não encontrado")
	}
	return nil
}
