package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FinancialRecordType classifica o registro como receita ou despesa.
type FinancialRecordType string

const (
	FinancialRecordTypeReceita FinancialRecordType = "receita"
	FinancialRecordTypeDespesa FinancialRecordType = "despesa"
)

var ValidFinancialRecordTypes = map[FinancialRecordType]bool{
	FinancialRecordTypeReceita: true,
	FinancialRecordTypeDespesa: true,
}

// FinancialRecord representa uma transação financeira no Painel Financeiro.
type FinancialRecord struct {
	ID          primitive.ObjectID  `bson:"_id,omitempty"       json:"id"`
	Title       string              `bson:"title"               json:"title"`
	Description string              `bson:"description"         json:"description"`
	Amount      float64             `bson:"amount"              json:"amount"`
	Type        FinancialRecordType `bson:"type"                json:"type"`
	Category    string              `bson:"category"            json:"category"`
	Date        time.Time           `bson:"date"                json:"date"`
	Tags        []string            `bson:"tags"                json:"tags"`
	Notes       string              `bson:"notes,omitempty"     json:"notes,omitempty"`
	CreatedAt   time.Time           `bson:"created_at"          json:"createdAt"`
	UpdatedAt   time.Time           `bson:"updated_at"          json:"updatedAt"`
}
