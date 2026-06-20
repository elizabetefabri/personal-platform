package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"studypanel-backend/internal/domain/entity"
	"studypanel-backend/internal/usecase"
	"studypanel-backend/pkg/response"
)

type CourseSectionHandler struct {
	createUC *usecase.CreateCourseSectionUseCase
	listUC   *usecase.ListCourseSectionsUseCase
	getUC    *usecase.GetCourseSectionUseCase
	updateUC *usecase.UpdateCourseSectionUseCase
	deleteUC *usecase.DeleteCourseSectionUseCase
}

func NewCourseSectionHandler(
	createUC *usecase.CreateCourseSectionUseCase,
	listUC *usecase.ListCourseSectionsUseCase,
	getUC *usecase.GetCourseSectionUseCase,
	updateUC *usecase.UpdateCourseSectionUseCase,
	deleteUC *usecase.DeleteCourseSectionUseCase,
) *CourseSectionHandler {
	return &CourseSectionHandler{createUC: createUC, listUC: listUC, getUC: getUC, updateUC: updateUC, deleteUC: deleteUC}
}

func (h *CourseSectionHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/v1/course-sections", h.List)
	mux.HandleFunc("POST /api/v1/course-sections", h.Create)
	mux.HandleFunc("GET /api/v1/course-sections/{id}", h.Get)
	mux.HandleFunc("PUT /api/v1/course-sections/{id}", h.Update)
	mux.HandleFunc("DELETE /api/v1/course-sections/{id}", h.Delete)
}

func (h *CourseSectionHandler) List(w http.ResponseWriter, r *http.Request) {
	items, err := h.listUC.Execute(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	if items == nil {
		items = []*entity.CourseSection{}
	}
	response.JSON(w, http.StatusOK, items)
}

func (h *CourseSectionHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateCourseSectionInput
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

func (h *CourseSectionHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		response.Error(w, http.StatusBadRequest, "id é obrigatório")
		return
	}
	item, err := h.getUC.Execute(r.Context(), id)
	if err != nil {
		if strings.Contains(err.Error(), "não encontrada") {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, item)
}

func (h *CourseSectionHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		response.Error(w, http.StatusBadRequest, "id é obrigatório")
		return
	}
	var input usecase.UpdateCourseSectionInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "corpo da requisição inválido")
		return
	}
	item, err := h.updateUC.Execute(r.Context(), id, input)
	if err != nil {
		if strings.Contains(err.Error(), "não encontrada") {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, item)
}

func (h *CourseSectionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		response.Error(w, http.StatusBadRequest, "id é obrigatório")
		return
	}
	if err := h.deleteUC.Execute(r.Context(), id); err != nil {
		if strings.Contains(err.Error(), "não encontrada") {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, map[string]string{"message": "seção excluída com sucesso"})
}
