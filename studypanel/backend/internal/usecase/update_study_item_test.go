package usecase_test

import (
	"context"
	"errors"
	"testing"

	"studypanel-backend/internal/domain/entity"
	"studypanel-backend/internal/usecase"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUpdateStudyItem_Success(t *testing.T) {
	repo := newMockRepo()
	id := primitive.NewObjectID()
	repo.items[id.Hex()] = &entity.StudyItem{
		ID:         id,
		Section:    "cloud",
		Topic:      "aws",
		CourseName: "AWS SAA — old",
		Status:     entity.StatusNotStarted,
	}
	uc := usecase.NewUpdateStudyItemUseCase(repo)

	input := usecase.UpdateStudyItemInput{
		CourseName: "AWS SAA — updated",
		Status:     entity.StatusCompleted,
		Date:       "2024-06-15",
		URL:        "https://udemy.com/aws",
	}

	result, err := uc.Execute(context.Background(), id.Hex(), input)
	if err != nil {
		t.Fatalf("esperava sem erro, recebeu: %v", err)
	}
	if result.CourseName != "AWS SAA — updated" {
		t.Errorf("courseName não foi atualizado: %q", result.CourseName)
	}
	if result.Status != entity.StatusCompleted {
		t.Errorf("status não foi atualizado: %q", result.Status)
	}
	if result.UpdatedAt.IsZero() {
		t.Error("updatedAt não deve ser zero")
	}
}

func TestUpdateStudyItem_EmptyID(t *testing.T) {
	repo := newMockRepo()
	uc := usecase.NewUpdateStudyItemUseCase(repo)

	_, err := uc.Execute(context.Background(), "", usecase.UpdateStudyItemInput{
		CourseName: "test",
		Status:     entity.StatusCompleted,
	})
	if err == nil {
		t.Fatal("esperava erro para id vazio")
	}
}

func TestUpdateStudyItem_EmptyCourseName(t *testing.T) {
	repo := newMockRepo()
	id := primitive.NewObjectID()
	repo.items[id.Hex()] = &entity.StudyItem{ID: id}
	uc := usecase.NewUpdateStudyItemUseCase(repo)

	_, err := uc.Execute(context.Background(), id.Hex(), usecase.UpdateStudyItemInput{
		Status: entity.StatusCompleted,
	})
	if err == nil {
		t.Fatal("esperava erro para courseName vazio")
	}
}

func TestUpdateStudyItem_InvalidStatus(t *testing.T) {
	repo := newMockRepo()
	id := primitive.NewObjectID()
	repo.items[id.Hex()] = &entity.StudyItem{ID: id}
	uc := usecase.NewUpdateStudyItemUseCase(repo)

	_, err := uc.Execute(context.Background(), id.Hex(), usecase.UpdateStudyItemInput{
		CourseName: "Test",
		Status:     entity.StudyStatus("invalido"),
	})
	if err == nil {
		t.Fatal("esperava erro para status inválido")
	}
}

func TestUpdateStudyItem_NotFound(t *testing.T) {
	repo := newMockRepo()
	uc := usecase.NewUpdateStudyItemUseCase(repo)

	_, err := uc.Execute(context.Background(), primitive.NewObjectID().Hex(), usecase.UpdateStudyItemInput{
		CourseName: "Test",
		Status:     entity.StatusCompleted,
	})
	if err == nil {
		t.Fatal("esperava erro para item não encontrado")
	}
}

func TestUpdateStudyItem_RepositoryError(t *testing.T) {
	repo := newMockRepo()
	id := primitive.NewObjectID()
	repo.items[id.Hex()] = &entity.StudyItem{ID: id}
	repo.err = errors.New("erro de banco")
	uc := usecase.NewUpdateStudyItemUseCase(repo)

	_, err := uc.Execute(context.Background(), id.Hex(), usecase.UpdateStudyItemInput{
		CourseName: "Test",
		Status:     entity.StatusCompleted,
	})
	if err == nil {
		t.Fatal("esperava erro do repositório")
	}
}
