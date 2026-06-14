package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"studypanel-backend/internal/domain/entity"
	"studypanel-backend/internal/usecase"
	"studypanel-backend/pkg/response"
)

type StudySessionHandler struct {
	createUC *usecase.CreateStudySessionUseCase
	listUC   *usecase.ListStudySessionsUseCase
	updateUC *usecase.UpdateStudySessionUseCase
	deleteUC *usecase.DeleteStudySessionUseCase
}

func NewStudySessionHandler(
	createUC *usecase.CreateStudySessionUseCase,
	listUC *usecase.ListStudySessionsUseCase,
	updateUC *usecase.UpdateStudySessionUseCase,
	deleteUC *usecase.DeleteStudySessionUseCase,
) *StudySessionHandler {
	return &StudySessionHandler{
		createUC: createUC,
		listUC:   listUC,
		updateUC: updateUC,
		deleteUC: deleteUC,
	}
}

func (h *StudySessionHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/v1/study-items/{itemId}/sessions", h.List)
	mux.HandleFunc("POST /api/v1/study-items/{itemId}/sessions", h.Create)
	mux.HandleFunc("PUT /api/v1/study-items/{itemId}/sessions/{sessionId}", h.Update)
	mux.HandleFunc("DELETE /api/v1/study-items/{itemId}/sessions/{sessionId}", h.Delete)
}

func (h *StudySessionHandler) List(w http.ResponseWriter, r *http.Request) {
	itemID := r.PathValue("itemId")
	if itemID == "" {
		response.Error(w, http.StatusBadRequest, "itemId é obrigatório")
		return
	}

	sessions, err := h.listUC.Execute(r.Context(), itemID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	if sessions == nil {
		sessions = []*entity.StudySession{}
	}
	response.JSON(w, http.StatusOK, sessions)
}

func (h *StudySessionHandler) Create(w http.ResponseWriter, r *http.Request) {
	itemID := r.PathValue("itemId")
	if itemID == "" {
		response.Error(w, http.StatusBadRequest, "itemId é obrigatório")
		return
	}

	var input usecase.CreateStudySessionInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "corpo da requisição inválido")
		return
	}

	session, err := h.createUC.Execute(r.Context(), itemID, input)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	response.JSON(w, http.StatusCreated, session)
}

func (h *StudySessionHandler) Update(w http.ResponseWriter, r *http.Request) {
	sessionID := r.PathValue("sessionId")
	if sessionID == "" {
		response.Error(w, http.StatusBadRequest, "sessionId é obrigatório")
		return
	}

	var input usecase.UpdateStudySessionInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "corpo da requisição inválido")
		return
	}

	session, err := h.updateUC.Execute(r.Context(), sessionID, input)
	if err != nil {
		if strings.Contains(err.Error(), "não encontrada") {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, session)
}

func (h *StudySessionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	sessionID := r.PathValue("sessionId")
	if sessionID == "" {
		response.Error(w, http.StatusBadRequest, "sessionId é obrigatório")
		return
	}

	if err := h.deleteUC.Execute(r.Context(), sessionID); err != nil {
		if strings.Contains(err.Error(), "não encontrada") {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, map[string]string{"message": "sessão excluída com sucesso"})
}
