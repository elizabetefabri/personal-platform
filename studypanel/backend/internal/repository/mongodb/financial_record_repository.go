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

const financialRecordCollection = "financial_records"

type FinancialRecordRepository struct {
	collection *mongo.Collection
}

func NewFinancialRecordRepository(db *mongo.Database) repository.FinancialRecordRepository {
	return &FinancialRecordRepository{collection: db.Collection(financialRecordCollection)}
}

func (r *FinancialRecordRepository) Create(ctx context.Context, record *entity.FinancialRecord) (*entity.FinancialRecord, error) {
	record.ID = primitive.NewObjectID()
	record.CreatedAt = time.Now().UTC()
	record.UpdatedAt = time.Now().UTC()
	_, err := r.collection.InsertOne(ctx, record)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (r *FinancialRecordRepository) GetByID(ctx context.Context, id string) (*entity.FinancialRecord, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("id inválido")
	}
	var record entity.FinancialRecord
	if err := r.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&record); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("registro não encontrado")
		}
		return nil, err
	}
	return &record, nil
}

func (r *FinancialRecordRepository) List(ctx context.Context, recordType string) ([]*entity.FinancialRecord, error) {
	query := bson.M{}
	if recordType != "" {
		query["type"] = recordType
	}
	opts := options.Find().SetSort(bson.D{{Key: "date", Value: -1}, {Key: "created_at", Value: -1}})
	cursor, err := r.collection.Find(ctx, query, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var records []*entity.FinancialRecord
	if err := cursor.All(ctx, &records); err != nil {
		return nil, err
	}
	return records, nil
}

func (r *FinancialRecordRepository) Update(ctx context.Context, id string, record *entity.FinancialRecord) (*entity.FinancialRecord, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("id inválido")
	}
	update := bson.M{"$set": bson.M{
		"title": record.Title, "description": record.Description,
		"amount": record.Amount, "type": record.Type,
		"category": record.Category, "date": record.Date,
		"tags": record.Tags, "notes": record.Notes,
		"updated_at": time.Now().UTC(),
	}}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updated entity.FinancialRecord
	if err := r.collection.FindOneAndUpdate(ctx, bson.M{"_id": oid}, update, opts).Decode(&updated); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("registro não encontrado")
		}
		return nil, err
	}
	return &updated, nil
}

func (r *FinancialRecordRepository) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("id inválido")
	}
	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("registro não encontrado")
	}
	return nil
}
