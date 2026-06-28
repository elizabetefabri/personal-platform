package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"studypanel-backend/internal/domain/entity"
	"studypanel-backend/internal/usecase"
	"studypanel-backend/pkg/response"
)

type CulinaryCategoryHandler struct {
	createUC *usecase.CreateCulinaryCategoryUseCase
	listUC   *usecase.ListCulinaryCategoriesUseCase
	getUC    *usecase.GetCulinaryCategoryUseCase
	updateUC *usecase.UpdateCulinaryCategoryUseCase
	deleteUC *usecase.DeleteCulinaryCategoryUseCase
}

func NewCulinaryCategoryHandler(
	createUC *usecase.CreateCulinaryCategoryUseCase,
	listUC *usecase.ListCulinaryCategoriesUseCase,
	getUC *usecase.GetCulinaryCategoryUseCase,
	updateUC *usecase.UpdateCulinaryCategoryUseCase,
	deleteUC *usecase.DeleteCulinaryCategoryUseCase,
) *CulinaryCategoryHandler {
	return &CulinaryCategoryHandler{createUC: createUC, listUC: listUC, getUC: getUC, updateUC: updateUC, deleteUC: deleteUC}
}

func (h *CulinaryCategoryHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/v1/culinary/categories", h.List)
	mux.HandleFunc("POST /api/v1/culinary/categories", h.Create)
	mux.HandleFunc("GET /api/v1/culinary/categories/{id}", h.Get)
	mux.HandleFunc("PUT /api/v1/culinary/categories/{id}", h.Update)
	mux.HandleFunc("DELETE /api/v1/culinary/categories/{id}", h.Delete)
}

func (h *CulinaryCategoryHandler) List(w http.ResponseWriter, r *http.Request) {
	items, err := h.listUC.Execute(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	if items == nil {
		items = []*entity.CulinaryCategory{}
	}
	response.JSON(w, http.StatusOK, items)
}

func (h *CulinaryCategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateCulinaryCategoryInput
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

func (h *CulinaryCategoryHandler) Get(w http.ResponseWriter, r *http.Request) {
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

func (h *CulinaryCategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		response.Error(w, http.StatusBadRequest, "id é obrigatório")
		return
	}
	var input usecase.UpdateCulinaryCategoryInput
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

func (h *CulinaryCategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
	response.JSON(w, http.StatusOK, map[string]string{"message": "categoria excluída com sucesso"})
}
