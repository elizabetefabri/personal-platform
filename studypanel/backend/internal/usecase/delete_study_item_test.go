package usecase_test

import (
	"context"
	"errors"
	"testing"

	"studypanel-backend/internal/domain/entity"
	"studypanel-backend/internal/usecase"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestDeleteStudyItem_Success(t *testing.T) {
	repo := newMockRepo()
	id := primitive.NewObjectID()
	repo.items[id.Hex()] = &entity.StudyItem{
		ID:         id,
		CourseName: "AWS SAA",
		Status:     entity.StatusInProgress,
	}
	uc := usecase.NewDeleteStudyItemUseCase(repo)

	err := uc.Execute(context.Background(), id.Hex())
	if err != nil {
		t.Fatalf("esperava sem erro, recebeu: %v", err)
	}
	if _, exists := repo.items[id.Hex()]; exists {
		t.Error("item ainda existe após delete")
	}
}

func TestDeleteStudyItem_EmptyID(t *testing.T) {
	repo := newMockRepo()
	uc := usecase.NewDeleteStudyItemUseCase(repo)

	err := uc.Execute(context.Background(), "")
	if err == nil {
		t.Fatal("esperava erro para id vazio")
	}
}

func TestDeleteStudyItem_NotFound(t *testing.T) {
	repo := newMockRepo()
	uc := usecase.NewDeleteStudyItemUseCase(repo)

	err := uc.Execute(context.Background(), primitive.NewObjectID().Hex())
	if err == nil {
		t.Fatal("esperava erro para item não encontrado")
	}
}

func TestDeleteStudyItem_RepositoryError(t *testing.T) {
	repo := newMockRepo()
	id := primitive.NewObjectID()
	repo.items[id.Hex()] = &entity.StudyItem{ID: id}
	repo.err = errors.New("falha de conexão")
	uc := usecase.NewDeleteStudyItemUseCase(repo)

	err := uc.Execute(context.Background(), id.Hex())
	if err == nil {
		t.Fatal("esperava erro do repositório")
	}
}
