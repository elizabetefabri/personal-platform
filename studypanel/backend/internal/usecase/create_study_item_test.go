package usecase_test

import (
	"context"
	"errors"
	"testing"

	"studypanel-backend/internal/domain/entity"
	"studypanel-backend/internal/usecase"
)

func TestCreateStudyItem_Success(t *testing.T) {
	repo := newMockRepo()
	uc := usecase.NewCreateStudyItemUseCase(repo)

	input := usecase.CreateStudyItemInput{
		Section:    "cloud",
		Topic:      "aws",
		CourseName: "AWS Solutions Architect Associate",
		Status:     entity.StatusInProgress,
		Date:       "2024-06-13",
		URL:        "https://udemy.com/course/aws-saa",
		ImageURL:   "https://example.com/aws.png",
	}

	result, err := uc.Execute(context.Background(), input)

	if err != nil {
		t.Fatalf("esperava sem erro, recebeu: %v", err)
	}
	if result == nil {
		t.Fatal("resultado não deve ser nil")
	}
	if result.CourseName != input.CourseName {
		t.Errorf("esperava courseName=%q, recebeu=%q", input.CourseName, result.CourseName)
	}
	if result.DetailRoute != "/cloud/aws" {
		t.Errorf("esperava detailRoute=/cloud/aws, recebeu=%q", result.DetailRoute)
	}
	if result.ID.IsZero() {
		t.Error("ID não deve ser zero")
	}
	if result.CreatedAt.IsZero() {
		t.Error("CreatedAt não deve ser zero")
	}
}

func TestCreateStudyItem_MissingSection(t *testing.T) {
	repo := newMockRepo()
	uc := usecase.NewCreateStudyItemUseCase(repo)

	input := usecase.CreateStudyItemInput{
		Topic:      "aws",
		CourseName: "AWS SAA",
		Status:     entity.StatusNotStarted,
	}

	_, err := uc.Execute(context.Background(), input)
	if err == nil {
		t.Fatal("esperava erro para section vazio")
	}
}

func TestCreateStudyItem_MissingTopic(t *testing.T) {
	repo := newMockRepo()
	uc := usecase.NewCreateStudyItemUseCase(repo)

	input := usecase.CreateStudyItemInput{
		Section:    "cloud",
		CourseName: "AWS SAA",
		Status:     entity.StatusNotStarted,
	}

	_, err := uc.Execute(context.Background(), input)
	if err == nil {
		t.Fatal("esperava erro para topic vazio")
	}
}

func TestCreateStudyItem_MissingCourseName(t *testing.T) {
	repo := newMockRepo()
	uc := usecase.NewCreateStudyItemUseCase(repo)

	input := usecase.CreateStudyItemInput{
		Section: "cloud",
		Topic:   "aws",
		Status:  entity.StatusNotStarted,
	}

	_, err := uc.Execute(context.Background(), input)
	if err == nil {
		t.Fatal("esperava erro para courseName vazio")
	}
}

func TestCreateStudyItem_InvalidStatus(t *testing.T) {
	repo := newMockRepo()
	uc := usecase.NewCreateStudyItemUseCase(repo)

	input := usecase.CreateStudyItemInput{
		Section:    "cloud",
		Topic:      "aws",
		CourseName: "AWS SAA",
		Status:     entity.StudyStatus("status-invalido"),
	}

	_, err := uc.Execute(context.Background(), input)
	if err == nil {
		t.Fatal("esperava erro para status inválido")
	}
}

func TestCreateStudyItem_RepositoryError(t *testing.T) {
	repo := newMockRepo()
	repo.err = errors.New("erro de banco de dados")
	uc := usecase.NewCreateStudyItemUseCase(repo)

	input := usecase.CreateStudyItemInput{
		Section:    "cloud",
		Topic:      "aws",
		CourseName: "AWS SAA",
		Status:     entity.StatusNotStarted,
	}

	_, err := uc.Execute(context.Background(), input)
	if err == nil {
		t.Fatal("esperava erro do repositório")
	}
}

func TestCreateStudyItem_AllStatuses(t *testing.T) {
	statuses := []entity.StudyStatus{
		entity.StatusNotStarted,
		entity.StatusInProgress,
		entity.StatusCompleted,
		entity.StatusPaused,
	}

	for _, status := range statuses {
		t.Run(string(status), func(t *testing.T) {
			repo := newMockRepo()
			uc := usecase.NewCreateStudyItemUseCase(repo)

			input := usecase.CreateStudyItemInput{
				Section:    "frontend",
				Topic:      "angular",
				CourseName: "Angular Avançado",
				Status:     status,
			}

			result, err := uc.Execute(context.Background(), input)
			if err != nil {
				t.Fatalf("status %q não deveria retornar erro: %v", status, err)
			}
			if result.Status != status {
				t.Errorf("esperava status=%q, recebeu=%q", status, result.Status)
			}
		})
	}
}
