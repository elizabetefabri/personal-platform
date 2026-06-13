package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"studypanel-backend/internal/domain/entity"
	"studypanel-backend/internal/usecase"
	"studypanel-backend/pkg/response"
)

type StudyItemHandler struct {
	createUC *usecase.CreateStudyItemUseCase
	getUC    *usecase.GetStudyItemUseCase
	listUC   *usecase.ListStudyItemsUseCase
	updateUC *usecase.UpdateStudyItemUseCase
	deleteUC *usecase.DeleteStudyItemUseCase
}

func NewStudyItemHandler(
	createUC *usecase.CreateStudyItemUseCase,
	getUC *usecase.GetStudyItemUseCase,
	listUC *usecase.ListStudyItemsUseCase,
	updateUC *usecase.UpdateStudyItemUseCase,
	deleteUC *usecase.DeleteStudyItemUseCase,
) *StudyItemHandler {
	return &StudyItemHandler{
		createUC: createUC,
		getUC:    getUC,
		listUC:   listUC,
		updateUC: updateUC,
		deleteUC: deleteUC,
	}
}

// RegisterRoutes registra os endpoints em um ServeMux (go 1.22+)
func (h *StudyItemHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/v1/study-items", h.List)
	mux.HandleFunc("POST /api/v1/study-items", h.Create)
	mux.HandleFunc("GET /api/v1/study-items/{id}", h.Get)
	mux.HandleFunc("PUT /api/v1/study-items/{id}", h.Update)
	mux.HandleFunc("DELETE /api/v1/study-items/{id}", h.Delete)
}

func (h *StudyItemHandler) List(w http.ResponseWriter, r *http.Request) {
	section := r.URL.Query().Get("section")
	topic := r.URL.Query().Get("topic")

	items, err := h.listUC.Execute(r.Context(), usecase.ListStudyItemsInput{
		Section: section,
		Topic:   topic,
	})
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	if items == nil {
		items = []*entity.StudyItem{}
	}
	response.JSON(w, http.StatusOK, items)
}

func (h *StudyItemHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateStudyItemInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "corpo da requisição inválido")
		return
	}

	item, err := h.createUC.Execute(r.Context(), input)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	response.JSON(w, http.StatusCreated, item)
}

func (h *StudyItemHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := extractID(r)
	if id == "" {
		response.Error(w, http.StatusBadRequest, "id é obrigatório")
		return
	}

	item, err := h.getUC.Execute(r.Context(), id)
	if err != nil {
		if strings.Contains(err.Error(), "não encontrado") {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, item)
}

func (h *StudyItemHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := extractID(r)
	if id == "" {
		response.Error(w, http.StatusBadRequest, "id é obrigatório")
		return
	}

	var input usecase.UpdateStudyItemInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "corpo da requisição inválido")
		return
	}

	item, err := h.updateUC.Execute(r.Context(), id, input)
	if err != nil {
		if strings.Contains(err.Error(), "não encontrado") {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, item)
}

func (h *StudyItemHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := extractID(r)
	if id == "" {
		response.Error(w, http.StatusBadRequest, "id é obrigatório")
		return
	}

	if err := h.deleteUC.Execute(r.Context(), id); err != nil {
		if strings.Contains(err.Error(), "não encontrado") {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, map[string]string{"message": "item excluído com sucesso"})
}

func extractID(r *http.Request) string {
	return r.PathValue("id")
}
