package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"studypanel-backend/internal/domain/entity"
	"studypanel-backend/internal/usecase"
	"studypanel-backend/pkg/response"
)

type CulinariaRecipeHandler struct {
	createUC *usecase.CreateCulinariaRecipeUseCase
	listUC   *usecase.ListCulinariaRecipesUseCase
	getUC    *usecase.GetCulinariaRecipeUseCase
	updateUC *usecase.UpdateCulinariaRecipeUseCase
	deleteUC *usecase.DeleteCulinariaRecipeUseCase
}

func NewCulinariaRecipeHandler(
	createUC *usecase.CreateCulinariaRecipeUseCase,
	listUC *usecase.ListCulinariaRecipesUseCase,
	getUC *usecase.GetCulinariaRecipeUseCase,
	updateUC *usecase.UpdateCulinariaRecipeUseCase,
	deleteUC *usecase.DeleteCulinariaRecipeUseCase,
) *CulinariaRecipeHandler {
	return &CulinariaRecipeHandler{createUC: createUC, listUC: listUC, getUC: getUC, updateUC: updateUC, deleteUC: deleteUC}
}

func (h *CulinariaRecipeHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/v1/culinaria/recipes", h.List)
	mux.HandleFunc("POST /api/v1/culinaria/recipes", h.Create)
	mux.HandleFunc("GET /api/v1/culinaria/recipes/{id}", h.Get)
	mux.HandleFunc("PUT /api/v1/culinaria/recipes/{id}", h.Update)
	mux.HandleFunc("DELETE /api/v1/culinaria/recipes/{id}", h.Delete)
}

func (h *CulinariaRecipeHandler) List(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")
	recipes, err := h.listUC.Execute(r.Context(), category)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	if recipes == nil {
		recipes = []*entity.CulinariaRecipe{}
	}
	response.JSON(w, http.StatusOK, recipes)
}

func (h *CulinariaRecipeHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateCulinariaRecipeInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "corpo da requisição inválido")
		return
	}
	recipe, err := h.createUC.Execute(r.Context(), input)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	response.JSON(w, http.StatusCreated, recipe)
}

func (h *CulinariaRecipeHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	recipe, err := h.getUC.Execute(r.Context(), id)
	if err != nil {
		if strings.Contains(err.Error(), "não encontrad") {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, recipe)
}

func (h *CulinariaRecipeHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var input usecase.UpdateCulinariaRecipeInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "corpo da requisição inválido")
		return
	}
	recipe, err := h.updateUC.Execute(r.Context(), id, input)
	if err != nil {
		if strings.Contains(err.Error(), "não encontrad") {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, recipe)
}

func (h *CulinariaRecipeHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := h.deleteUC.Execute(r.Context(), id); err != nil {
		if strings.Contains(err.Error(), "não encontrad") {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, map[string]string{"message": "receita excluída com sucesso"})
}
