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

const quizQuestionsCollection = "quiz_questions"

type QuizQuestionRepository struct {
	collection *mongo.Collection
}

func NewQuizQuestionRepository(db *mongo.Database) *QuizQuestionRepository {
	return &QuizQuestionRepository{
		collection: db.Collection(quizQuestionsCollection),
	}
}

func (r *QuizQuestionRepository) Create(ctx context.Context, question *entity.QuizQuestion) (*entity.QuizQuestion, error) {
	question.ID = primitive.NewObjectID()
	question.CreatedAt = time.Now().UTC()

	_, err := r.collection.InsertOne(ctx, question)
	if err != nil {
		return nil, err
	}
	return question, nil
}

func (r *QuizQuestionRepository) ListByTopic(ctx context.Context, section, topic string) ([]*entity.QuizQuestion, error) {
	query := bson.M{}
	if section != "" {
		query["section"] = section
	}
	if topic != "" {
		query["topic"] = topic
	}

	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})
	cursor, err := r.collection.Find(ctx, query, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var questions []*entity.QuizQuestion
	if err := cursor.All(ctx, &questions); err != nil {
		return nil, err
	}
	return questions, nil
}

func (r *QuizQuestionRepository) Update(ctx context.Context, id string, question *entity.QuizQuestion) (*entity.QuizQuestion, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("id inválido")
	}

	update := bson.M{
		"$set": bson.M{
			"section":     question.Section,
			"topic":       question.Topic,
			"question":    question.Question,
			"answer":      question.Answer,
			"explanation": question.Explanation,
			"difficulty":  question.Difficulty,
			"tags":        question.Tags,
		},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updated entity.QuizQuestion
	err = r.collection.FindOneAndUpdate(ctx, bson.M{"_id": oid}, update, opts).Decode(&updated)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("questão não encontrada")
		}
		return nil, err
	}
	return &updated, nil
}

func (r *QuizQuestionRepository) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("id inválido")
	}

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("questão não encontrada")
	}
	return nil
}
