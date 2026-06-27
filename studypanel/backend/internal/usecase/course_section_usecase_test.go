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

type mockCourseSectionRepo struct {
	items map[string]*entity.CourseSection
	err   error
}

func newMockCourseSectionRepo() *mockCourseSectionRepo {
	return &mockCourseSectionRepo{items: make(map[string]*entity.CourseSection)}
}

func (m *mockCourseSectionRepo) Create(_ context.Context, s *entity.CourseSection) (*entity.CourseSection, error) {
	if m.err != nil {
		return nil, m.err
	}
	s.ID = primitive.NewObjectID()
	m.items[s.ID.Hex()] = s
	return s, nil
}

func (m *mockCourseSectionRepo) GetByID(_ context.Context, id string) (*entity.CourseSection, error) {
	if m.err != nil {
		return nil, m.err
	}
	s, ok := m.items[id]
	if !ok {
		return nil, errors.New("seção não encontrada")
	}
	return s, nil
}

func (m *mockCourseSectionRepo) GetBySlug(_ context.Context, slug string) (*entity.CourseSection, error) {
	if m.err != nil {
		return nil, m.err
	}
	for _, s := range m.items {
		if s.Slug == slug {
			return s, nil
		}
	}
	return nil, errors.New("seção não encontrada")
}

func (m *mockCourseSectionRepo) List(_ context.Context) ([]*entity.CourseSection, error) {
	if m.err != nil {
		return nil, m.err
	}
	var result []*entity.CourseSection
	for _, s := range m.items {
		result = append(result, s)
	}
	return result, nil
}

func (m *mockCourseSectionRepo) Update(_ context.Context, id string, s *entity.CourseSection) (*entity.CourseSection, error) {
	if m.err != nil {
		return nil, m.err
	}
	m.items[id] = s
	return s, nil
}

func (m *mockCourseSectionRepo) Delete(_ context.Context, id string) error {
	if m.err != nil {
		return m.err
	}
	delete(m.items, id)
	return nil
}

// ── tests ─────────────────────────────────────────────────────────────────────

func TestCreateCourseSection_Success(t *testing.T) {
	repo := newMockCourseSectionRepo()
	uc := usecase.NewCreateCourseSectionUseCase(repo)

	input := usecase.CreateCourseSectionInput{
		Slug:        "cloud",
		Name:        "Cloud Computing",
		Description: "AWS, GCP, Azure",
		BannerColor: "#FF5733",
		IconClass:   "fa-cloud",
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
	if result.Slug != input.Slug {
		t.Errorf("esperava slug=%q, recebeu=%q", input.Slug, result.Slug)
	}
	if result.ID.IsZero() {
		t.Error("ID não deve ser zero")
	}
	if result.CreatedAt.IsZero() {
		t.Error("CreatedAt não deve ser zero")
	}
}

func TestCreateCourseSection_MissingName(t *testing.T) {
	repo := newMockCourseSectionRepo()
	uc := usecase.NewCreateCourseSectionUseCase(repo)

	input := usecase.CreateCourseSectionInput{
		Slug: "cloud",
	}

	_, err := uc.Execute(context.Background(), input)
	if err == nil {
		t.Fatal("esperava erro para name vazio")
	}
}

func TestCreateCourseSection_MissingSlug(t *testing.T) {
	repo := newMockCourseSectionRepo()
	uc := usecase.NewCreateCourseSectionUseCase(repo)

	input := usecase.CreateCourseSectionInput{
		Name: "Cloud Computing",
	}

	_, err := uc.Execute(context.Background(), input)
	if err == nil {
		t.Fatal("esperava erro para slug vazio")
	}
}

func TestCreateCourseSection_RepositoryError(t *testing.T) {
	repo := newMockCourseSectionRepo()
	repo.err = errors.New("erro de banco de dados")
	uc := usecase.NewCreateCourseSectionUseCase(repo)

	input := usecase.CreateCourseSectionInput{
		Slug: "cloud",
		Name: "Cloud Computing",
	}

	_, err := uc.Execute(context.Background(), input)
	if err == nil {
		t.Fatal("esperava erro do repositório")
	}
}

func TestListCourseSections_ReturnsAll(t *testing.T) {
	repo := newMockCourseSectionRepo()
	createUC := usecase.NewCreateCourseSectionUseCase(repo)
	listUC := usecase.NewListCourseSectionsUseCase(repo)

	sections := []usecase.CreateCourseSectionInput{
		{Slug: "cloud", Name: "Cloud"},
		{Slug: "frontend", Name: "Frontend"},
		{Slug: "backend", Name: "Backend"},
	}
	for _, s := range sections {
		if _, err := createUC.Execute(context.Background(), s); err != nil {
			t.Fatalf("erro ao criar seção: %v", err)
		}
	}

	result, err := listUC.Execute(context.Background())
	if err != nil {
		t.Fatalf("esperava sem erro, recebeu: %v", err)
	}
	if len(result) != len(sections) {
		t.Errorf("esperava %d seções, recebeu %d", len(sections), len(result))
	}
}

func TestGetCourseSection_Success(t *testing.T) {
	repo := newMockCourseSectionRepo()
	createUC := usecase.NewCreateCourseSectionUseCase(repo)
	getUC := usecase.NewGetCourseSectionUseCase(repo)

	created, err := createUC.Execute(context.Background(), usecase.CreateCourseSectionInput{
		Slug: "cloud",
		Name: "Cloud Computing",
	})
	if err != nil {
		t.Fatalf("erro ao criar seção: %v", err)
	}

	result, err := getUC.Execute(context.Background(), created.ID.Hex())
	if err != nil {
		t.Fatalf("esperava sem erro, recebeu: %v", err)
	}
	if result.ID != created.ID {
		t.Errorf("esperava id=%q, recebeu=%q", created.ID.Hex(), result.ID.Hex())
	}
}

func TestGetCourseSection_NotFound(t *testing.T) {
	repo := newMockCourseSectionRepo()
	getUC := usecase.NewGetCourseSectionUseCase(repo)

	_, err := getUC.Execute(context.Background(), primitive.NewObjectID().Hex())
	if err == nil {
		t.Fatal("esperava erro para seção não encontrada")
	}
}

func TestUpdateCourseSection_Success(t *testing.T) {
	repo := newMockCourseSectionRepo()
	createUC := usecase.NewCreateCourseSectionUseCase(repo)
	updateUC := usecase.NewUpdateCourseSectionUseCase(repo)

	created, err := createUC.Execute(context.Background(), usecase.CreateCourseSectionInput{
		Slug: "cloud",
		Name: "Cloud Computing",
	})
	if err != nil {
		t.Fatalf("erro ao criar seção: %v", err)
	}

	updateInput := usecase.UpdateCourseSectionInput{
		Name:        "Cloud Computing Atualizado",
		Description: "Descrição atualizada",
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

func TestDeleteCourseSection_Success(t *testing.T) {
	repo := newMockCourseSectionRepo()
	createUC := usecase.NewCreateCourseSectionUseCase(repo)
	deleteUC := usecase.NewDeleteCourseSectionUseCase(repo)

	created, err := createUC.Execute(context.Background(), usecase.CreateCourseSectionInput{
		Slug: "cloud",
		Name: "Cloud Computing",
	})
	if err != nil {
		t.Fatalf("erro ao criar seção: %v", err)
	}

	if err := deleteUC.Execute(context.Background(), created.ID.Hex()); err != nil {
		t.Fatalf("esperava sem erro, recebeu: %v", err)
	}

	if _, exists := repo.items[created.ID.Hex()]; exists {
		t.Error("seção deveria ter sido removida do repositório")
	}
}
