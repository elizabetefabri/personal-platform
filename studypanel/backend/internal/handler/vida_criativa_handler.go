package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"studypanel-backend/internal/domain/entity"
	"studypanel-backend/internal/usecase"
	"studypanel-backend/pkg/response"
)

type VidaCriativaHandler struct {
	createUC *usecase.CreateVidaCriativaUseCase
	listUC   *usecase.ListVidaCriativaUseCase
	getUC    *usecase.GetVidaCriativaUseCase
	updateUC *usecase.UpdateVidaCriativaUseCase
	deleteUC *usecase.DeleteVidaCriativaUseCase
}

func NewVidaCriativaHandler(
	createUC *usecase.CreateVidaCriativaUseCase,
	listUC *usecase.ListVidaCriativaUseCase,
	getUC *usecase.GetVidaCriativaUseCase,
	updateUC *usecase.UpdateVidaCriativaUseCase,
	deleteUC *usecase.DeleteVidaCriativaUseCase,
) *VidaCriativaHandler {
	return &VidaCriativaHandler{createUC: createUC, listUC: listUC, getUC: getUC, updateUC: updateUC, deleteUC: deleteUC}
}

func (h *VidaCriativaHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/v1/vida-criativa", h.List)
	mux.HandleFunc("POST /api/v1/vida-criativa", h.Create)
	mux.HandleFunc("GET /api/v1/vida-criativa/{id}", h.Get)
	mux.HandleFunc("PUT /api/v1/vida-criativa/{id}", h.Update)
	mux.HandleFunc("DELETE /api/v1/vida-criativa/{id}", h.Delete)
}

func (h *VidaCriativaHandler) List(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")
	items, err := h.listUC.Execute(r.Context(), category)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	if items == nil {
		items = []*entity.VidaCriativaItem{}
	}
	response.JSON(w, http.StatusOK, items)
}

func (h *VidaCriativaHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateVidaCriativaInput
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

func (h *VidaCriativaHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
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

func (h *VidaCriativaHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var input usecase.UpdateVidaCriativaInput
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

func (h *VidaCriativaHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
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
