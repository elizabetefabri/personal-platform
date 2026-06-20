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

const courseTopicCollection = "course_topics"

type CourseTopicRepository struct {
	collection *mongo.Collection
}

func NewCourseTopicRepository(db *mongo.Database) repository.CourseTopicRepository {
	return &CourseTopicRepository{collection: db.Collection(courseTopicCollection)}
}

func (r *CourseTopicRepository) Create(ctx context.Context, t *entity.CourseTopic) (*entity.CourseTopic, error) {
	t.ID = primitive.NewObjectID()
	t.CreatedAt = time.Now().UTC()
	t.UpdatedAt = time.Now().UTC()
	_, err := r.collection.InsertOne(ctx, t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (r *CourseTopicRepository) GetByID(ctx context.Context, id string) (*entity.CourseTopic, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("id inválido")
	}
	var t entity.CourseTopic
	if err := r.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&t); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("tópico não encontrado")
		}
		return nil, err
	}
	return &t, nil
}

func (r *CourseTopicRepository) List(ctx context.Context, sectionSlug string) ([]*entity.CourseTopic, error) {
	query := bson.M{}
	if sectionSlug != "" {
		query["section_slug"] = sectionSlug
	}
	opts := options.Find().SetSort(bson.D{{Key: "order", Value: 1}, {Key: "created_at", Value: 1}})
	cursor, err := r.collection.Find(ctx, query, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var items []*entity.CourseTopic
	if err := cursor.All(ctx, &items); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *CourseTopicRepository) Update(ctx context.Context, id string, t *entity.CourseTopic) (*entity.CourseTopic, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("id inválido")
	}
	update := bson.M{"$set": bson.M{
		"label": t.Label, "description": t.Description,
		"banner_color": t.BannerColor, "icon_class": t.IconClass,
		"skill": t.Skill, "image_url": t.ImageURL,
		"active": t.Active, "order": t.Order,
		"updated_at": time.Now().UTC(),
	}}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updated entity.CourseTopic
	if err := r.collection.FindOneAndUpdate(ctx, bson.M{"_id": oid}, update, opts).Decode(&updated); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("tópico não encontrado")
		}
		return nil, err
	}
	return &updated, nil
}

func (r *CourseTopicRepository) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("id inválido")
	}
	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("tópico não encontrado")
	}
	return nil
}
