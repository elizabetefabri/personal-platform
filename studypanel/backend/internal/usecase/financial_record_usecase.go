package usecase

import (
	"context"
	"errors"
	"time"

	"studypanel-backend/internal/domain/entity"
	"studypanel-backend/internal/domain/repository"
)

// ── Create ───────────────────────────────────────────────────────────────────

type CreateFinancialRecordInput struct {
	Title       string                      `json:"title"`
	Description string                      `json:"description"`
	Amount      float64                     `json:"amount"`
	Type        entity.FinancialRecordType  `json:"type"`
	Category    string                      `json:"category"`
	Date        time.Time                   `json:"date"`
	Tags        []string                    `json:"tags"`
	Notes       string                      `json:"notes"`
}

type CreateFinancialRecordUseCase struct{ repo repository.FinancialRecordRepository }

func NewCreateFinancialRecordUseCase(r repository.FinancialRecordRepository) *CreateFinancialRecordUseCase {
	return &CreateFinancialRecordUseCase{repo: r}
}

func (uc *CreateFinancialRecordUseCase) Execute(ctx context.Context, in CreateFinancialRecordInput) (*entity.FinancialRecord, error) {
	if in.Title == "" {
		return nil, errors.New("title é obrigatório")
	}
	if !entity.ValidFinancialRecordTypes[in.Type] {
		return nil, errors.New("type deve ser 'receita' ou 'despesa'")
	}
	if in.Tags == nil {
		in.Tags = []string{}
	}
	if in.Date.IsZero() {
		in.Date = time.Now().UTC()
	}
	record := &entity.FinancialRecord{
		Title: in.Title, Description: in.Description,
		Amount: in.Amount, Type: in.Type,
		Category: in.Category, Date: in.Date,
		Tags: in.Tags, Notes: in.Notes,
		CreatedAt: time.Now().UTC(), UpdatedAt: time.Now().UTC(),
	}
	return uc.repo.Create(ctx, record)
}

// ── List ─────────────────────────────────────────────────────────────────────

type ListFinancialRecordsUseCase struct{ repo repository.FinancialRecordRepository }

func NewListFinancialRecordsUseCase(r repository.FinancialRecordRepository) *ListFinancialRecordsUseCase {
	return &ListFinancialRecordsUseCase{repo: r}
}

func (uc *ListFinancialRecordsUseCase) Execute(ctx context.Context, recordType string) ([]*entity.FinancialRecord, error) {
	return uc.repo.List(ctx, recordType)
}

// ── Get ──────────────────────────────────────────────────────────────────────

type GetFinancialRecordUseCase struct{ repo repository.FinancialRecordRepository }

func NewGetFinancialRecordUseCase(r repository.FinancialRecordRepository) *GetFinancialRecordUseCase {
	return &GetFinancialRecordUseCase{repo: r}
}

func (uc *GetFinancialRecordUseCase) Execute(ctx context.Context, id string) (*entity.FinancialRecord, error) {
	if id == "" {
		return nil, errors.New("id é obrigatório")
	}
	return uc.repo.GetByID(ctx, id)
}

// ── Update ───────────────────────────────────────────────────────────────────

type UpdateFinancialRecordInput struct {
	Title       string                      `json:"title"`
	Description string                      `json:"description"`
	Amount      float64                     `json:"amount"`
	Type        entity.FinancialRecordType  `json:"type"`
	Category    string                      `json:"category"`
	Date        time.Time                   `json:"date"`
	Tags        []string                    `json:"tags"`
	Notes       string                      `json:"notes"`
}

type UpdateFinancialRecordUseCase struct{ repo repository.FinancialRecordRepository }

func NewUpdateFinancialRecordUseCase(r repository.FinancialRecordRepository) *UpdateFinancialRecordUseCase {
	return &UpdateFinancialRecordUseCase{repo: r}
}

func (uc *UpdateFinancialRecordUseCase) Execute(ctx context.Context, id string, in UpdateFinancialRecordInput) (*entity.FinancialRecord, error) {
	if id == "" {
		return nil, errors.New("id é obrigatório")
	}
	if in.Tags == nil {
		in.Tags = []string{}
	}
	record := &entity.FinancialRecord{
		Title: in.Title, Description: in.Description,
		Amount: in.Amount, Type: in.Type,
		Category: in.Category, Date: in.Date,
		Tags: in.Tags, Notes: in.Notes,
	}
	return uc.repo.Update(ctx, id, record)
}

// ── Delete ───────────────────────────────────────────────────────────────────

type DeleteFinancialRecordUseCase struct{ repo repository.FinancialRecordRepository }

func NewDeleteFinancialRecordUseCase(r repository.FinancialRecordRepository) *DeleteFinancialRecordUseCase {
	return &DeleteFinancialRecordUseCase{repo: r}
}

func (uc *DeleteFinancialRecordUseCase) Execute(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id é obrigatório")
	}
	return uc.repo.Delete(ctx, id)
}
