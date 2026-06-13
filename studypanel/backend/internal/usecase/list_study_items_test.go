package usecase_test

import (
	"context"
	"errors"
	"testing"

	"studypanel-backend/internal/domain/entity"
	"studypanel-backend/internal/usecase"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func seedItems(repo *mockRepo) {
	items := []*entity.StudyItem{
		{ID: primitive.NewObjectID(), Section: "cloud", Topic: "aws", CourseName: "AWS SAA", Status: entity.StatusInProgress},
		{ID: primitive.NewObjectID(), Section: "cloud", Topic: "azure", CourseName: "Azure Fundamentals", Status: entity.StatusNotStarted},
		{ID: primitive.NewObjectID(), Section: "frontend", Topic: "angular", CourseName: "Angular Avançado", Status: entity.StatusCompleted},
	}
	for _, item := range items {
		repo.items[item.ID.Hex()] = item
	}
}

func TestListStudyItems_NoFilter(t *testing.T) {
	repo := newMockRepo()
	seedItems(repo)
	uc := usecase.NewListStudyItemsUseCase(repo)

	result, err := uc.Execute(context.Background(), usecase.ListStudyItemsInput{})
	if err != nil {
		t.Fatalf("esperava sem erro, recebeu: %v", err)
	}
	if len(result) != 3 {
		t.Errorf("esperava 3 itens, recebeu %d", len(result))
	}
}

func TestListStudyItems_FilterBySection(t *testing.T) {
	repo := newMockRepo()
	seedItems(repo)
	uc := usecase.NewListStudyItemsUseCase(repo)

	result, err := uc.Execute(context.Background(), usecase.ListStudyItemsInput{Section: "cloud"})
	if err != nil {
		t.Fatalf("esperava sem erro, recebeu: %v", err)
	}
	if len(result) != 2 {
		t.Errorf("esperava 2 itens cloud, recebeu %d", len(result))
	}
	for _, item := range result {
		if item.Section != "cloud" {
			t.Errorf("item com section incorreta: %q", item.Section)
		}
	}
}

func TestListStudyItems_FilterByTopic(t *testing.T) {
	repo := newMockRepo()
	seedItems(repo)
	uc := usecase.NewListStudyItemsUseCase(repo)

	result, err := uc.Execute(context.Background(), usecase.ListStudyItemsInput{Section: "cloud", Topic: "aws"})
	if err != nil {
		t.Fatalf("esperava sem erro, recebeu: %v", err)
	}
	if len(result) != 1 {
		t.Errorf("esperava 1 item aws, recebeu %d", len(result))
	}
	if result[0].Topic != "aws" {
		t.Errorf("esperava topic=aws, recebeu=%q", result[0].Topic)
	}
}

func TestListStudyItems_EmptyResult(t *testing.T) {
	repo := newMockRepo()
	uc := usecase.NewListStudyItemsUseCase(repo)

	result, err := uc.Execute(context.Background(), usecase.ListStudyItemsInput{Section: "nao-existe"})
	if err != nil {
		t.Fatalf("esperava sem erro, recebeu: %v", err)
	}
	if len(result) != 0 {
		t.Errorf("esperava 0 itens, recebeu %d", len(result))
	}
}

func TestListStudyItems_RepositoryError(t *testing.T) {
	repo := newMockRepo()
	repo.err = errors.New("erro de conexão")
	uc := usecase.NewListStudyItemsUseCase(repo)

	_, err := uc.Execute(context.Background(), usecase.ListStudyItemsInput{})
	if err == nil {
		t.Fatal("esperava erro do repositório")
	}
}
