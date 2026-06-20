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

const courseSectionCollection = "course_sections"

type CourseSectionRepository struct {
	collection *mongo.Collection
}

func NewCourseSectionRepository(db *mongo.Database) repository.CourseSectionRepository {
	return &CourseSectionRepository{collection: db.Collection(courseSectionCollection)}
}

func (r *CourseSectionRepository) Create(ctx context.Context, s *entity.CourseSection) (*entity.CourseSection, error) {
	s.ID = primitive.NewObjectID()
	s.CreatedAt = time.Now().UTC()
	s.UpdatedAt = time.Now().UTC()
	_, err := r.collection.InsertOne(ctx, s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (r *CourseSectionRepository) GetByID(ctx context.Context, id string) (*entity.CourseSection, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("id inválido")
	}
	var s entity.CourseSection
	if err := r.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&s); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("seção não encontrada")
		}
		return nil, err
	}
	return &s, nil
}

func (r *CourseSectionRepository) GetBySlug(ctx context.Context, slug string) (*entity.CourseSection, error) {
	var s entity.CourseSection
	if err := r.collection.FindOne(ctx, bson.M{"slug": slug}).Decode(&s); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("seção não encontrada")
		}
		return nil, err
	}
	return &s, nil
}

func (r *CourseSectionRepository) List(ctx context.Context) ([]*entity.CourseSection, error) {
	opts := options.Find().SetSort(bson.D{{Key: "order", Value: 1}, {Key: "created_at", Value: 1}})
	cursor, err := r.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var items []*entity.CourseSection
	if err := cursor.All(ctx, &items); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *CourseSectionRepository) Update(ctx context.Context, id string, s *entity.CourseSection) (*entity.CourseSection, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("id inválido")
	}
	update := bson.M{"$set": bson.M{
		"name": s.Name, "description": s.Description,
		"banner_color": s.BannerColor, "icon_class": s.IconClass,
		"image_url": s.ImageURL, "active": s.Active, "order": s.Order,
		"updated_at": time.Now().UTC(),
	}}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updated entity.CourseSection
	if err := r.collection.FindOneAndUpdate(ctx, bson.M{"_id": oid}, update, opts).Decode(&updated); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("seção não encontrada")
		}
		return nil, err
	}
	return &updated, nil
}

func (r *CourseSectionRepository) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("id inválido")
	}
	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("seção não encontrada")
	}
	return nil
}
