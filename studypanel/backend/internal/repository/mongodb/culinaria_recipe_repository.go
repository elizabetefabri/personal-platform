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

const culinariaRecipeCollection = "culinaria_recipes"

type CulinariaRecipeRepository struct {
	collection *mongo.Collection
}

func NewCulinariaRecipeRepository(db *mongo.Database) repository.CulinariaRecipeRepository {
	return &CulinariaRecipeRepository{collection: db.Collection(culinariaRecipeCollection)}
}

func (r *CulinariaRecipeRepository) Create(ctx context.Context, recipe *entity.CulinariaRecipe) (*entity.CulinariaRecipe, error) {
	recipe.ID = primitive.NewObjectID()
	recipe.CreatedAt = time.Now().UTC()
	recipe.UpdatedAt = time.Now().UTC()
	_, err := r.collection.InsertOne(ctx, recipe)
	if err != nil {
		return nil, err
	}
	return recipe, nil
}

func (r *CulinariaRecipeRepository) GetByID(ctx context.Context, id string) (*entity.CulinariaRecipe, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("id inválido")
	}
	var recipe entity.CulinariaRecipe
	if err := r.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&recipe); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("receita não encontrada")
		}
		return nil, err
	}
	return &recipe, nil
}

func (r *CulinariaRecipeRepository) List(ctx context.Context, category string) ([]*entity.CulinariaRecipe, error) {
	query := bson.M{}
	if category != "" {
		query["category"] = category
	}
	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})
	cursor, err := r.collection.Find(ctx, query, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var recipes []*entity.CulinariaRecipe
	if err := cursor.All(ctx, &recipes); err != nil {
		return nil, err
	}
	return recipes, nil
}

func (r *CulinariaRecipeRepository) Update(ctx context.Context, id string, recipe *entity.CulinariaRecipe) (*entity.CulinariaRecipe, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("id inválido")
	}
	update := bson.M{"$set": bson.M{
		"title": recipe.Title, "description": recipe.Description,
		"category": recipe.Category, "ingredients": recipe.Ingredients,
		"instructions": recipe.Instructions, "prep_time": recipe.PrepTime,
		"difficulty": recipe.Difficulty, "tags": recipe.Tags,
		"image_url": recipe.ImageURL, "servings": recipe.Servings,
		"active": recipe.Active, "updated_at": time.Now().UTC(),
	}}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updated entity.CulinariaRecipe
	if err := r.collection.FindOneAndUpdate(ctx, bson.M{"_id": oid}, update, opts).Decode(&updated); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("receita não encontrada")
		}
		return nil, err
	}
	return &updated, nil
}

func (r *CulinariaRecipeRepository) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("id inválido")
	}
	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("receita não encontrada")
	}
	return nil
}
