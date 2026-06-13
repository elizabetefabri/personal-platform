package entity_test

import (
	"testing"

	"studypanel-backend/internal/domain/entity"
)

func TestStudyItem_IsValid_AllFields(t *testing.T) {
	item := &entity.StudyItem{
		Section:    "cloud",
		Topic:      "aws",
		CourseName: "AWS SAA",
		Status:     entity.StatusInProgress,
	}
	if !item.IsValid() {
		t.Error("item válido retornou false")
	}
}

func TestStudyItem_IsValid_MissingSection(t *testing.T) {
	item := &entity.StudyItem{
		Topic:      "aws",
		CourseName: "AWS SAA",
		Status:     entity.StatusInProgress,
	}
	if item.IsValid() {
		t.Error("item sem section deve ser inválido")
	}
}

func TestStudyItem_IsValid_MissingTopic(t *testing.T) {
	item := &entity.StudyItem{
		Section:    "cloud",
		CourseName: "AWS SAA",
		Status:     entity.StatusInProgress,
	}
	if item.IsValid() {
		t.Error("item sem topic deve ser inválido")
	}
}

func TestStudyItem_IsValid_MissingCourseName(t *testing.T) {
	item := &entity.StudyItem{
		Section: "cloud",
		Topic:   "aws",
		Status:  entity.StatusInProgress,
	}
	if item.IsValid() {
		t.Error("item sem courseName deve ser inválido")
	}
}

func TestStudyItem_IsValid_InvalidStatus(t *testing.T) {
	item := &entity.StudyItem{
		Section:    "cloud",
		Topic:      "aws",
		CourseName: "AWS SAA",
		Status:     entity.StudyStatus("invalido"),
	}
	if item.IsValid() {
		t.Error("item com status inválido deve ser inválido")
	}
}

func TestStudyItem_IsValid_AllValidStatuses(t *testing.T) {
	validStatuses := []entity.StudyStatus{
		entity.StatusNotStarted,
		entity.StatusInProgress,
		entity.StatusCompleted,
		entity.StatusPaused,
	}

	for _, status := range validStatuses {
		t.Run(string(status), func(t *testing.T) {
			item := &entity.StudyItem{
				Section:    "cloud",
				Topic:      "aws",
				CourseName: "AWS SAA",
				Status:     status,
			}
			if !item.IsValid() {
				t.Errorf("status %q deve ser válido", status)
			}
		})
	}
}

func TestBuildDetailRoute(t *testing.T) {
	tests := []struct {
		section  string
		topic    string
		expected string
	}{
		{"cloud", "aws", "/cloud/aws"},
		{"frontend", "angular", "/frontend/angular"},
		{"devops", "terraform", "/devops/terraform"},
		{"banco-de-dados", "mongodb", "/banco-de-dados/mongodb"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := entity.BuildDetailRoute(tt.section, tt.topic)
			if result != tt.expected {
				t.Errorf("esperava %q, recebeu %q", tt.expected, result)
			}
		})
	}
}

func TestValidStatuses_ContainsAllStatuses(t *testing.T) {
	expected := []entity.StudyStatus{
		entity.StatusNotStarted,
		entity.StatusInProgress,
		entity.StatusCompleted,
		entity.StatusPaused,
	}

	for _, status := range expected {
		if !entity.ValidStatuses[status] {
			t.Errorf("status %q não está em ValidStatuses", status)
		}
	}
}
