package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"studypanel-backend/internal/domain/entity"
	"studypanel-backend/internal/handler"
	"studypanel-backend/internal/usecase"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ── mock ──────────────────────────────────────────────────────────────────────

type projectHandlerMockRepo struct {
	items map[string]*entity.Project
	err   error
}

func newProjectHandlerMockRepo() *projectHandlerMockRepo {
	return &projectHandlerMockRepo{items: make(map[string]*entity.Project)}
}

func (m *projectHandlerMockRepo) Create(_ context.Context, p *entity.Project) (*entity.Project, error) {
	if m.err != nil {
		return nil, m.err
	}
	p.ID = primitive.NewObjectID()
	m.items[p.ID.Hex()] = p
	return p, nil
}

func (m *projectHandlerMockRepo) GetByID(_ context.Context, id string) (*entity.Project, error) {
	if m.err != nil {
		return nil, m.err
	}
	p, ok := m.items[id]
	if !ok {
		return nil, errors.New("projeto não encontrado")
	}
	return p, nil
}

func (m *projectHandlerMockRepo) List(_ context.Context, projectType string) ([]*entity.Project, error) {
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

func (m *projectHandlerMockRepo) Update(_ context.Context, id string, p *entity.Project) (*entity.Project, error) {
	if m.err != nil {
		return nil, m.err
	}
	m.items[id] = p
	return p, nil
}

func (m *projectHandlerMockRepo) Delete(_ context.Context, id string) error {
	if m.err != nil {
		return m.err
	}
	delete(m.items, id)
	return nil
}

func buildProjectHandler(repo *projectHandlerMockRepo) *handler.ProjectHandler {
	return handler.NewProjectHandler(
		usecase.NewCreateProjectUseCase(repo),
		usecase.NewListProjectsUseCase(repo),
		usecase.NewGetProjectUseCase(repo),
		usecase.NewUpdateProjectUseCase(repo),
		usecase.NewDeleteProjectUseCase(repo),
	)
}

// ── tests ─────────────────────────────────────────────────────────────────────

func TestProjectHandler_List_Empty(t *testing.T) {
	repo := newProjectHandlerMockRepo()
	h := buildProjectHandler(repo)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/projects", nil)
	w := httptest.NewRecorder()

	h.List(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("esperava 200, recebeu %d", w.Code)
	}
}

func TestProjectHandler_List_ByType(t *testing.T) {
	repo := newProjectHandlerMockRepo()
	h := buildProjectHandler(repo)

	// Pré-popula com projetos de tipos diferentes
	repo.items["1"] = &entity.Project{ID: primitive.NewObjectID(), Name: "Pessoal 1", Type: entity.ProjectTypePessoal}
	repo.items["2"] = &entity.Project{ID: primitive.NewObjectID(), Name: "Pessoal 2", Type: entity.ProjectTypePessoal}
	repo.items["3"] = &entity.Project{ID: primitive.NewObjectID(), Name: "Profissional 1", Type: entity.ProjectTypeProfissional}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/projects?type=pessoal", nil)
	w := httptest.NewRecorder()

	h.List(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("esperava 200, recebeu %d", w.Code)
	}
}

func TestProjectHandler_Create_Success(t *testing.T) {
	repo := newProjectHandlerMockRepo()
	h := buildProjectHandler(repo)

	body := usecase.CreateProjectInput{
		Name:   "StudyPanel",
		Type:   entity.ProjectTypePessoal,
		Slug:   "studypanel",
		Active: true,
		Order:  1,
	}
	b, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/projects", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.Create(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("esperava 201, recebeu %d — body: %s", w.Code, w.Body.String())
	}
}

func TestProjectHandler_Create_MissingName(t *testing.T) {
	repo := newProjectHandlerMockRepo()
	h := buildProjectHandler(repo)

	body := usecase.CreateProjectInput{
		Type: entity.ProjectTypePessoal,
		// Name omitido
	}
	b, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/projects", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.Create(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("esperava 400, recebeu %d", w.Code)
	}
}

func TestProjectHandler_Create_InvalidType(t *testing.T) {
	repo := newProjectHandlerMockRepo()
	h := buildProjectHandler(repo)

	body := usecase.CreateProjectInput{
		Name: "StudyPanel",
		Type: entity.ProjectType("tipo-invalido"),
	}
	b, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/projects", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.Create(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("esperava 400, recebeu %d", w.Code)
	}
}

func TestProjectHandler_Get_Success(t *testing.T) {
	repo := newProjectHandlerMockRepo()
	h := buildProjectHandler(repo)

	project := &entity.Project{
		ID:   primitive.NewObjectID(),
		Name: "StudyPanel",
		Type: entity.ProjectTypePessoal,
	}
	repo.items[project.ID.Hex()] = project

	req := httptest.NewRequest(http.MethodGet, "/api/v1/projects/"+project.ID.Hex(), nil)
	req.SetPathValue("id", project.ID.Hex())
	w := httptest.NewRecorder()

	h.Get(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("esperava 200, recebeu %d — body: %s", w.Code, w.Body.String())
	}
}

func TestProjectHandler_Get_NotFound(t *testing.T) {
	repo := newProjectHandlerMockRepo()
	h := buildProjectHandler(repo)

	id := primitive.NewObjectID().Hex()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/projects/"+id, nil)
	req.SetPathValue("id", id)
	w := httptest.NewRecorder()

	h.Get(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("esperava 404, recebeu %d", w.Code)
	}
}

func TestProjectHandler_Update_Success(t *testing.T) {
	repo := newProjectHandlerMockRepo()
	h := buildProjectHandler(repo)

	project := &entity.Project{
		ID:   primitive.NewObjectID(),
		Name: "StudyPanel",
		Type: entity.ProjectTypePessoal,
	}
	repo.items[project.ID.Hex()] = project

	updateBody := usecase.UpdateProjectInput{
		Name:        "StudyPanel Atualizado",
		Description: "Descrição nova",
		Tags:        []string{"go", "angular"},
		Active:      true,
		Order:       2,
	}
	b, _ := json.Marshal(updateBody)

	req := httptest.NewRequest(http.MethodPut, "/api/v1/projects/"+project.ID.Hex(), bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.SetPathValue("id", project.ID.Hex())
	w := httptest.NewRecorder()

	h.Update(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("esperava 200, recebeu %d — body: %s", w.Code, w.Body.String())
	}
}

func TestProjectHandler_Delete_Success(t *testing.T) {
	repo := newProjectHandlerMockRepo()
	h := buildProjectHandler(repo)

	project := &entity.Project{
		ID:   primitive.NewObjectID(),
		Name: "StudyPanel",
		Type: entity.ProjectTypePessoal,
	}
	repo.items[project.ID.Hex()] = project

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/projects/"+project.ID.Hex(), nil)
	req.SetPathValue("id", project.ID.Hex())
	w := httptest.NewRecorder()

	h.Delete(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("esperava 200, recebeu %d — body: %s", w.Code, w.Body.String())
	}
}
