package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"studypanel-backend/internal/domain/entity"
	"studypanel-backend/internal/usecase"
	"studypanel-backend/pkg/response"
)

type ProjectHandler struct {
	createUC *usecase.CreateProjectUseCase
	listUC   *usecase.ListProjectsUseCase
	getUC    *usecase.GetProjectUseCase
	updateUC *usecase.UpdateProjectUseCase
	deleteUC *usecase.DeleteProjectUseCase
}

func NewProjectHandler(
	createUC *usecase.CreateProjectUseCase,
	listUC *usecase.ListProjectsUseCase,
	getUC *usecase.GetProjectUseCase,
	updateUC *usecase.UpdateProjectUseCase,
	deleteUC *usecase.DeleteProjectUseCase,
) *ProjectHandler {
	return &ProjectHandler{createUC: createUC, listUC: listUC, getUC: getUC, updateUC: updateUC, deleteUC: deleteUC}
}

func (h *ProjectHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/v1/projects", h.List)
	mux.HandleFunc("POST /api/v1/projects", h.Create)
	mux.HandleFunc("GET /api/v1/projects/{id}", h.Get)
	mux.HandleFunc("PUT /api/v1/projects/{id}", h.Update)
	mux.HandleFunc("DELETE /api/v1/projects/{id}", h.Delete)
}

func (h *ProjectHandler) List(w http.ResponseWriter, r *http.Request) {
	projectType := r.URL.Query().Get("type")
	items, err := h.listUC.Execute(r.Context(), projectType)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	if items == nil {
		items = []*entity.Project{}
	}
	response.JSON(w, http.StatusOK, items)
}

func (h *ProjectHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateProjectInput
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

func (h *ProjectHandler) Get(w http.ResponseWriter, r *http.Request) {
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

func (h *ProjectHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		response.Error(w, http.StatusBadRequest, "id é obrigatório")
		return
	}
	var input usecase.UpdateProjectInput
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

func (h *ProjectHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
	response.JSON(w, http.StatusOK, map[string]string{"message": "projeto excluído com sucesso"})
}
