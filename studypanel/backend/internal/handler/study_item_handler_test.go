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
	"studypanel-backend/internal/domain/repository"
	"studypanel-backend/internal/handler"
	"studypanel-backend/internal/usecase"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// handlerMockRepo implementa repository.StudyItemRepository para testes de handler
type handlerMockRepo struct {
	items map[string]*entity.StudyItem
	err   error
}

func newHandlerMockRepo() *handlerMockRepo {
	return &handlerMockRepo{items: make(map[string]*entity.StudyItem)}
}

func (m *handlerMockRepo) Create(_ context.Context, item *entity.StudyItem) (*entity.StudyItem, error) {
	if m.err != nil {
		return nil, m.err
	}
	item.ID = primitive.NewObjectID()
	m.items[item.ID.Hex()] = item
	return item, nil
}

func (m *handlerMockRepo) GetByID(_ context.Context, id string) (*entity.StudyItem, error) {
	if m.err != nil {
		return nil, m.err
	}
	item, ok := m.items[id]
	if !ok {
		return nil, errors.New("item não encontrado")
	}
	return item, nil
}

func (m *handlerMockRepo) List(_ context.Context, _ repository.Filter) ([]*entity.StudyItem, error) {
	if m.err != nil {
		return nil, m.err
	}
	var result []*entity.StudyItem
	for _, item := range m.items {
		result = append(result, item)
	}
	return result, nil
}

func (m *handlerMockRepo) Update(_ context.Context, id string, item *entity.StudyItem) (*entity.StudyItem, error) {
	if m.err != nil {
		return nil, m.err
	}
	m.items[id] = item
	return item, nil
}

func (m *handlerMockRepo) Delete(_ context.Context, id string) error {
	if m.err != nil {
		return m.err
	}
	delete(m.items, id)
	return nil
}

func buildHandler(repo *handlerMockRepo) *handler.StudyItemHandler {
	return handler.NewStudyItemHandler(
		usecase.NewCreateStudyItemUseCase(repo),
		usecase.NewGetStudyItemUseCase(repo),
		usecase.NewListStudyItemsUseCase(repo),
		usecase.NewUpdateStudyItemUseCase(repo),
		usecase.NewDeleteStudyItemUseCase(repo),
	)
}

func TestHandler_List_Empty(t *testing.T) {
	repo := newHandlerMockRepo()
	h := buildHandler(repo)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/study-items", nil)
	w := httptest.NewRecorder()

	h.List(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("esperava 200, recebeu %d", w.Code)
	}
}

func TestHandler_Create_Success(t *testing.T) {
	repo := newHandlerMockRepo()
	h := buildHandler(repo)

	body := usecase.CreateStudyItemInput{
		Section:    "cloud",
		Topic:      "aws",
		CourseName: "AWS SAA",
		Status:     entity.StatusNotStarted,
	}
	b, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/study-items", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.Create(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("esperava 201, recebeu %d — body: %s", w.Code, w.Body.String())
	}
}

func TestHandler_Create_InvalidBody(t *testing.T) {
	repo := newHandlerMockRepo()
	h := buildHandler(repo)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/study-items", bytes.NewReader([]byte("invalid json")))
	w := httptest.NewRecorder()

	h.Create(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("esperava 400, recebeu %d", w.Code)
	}
}

func TestHandler_Create_ValidationError(t *testing.T) {
	repo := newHandlerMockRepo()
	h := buildHandler(repo)

	body := usecase.CreateStudyItemInput{
		Status: entity.StatusNotStarted,
	}
	b, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/study-items", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.Create(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("esperava 400, recebeu %d", w.Code)
	}
}

func TestHandler_Get_EmptyID(t *testing.T) {
	repo := newHandlerMockRepo()
	h := buildHandler(repo)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/study-items/", nil)
	w := httptest.NewRecorder()

	h.Get(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("esperava 400, recebeu %d", w.Code)
	}
}

func TestHandler_Get_NotFound(t *testing.T) {
	repo := newHandlerMockRepo()
	h := buildHandler(repo)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/study-items/"+primitive.NewObjectID().Hex(), nil)
	req.SetPathValue("id", primitive.NewObjectID().Hex())
	w := httptest.NewRecorder()

	h.Get(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("esperava 404, recebeu %d", w.Code)
	}
}

func TestHandler_Delete_EmptyID(t *testing.T) {
	repo := newHandlerMockRepo()
	h := buildHandler(repo)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/study-items/", nil)
	w := httptest.NewRecorder()

	h.Delete(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("esperava 400, recebeu %d", w.Code)
	}
}

func TestHandler_Update_EmptyID(t *testing.T) {
	repo := newHandlerMockRepo()
	h := buildHandler(repo)

	body := usecase.UpdateStudyItemInput{
		CourseName: "Test",
		Status:     entity.StatusCompleted,
	}
	b, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPut, "/api/v1/study-items/", bytes.NewReader(b))
	w := httptest.NewRecorder()

	h.Update(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("esperava 400, recebeu %d", w.Code)
	}
}

func TestHandler_List_RepositoryError(t *testing.T) {
	repo := newHandlerMockRepo()
	repo.err = errors.New("erro de banco")
	h := buildHandler(repo)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/study-items", nil)
	w := httptest.NewRecorder()

	h.List(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("esperava 500, recebeu %d", w.Code)
	}
}
