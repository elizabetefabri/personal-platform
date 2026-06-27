package usecase_test

import (
	"context"
	"errors"
	"testing"

	"studypanel-backend/internal/domain/entity"
	"studypanel-backend/internal/usecase"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ── mock ──────────────────────────────────────────────────────────────────────

type mockProjectRepo struct {
	items map[string]*entity.Project
	err   error
}

func newMockProjectRepo() *mockProjectRepo {
	return &mockProjectRepo{items: make(map[string]*entity.Project)}
}

func (m *mockProjectRepo) Create(_ context.Context, p *entity.Project) (*entity.Project, error) {
	if m.err != nil {
		return nil, m.err
	}
	p.ID = primitive.NewObjectID()
	m.items[p.ID.Hex()] = p
	return p, nil
}

func (m *mockProjectRepo) GetByID(_ context.Context, id string) (*entity.Project, error) {
	if m.err != nil {
		return nil, m.err
	}
	p, ok := m.items[id]
	if !ok {
		return nil, errors.New("projeto não encontrado")
	}
	return p, nil
}

func (m *mockProjectRepo) List(_ context.Context, projectType string) ([]*entity.Project, error) {
	if m.err != nil {
		return nil, m.err
	}
	var result []*entity.Project
	for _, p := range m.items {
		if projectType != "" && string(p.Type) != projectType {
			continue
		}
		result = append(result, p)
	}
	return result, nil
}

func (m *mockProjectRepo) Update(_ context.Context, id string, p *entity.Project) (*entity.Project, error) {
	if m.err != nil {
		return nil, m.err
	}
	m.items[id] = p
	return p, nil
}

func (m *mockProjectRepo) Delete(_ context.Context, id string) error {
	if m.err != nil {
		return m.err
	}
	delete(m.items, id)
	return nil
}

// ── tests ─────────────────────────────────────────────────────────────────────

func TestCreateProject_Success(t *testing.T) {
	repo := newMockProjectRepo()
	uc := usecase.NewCreateProjectUseCase(repo)

	input := usecase.CreateProjectInput{
		Name:        "StudyPanel",
		Type:        entity.ProjectTypePessoal,
		Description: "Plataforma de estudos pessoal",
		Tags:        []string{"go", "angular"},
		Slug:        "studypanel",
		BannerColor: "#3B82F6",
		Active:      true,
		Order:       1,
	}

	result, err := uc.Execute(context.Background(), input)

	if err != nil {
		t.Fatalf("esperava sem erro, recebeu: %v", err)
	}
	if result == nil {
		t.Fatal("resultado não deve ser nil")
	}
	if result.Name != input.Name {
		t.Errorf("esperava name=%q, recebeu=%q", input.Name, result.Name)
	}
	if result.Type != input.Type {
		t.Errorf("esperava type=%q, recebeu=%q", input.Type, result.Type)
	}
	if result.ID.IsZero() {
		t.Error("ID não deve ser zero")
	}
	if result.CreatedAt.IsZero() {
		t.Error("CreatedAt não deve ser zero")
	}
}

func TestCreateProject_MissingName(t *testing.T) {
	repo := newMockProjectRepo()
	uc := usecase.NewCreateProjectUseCase(repo)

	input := usecase.CreateProjectInput{
		Type: entity.ProjectTypePessoal,
	}

	_, err := uc.Execute(context.Background(), input)
	if err == nil {
		t.Fatal("esperava erro para name vazio")
	}
}

func TestCreateProject_InvalidType(t *testing.T) {
	repo := newMockProjectRepo()
	uc := usecase.NewCreateProjectUseCase(repo)

	input := usecase.CreateProjectInput{
		Name: "Projeto Inválido",
		Type: entity.ProjectType("tipo-invalido"),
	}

	_, err := uc.Execute(context.Background(), input)
	if err == nil {
		t.Fatal("esperava erro para type inválido")
	}
}

func TestCreateProject_RepositoryError(t *testing.T) {
	repo := newMockProjectRepo()
	repo.err = errors.New("erro de banco de dados")
	uc := usecase.NewCreateProjectUseCase(repo)

	input := usecase.CreateProjectInput{
		Name: "StudyPanel",
		Type: entity.ProjectTypePessoal,
	}

	_, err := uc.Execute(context.Background(), input)
	if err == nil {
		t.Fatal("esperava erro do repositório")
	}
}

func TestListProjects_ByType(t *testing.T) {
	repo := newMockProjectRepo()
	createUC := usecase.NewCreateProjectUseCase(repo)
	listUC := usecase.NewListProjectsUseCase(repo)

	projects := []usecase.CreateProjectInput{
		{Name: "Pessoal 1", Type: entity.ProjectTypePessoal},
		{Name: "Pessoal 2", Type: entity.ProjectTypePessoal},
		{Name: "Profissional 1", Type: entity.ProjectTypeProfissional},
	}
	for _, p := range projects {
		if _, err := createUC.Execute(context.Background(), p); err != nil {
			t.Fatalf("erro ao criar projeto: %v", err)
		}
	}

	pessoais, err := listUC.Execute(context.Background(), "pessoal")
	if err != nil {
		t.Fatalf("esperava sem erro, recebeu: %v", err)
	}
	if len(pessoais) != 2 {
		t.Errorf("esperava 2 projetos pessoais, recebeu %d", len(pessoais))
	}

	profissionais, err := listUC.Execute(context.Background(), "profissional")
	if err != nil {
		t.Fatalf("esperava sem erro, recebeu: %v", err)
	}
	if len(profissionais) != 1 {
		t.Errorf("esperava 1 projeto profissional, recebeu %d", len(profissionais))
	}

	todos, err := listUC.Execute(context.Background(), "")
	if err != nil {
		t.Fatalf("esperava sem erro, recebeu: %v", err)
	}
	if len(todos) != 3 {
		t.Errorf("esperava 3 projetos no total, recebeu %d", len(todos))
	}
}

func TestGetProject_Success(t *testing.T) {
	repo := newMockProjectRepo()
	createUC := usecase.NewCreateProjectUseCase(repo)
	getUC := usecase.NewGetProjectUseCase(repo)

	created, err := createUC.Execute(context.Background(), usecase.CreateProjectInput{
		Name: "StudyPanel",
		Type: entity.ProjectTypePessoal,
	})
	if err != nil {
		t.Fatalf("erro ao criar projeto: %v", err)
	}

	result, err := getUC.Execute(context.Background(), created.ID.Hex())
	if err != nil {
		t.Fatalf("esperava sem erro, recebeu: %v", err)
	}
	if result.ID != created.ID {
		t.Errorf("esperava id=%q, recebeu=%q", created.ID.Hex(), result.ID.Hex())
	}
}

func TestGetProject_NotFound(t *testing.T) {
	repo := newMockProjectRepo()
	getUC := usecase.NewGetProjectUseCase(repo)

	_, err := getUC.Execute(context.Background(), primitive.NewObjectID().Hex())
	if err == nil {
		t.Fatal("esperava erro para projeto não encontrado")
	}
}

func TestUpdateProject_Success(t *testing.T) {
	repo := newMockProjectRepo()
	createUC := usecase.NewCreateProjectUseCase(repo)
	updateUC := usecase.NewUpdateProjectUseCase(repo)

	created, err := createUC.Execute(context.Background(), usecase.CreateProjectInput{
		Name: "StudyPanel",
		Type: entity.ProjectTypePessoal,
	})
	if err != nil {
		t.Fatalf("erro ao criar projeto: %v", err)
	}

	updateInput := usecase.UpdateProjectInput{
		Name:        "StudyPanel Atualizado",
		Description: "Descrição atualizada",
		Tags:        []string{"go", "angular", "mongodb"},
		Active:      true,
		Order:       2,
	}

	result, err := updateUC.Execute(context.Background(), created.ID.Hex(), updateInput)
	if err != nil {
		t.Fatalf("esperava sem erro, recebeu: %v", err)
	}
	if result.Name != updateInput.Name {
		t.Errorf("esperava name=%q, recebeu=%q", updateInput.Name, result.Name)
	}
}

func TestDeleteProject_Success(t *testing.T) {
	repo := newMockProjectRepo()
	createUC := usecase.NewCreateProjectUseCase(repo)
	deleteUC := usecase.NewDeleteProjectUseCase(repo)

	created, err := createUC.Execute(context.Background(), usecase.CreateProjectInput{
		Name: "StudyPanel",
		Type: entity.ProjectTypePessoal,
	})
	if err != nil {
		t.Fatalf("erro ao criar projeto: %v", err)
	}

	if err := deleteUC.Execute(context.Background(), created.ID.Hex()); err != nil {
		t.Fatalf("esperava sem erro, recebeu: %v", err)
	}

	if _, exists := repo.items[created.ID.Hex()]; exists {
		t.Error("projeto deveria ter sido removido do repositório")
	}
}
