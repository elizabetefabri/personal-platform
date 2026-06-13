package usecase_test

import (
	"context"
	"errors"
	"testing"

	"studypanel-backend/internal/domain/entity"
	"studypanel-backend/internal/usecase"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetStudyItem_Success(t *testing.T) {
	repo := newMockRepo()
	uc := usecase.NewGetStudyItemUseCase(repo)

	id := primitive.NewObjectID()
	repo.items[id.Hex()] = &entity.StudyItem{
		ID:         id,
		Section:    "cloud",
		Topic:      "aws",
		CourseName: "AWS SAA",
		Status:     entity.StatusInProgress,
	}

	result, err := uc.Execute(context.Background(), id.Hex())
	if err != nil {
		t.Fatalf("esperava sem erro, recebeu: %v", err)
	}
	if result.ID != id {
		t.Errorf("IDs não coincidem: esperava %v, recebeu %v", id, result.ID)
	}
}

func TestGetStudyItem_EmptyID(t *testing.T) {
	repo := newMockRepo()
	uc := usecase.NewGetStudyItemUseCase(repo)

	_, err := uc.Execute(context.Background(), "")
	if err == nil {
		t.Fatal("esperava erro para id vazio")
	}
}

func TestGetStudyItem_NotFound(t *testing.T) {
	repo := newMockRepo()
	uc := usecase.NewGetStudyItemUseCase(repo)

	_, err := uc.Execute(context.Background(), primitive.NewObjectID().Hex())
	if err == nil {
		t.Fatal("esperava erro para item não encontrado")
	}
}

func TestGetStudyItem_RepositoryError(t *testing.T) {
	repo := newMockRepo()
	repo.err = errors.New("timeout de conexão")
	uc := usecase.NewGetStudyItemUseCase(repo)

	id := primitive.NewObjectID()
	repo.items[id.Hex()] = &entity.StudyItem{ID: id}

	_, err := uc.Execute(context.Background(), id.Hex())
	if err == nil {
		t.Fatal("esperava erro do repositório")
	}
}
