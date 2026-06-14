package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"studypanel-backend/internal/domain/entity"
	"studypanel-backend/internal/usecase"
	"studypanel-backend/pkg/response"
)

type StudyNoteHandler struct {
	createUC *usecase.CreateStudyNoteUseCase
	listUC   *usecase.ListStudyNotesUseCase
	updateUC *usecase.UpdateStudyNoteUseCase
	deleteUC *usecase.DeleteStudyNoteUseCase
}

func NewStudyNoteHandler(
	createUC *usecase.CreateStudyNoteUseCase,
	listUC *usecase.ListStudyNotesUseCase,
	updateUC *usecase.UpdateStudyNoteUseCase,
	deleteUC *usecase.DeleteStudyNoteUseCase,
) *StudyNoteHandler {
	return &StudyNoteHandler{
		createUC: createUC,
		listUC:   listUC,
		updateUC: updateUC,
		deleteUC: deleteUC,
	}
}

func (h *StudyNoteHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/v1/study-items/{itemId}/notes", h.List)
	mux.HandleFunc("POST /api/v1/study-items/{itemId}/notes", h.Create)
	mux.HandleFunc("PUT /api/v1/study-items/{itemId}/notes/{noteId}", h.Update)
	mux.HandleFunc("DELETE /api/v1/study-items/{itemId}/notes/{noteId}", h.Delete)
}

func (h *StudyNoteHandler) List(w http.ResponseWriter, r *http.Request) {
	itemID := r.PathValue("itemId")
	if itemID == "" {
		response.Error(w, http.StatusBadRequest, "itemId é obrigatório")
		return
	}

	notes, err := h.listUC.Execute(r.Context(), itemID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	if notes == nil {
		notes = []*entity.StudyNote{}
	}
	response.JSON(w, http.StatusOK, notes)
}

func (h *StudyNoteHandler) Create(w http.ResponseWriter, r *http.Request) {
	itemID := r.PathValue("itemId")
	if itemID == "" {
		response.Error(w, http.StatusBadRequest, "itemId é obrigatório")
		return
	}

	var input usecase.CreateStudyNoteInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "corpo da requisição inválido")
		return
	}

	note, err := h.createUC.Execute(r.Context(), itemID, input)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	response.JSON(w, http.StatusCreated, note)
}

func (h *StudyNoteHandler) Update(w http.ResponseWriter, r *http.Request) {
	noteID := r.PathValue("noteId")
	if noteID == "" {
		response.Error(w, http.StatusBadRequest, "noteId é obrigatório")
		return
	}

	var input usecase.UpdateStudyNoteInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "corpo da requisição inválido")
		return
	}

	note, err := h.updateUC.Execute(r.Context(), noteID, input)
	if err != nil {
		if strings.Contains(err.Error(), "não encontrada") {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, note)
}

func (h *StudyNoteHandler) Delete(w http.ResponseWriter, r *http.Request) {
	noteID := r.PathValue("noteId")
	if noteID == "" {
		response.Error(w, http.StatusBadRequest, "noteId é obrigatório")
		return
	}

	if err := h.deleteUC.Execute(r.Context(), noteID); err != nil {
		if strings.Contains(err.Error(), "não encontrada") {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, map[string]string{"message": "nota excluída com sucesso"})
}
