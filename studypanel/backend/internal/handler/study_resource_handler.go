package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"studypanel-backend/internal/domain/entity"
	"studypanel-backend/internal/usecase"
	"studypanel-backend/pkg/response"
)

type StudyResourceHandler struct {
	createUC *usecase.CreateStudyResourceUseCase
	listUC   *usecase.ListStudyResourcesUseCase
	updateUC *usecase.UpdateStudyResourceUseCase
	deleteUC *usecase.DeleteStudyResourceUseCase
}

func NewStudyResourceHandler(
	createUC *usecase.CreateStudyResourceUseCase,
	listUC *usecase.ListStudyResourcesUseCase,
	updateUC *usecase.UpdateStudyResourceUseCase,
	deleteUC *usecase.DeleteStudyResourceUseCase,
) *StudyResourceHandler {
	return &StudyResourceHandler{
		createUC: createUC,
		listUC:   listUC,
		updateUC: updateUC,
		deleteUC: deleteUC,
	}
}

func (h *StudyResourceHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/v1/study-items/{itemId}/resources", h.List)
	mux.HandleFunc("POST /api/v1/study-items/{itemId}/resources", h.Create)
	mux.HandleFunc("PUT /api/v1/study-items/{itemId}/resources/{resourceId}", h.Update)
	mux.HandleFunc("DELETE /api/v1/study-items/{itemId}/resources/{resourceId}", h.Delete)
}

func (h *StudyResourceHandler) List(w http.ResponseWriter, r *http.Request) {
	itemID := r.PathValue("itemId")
	if itemID == "" {
		response.Error(w, http.StatusBadRequest, "itemId é obrigatório")
		return
	}

	resources, err := h.listUC.Execute(r.Context(), itemID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	if resources == nil {
		resources = []*entity.StudyResource{}
	}
	response.JSON(w, http.StatusOK, resources)
}

func (h *StudyResourceHandler) Create(w http.ResponseWriter, r *http.Request) {
	itemID := r.PathValue("itemId")
	if itemID == "" {
		response.Error(w, http.StatusBadRequest, "itemId é obrigatório")
		return
	}

	var input usecase.CreateStudyResourceInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "corpo da requisição inválido")
		return
	}

	resource, err := h.createUC.Execute(r.Context(), itemID, input)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	response.JSON(w, http.StatusCreated, resource)
}

func (h *StudyResourceHandler) Update(w http.ResponseWriter, r *http.Request) {
	resourceID := r.PathValue("resourceId")
	if resourceID == "" {
		response.Error(w, http.StatusBadRequest, "resourceId é obrigatório")
		return
	}

	var input usecase.UpdateStudyResourceInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "corpo da requisição inválido")
		return
	}

	resource, err := h.updateUC.Execute(r.Context(), resourceID, input)
	if err != nil {
		if strings.Contains(err.Error(), "não encontrado") {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, resource)
}

func (h *StudyResourceHandler) Delete(w http.ResponseWriter, r *http.Request) {
	resourceID := r.PathValue("resourceId")
	if resourceID == "" {
		response.Error(w, http.StatusBadRequest, "resourceId é obrigatório")
		return
	}

	if err := h.deleteUC.Execute(r.Context(), resourceID); err != nil {
		if strings.Contains(err.Error(), "não encontrado") {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, map[string]string{"message": "recurso excluído com sucesso"})
}
