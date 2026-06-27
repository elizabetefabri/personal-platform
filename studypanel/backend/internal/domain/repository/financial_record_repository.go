package repository

import (
	"context"

	"studypanel-backend/internal/domain/entity"
)

type FinancialRecordRepository interface {
	Create(ctx context.Context, record *entity.FinancialRecord) (*entity.FinancialRecord, error)
	GetByID(ctx context.Context, id string) (*entity.FinancialRecord, error)
	List(ctx context.Context, recordType string) ([]*entity.FinancialRecord, error)
	Update(ctx context.Context, id string, record *entity.FinancialRecord) (*entity.FinancialRecord, error)
	Delete(ctx context.Context, id string) error
}
