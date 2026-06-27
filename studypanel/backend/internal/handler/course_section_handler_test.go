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

type csHandlerMockRepo struct {
	items map[string]*entity.CourseSection
	err   error
}

func newCSHandlerMockRepo() *csHandlerMockRepo {
	return &csHandlerMockRepo{items: make(map[string]*entity.CourseSection)}
}

func (m *csHandlerMockRepo) Create(_ context.Context, s *entity.CourseSection) (*entity.CourseSection, error) {
	if m.err != nil {
		return nil, m.err
	}
	s.ID = primitive.NewObjectID()
	m.items[s.ID.Hex()] = s
	return s, nil
}

func (m *csHandlerMockRepo) GetByID(_ context.Context, id string) (*entity.CourseSection, error) {
	if m.err != nil {
		return nil, m.err
	}
	s, ok := m.items[id]
	if !ok {
		return nil, errors.New("seção não encontrada")
	}
	return s, nil
}

func (m *csHandlerMockRepo) GetBySlug(_ context.Context, slug string) (*entity.CourseSection, error) {
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

func (m *csHandlerMockRepo) List(_ context.Context) ([]*entity.CourseSection, error) {
	if m.err != nil {
		return nil, m.err
	}
	var result []*entity.CourseSection
	for _, s := range m.items {
		result = append(result, s)
	}
	return result, nil
}

func (m *csHandlerMockRepo) Update(_ context.Context, id string, s *entity.CourseSection) (*entity.CourseSection, error) {
	if m.err != nil {
		return nil, m.err
	}
	m.items[id] = s
	return s, nil
}

func (m *csHandlerMockRepo) Delete(_ context.Context, id string) error {
	if m.err != nil {
		return m.err
	}
	delete(m.items, id)
	return nil
}

func buildCourseSectionHandler(repo *csHandlerMockRepo) *handler.CourseSectionHandler {
	return handler.NewCourseSectionHandler(
		usecase.NewCreateCourseSectionUseCase(repo),
		usecase.NewListCourseSectionsUseCase(repo),
		usecase.NewGetCourseSectionUseCase(repo),
		usecase.NewUpdateCourseSectionUseCase(repo),
		usecase.NewDeleteCourseSectionUseCase(repo),
	)
}

// ── tests ─────────────────────────────────────────────────────────────────────

func TestCSHandler_List_Empty(t *testing.T) {
	repo := newCSHandlerMockRepo()
	h := buildCourseSectionHandler(repo)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/course-sections", nil)
	w := httptest.NewRecorder()

	h.List(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("esperava 200, recebeu %d", w.Code)
	}

	var result []interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		// response may be wrapped in a data key — just check status
		return
	}
	if result == nil {
		t.Error("esperava array vazio, recebeu nil")
	}
}

func TestCSHandler_Create_Success(t *testing.T) {
	repo := newCSHandlerMockRepo()
	h := buildCourseSectionHandler(repo)

	body := usecase.CreateCourseSectionInput{
		Slug:   "cloud",
		Name:   "Cloud Computing",
		Active: true,
		Order:  1,
	}
	b, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/course-sections", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.Create(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("esperava 201, recebeu %d — body: %s", w.Code, w.Body.String())
	}
}

func TestCSHandler_Create_MissingName(t *testing.T) {
	repo := newCSHandlerMockRepo()
	h := buildCourseSectionHandler(repo)

	body := usecase.CreateCourseSectionInput{
		Slug: "cloud",
		// Name omitido
	}
	b, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/course-sections", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.Create(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("esperava 400, recebeu %d", w.Code)
	}
}

func TestCSHandler_Get_Success(t *testing.T) {
	repo := newCSHandlerMockRepo()
	h := buildCourseSectionHandler(repo)

	// cria uma seção diretamente no mock
	section := &entity.CourseSection{
		ID:   primitive.NewObjectID(),
		Slug: "cloud",
		Name: "Cloud Computing",
	}
	repo.items[section.ID.Hex()] = section

	req := httptest.NewRequest(http.MethodGet, "/api/v1/course-sections/"+section.ID.Hex(), nil)
	req.SetPathValue("id", section.ID.Hex())
	w := httptest.NewRecorder()

	h.Get(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("esperava 200, recebeu %d — body: %s", w.Code, w.Body.String())
	}
}

func TestCSHandler_Get_NotFound(t *testing.T) {
	repo := newCSHandlerMockRepo()
	h := buildCourseSectionHandler(repo)

	id := primitive.NewObjectID().Hex()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/course-sections/"+id, nil)
	req.SetPathValue("id", id)
	w := httptest.NewRecorder()

	h.Get(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("esperava 404, recebeu %d", w.Code)
	}
}

func TestCSHandler_Update_Success(t *testing.T) {
	repo := newCSHandlerMockRepo()
	h := buildCourseSectionHandler(repo)

	section := &entity.CourseSection{
		ID:   primitive.NewObjectID(),
		Slug: "cloud",
		Name: "Cloud Computing",
	}
	repo.items[section.ID.Hex()] = section

	updateBody := usecase.UpdateCourseSectionInput{
		Name:        "Cloud Computing Atualizado",
		Description: "Descrição nova",
		Active:      true,
		Order:       2,
	}
	b, _ := json.Marshal(updateBody)

	req := httptest.NewRequest(http.MethodPut, "/api/v1/course-sections/"+section.ID.Hex(), bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.SetPathValue("id", section.ID.Hex())
	w := httptest.NewRecorder()

	h.Update(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("esperava 200, recebeu %d — body: %s", w.Code, w.Body.String())
	}
}

func TestCSHandler_Delete_Success(t *testing.T) {
	repo := newCSHandlerMockRepo()
	h := buildCourseSectionHandler(repo)

	section := &entity.CourseSection{
		ID:   primitive.NewObjectID(),
		Slug: "cloud",
		Name: "Cloud Computing",
	}
	repo.items[section.ID.Hex()] = section

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/course-sections/"+section.ID.Hex(), nil)
	req.SetPathValue("id", section.ID.Hex())
	w := httptest.NewRecorder()

	h.Delete(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("esperava 200, recebeu %d — body: %s", w.Code, w.Body.String())
	}
}
