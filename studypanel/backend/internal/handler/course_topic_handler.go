package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"studypanel-backend/internal/domain/entity"
	"studypanel-backend/internal/usecase"
	"studypanel-backend/pkg/response"
)

type CourseTopicHandler struct {
	createUC *usecase.CreateCourseTopicUseCase
	listUC   *usecase.ListCourseTopicsUseCase
	getUC    *usecase.GetCourseTopicUseCase
	updateUC *usecase.UpdateCourseTopicUseCase
	deleteUC *usecase.DeleteCourseTopicUseCase
}

func NewCourseTopicHandler(
	createUC *usecase.CreateCourseTopicUseCase,
	listUC *usecase.ListCourseTopicsUseCase,
	getUC *usecase.GetCourseTopicUseCase,
	updateUC *usecase.UpdateCourseTopicUseCase,
	deleteUC *usecase.DeleteCourseTopicUseCase,
) *CourseTopicHandler {
	return &CourseTopicHandler{createUC: createUC, listUC: listUC, getUC: getUC, updateUC: updateUC, deleteUC: deleteUC}
}

func (h *CourseTopicHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/v1/course-topics", h.List)
	mux.HandleFunc("POST /api/v1/course-topics", h.Create)
	mux.HandleFunc("GET /api/v1/course-topics/{id}", h.Get)
	mux.HandleFunc("PUT /api/v1/course-topics/{id}", h.Update)
	mux.HandleFunc("DELETE /api/v1/course-topics/{id}", h.Delete)
}

func (h *CourseTopicHandler) List(w http.ResponseWriter, r *http.Request) {
	sectionSlug := r.URL.Query().Get("sectionSlug")
	items, err := h.listUC.Execute(r.Context(), sectionSlug)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	if items == nil {
		items = []*entity.CourseTopic{}
	}
	response.JSON(w, http.StatusOK, items)
}

func (h *CourseTopicHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateCourseTopicInput
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

func (h *CourseTopicHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
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

func (h *CourseTopicHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		response.Error(w, http.StatusBadRequest, "id é obrigatório")
		return
	}
	var input usecase.UpdateCourseTopicInput
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

func (h *CourseTopicHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
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
	response.JSON(w, http.StatusOK, map[string]string{"message": "tópico excluído com sucesso"})
}
