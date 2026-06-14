package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"studypanel-backend/internal/domain/entity"
	"studypanel-backend/internal/usecase"
	"studypanel-backend/pkg/response"
)

type QuizQuestionHandler struct {
	createUC *usecase.CreateQuizQuestionUseCase
	listUC   *usecase.ListQuizQuestionsUseCase
	updateUC *usecase.UpdateQuizQuestionUseCase
	deleteUC *usecase.DeleteQuizQuestionUseCase
}

func NewQuizQuestionHandler(
	createUC *usecase.CreateQuizQuestionUseCase,
	listUC *usecase.ListQuizQuestionsUseCase,
	updateUC *usecase.UpdateQuizQuestionUseCase,
	deleteUC *usecase.DeleteQuizQuestionUseCase,
) *QuizQuestionHandler {
	return &QuizQuestionHandler{
		createUC: createUC,
		listUC:   listUC,
		updateUC: updateUC,
		deleteUC: deleteUC,
	}
}

func (h *QuizQuestionHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/v1/quiz-questions", h.List)
	mux.HandleFunc("POST /api/v1/quiz-questions", h.Create)
	mux.HandleFunc("PUT /api/v1/quiz-questions/{id}", h.Update)
	mux.HandleFunc("DELETE /api/v1/quiz-questions/{id}", h.Delete)
}

func (h *QuizQuestionHandler) List(w http.ResponseWriter, r *http.Request) {
	section := r.URL.Query().Get("section")
	topic := r.URL.Query().Get("topic")

	questions, err := h.listUC.Execute(r.Context(), section, topic)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	if questions == nil {
		questions = []*entity.QuizQuestion{}
	}
	response.JSON(w, http.StatusOK, questions)
}

func (h *QuizQuestionHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateQuizQuestionInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "corpo da requisição inválido")
		return
	}

	question, err := h.createUC.Execute(r.Context(), input)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	response.JSON(w, http.StatusCreated, question)
}

func (h *QuizQuestionHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		response.Error(w, http.StatusBadRequest, "id é obrigatório")
		return
	}

	var input usecase.UpdateQuizQuestionInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "corpo da requisição inválido")
		return
	}

	question, err := h.updateUC.Execute(r.Context(), id, input)
	if err != nil {
		if strings.Contains(err.Error(), "não encontrada") {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, question)
}

func (h *QuizQuestionHandler) Delete(w http.ResponseWriter, r *http.Request) {
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

	response.JSON(w, http.StatusOK, map[string]string{"message": "questão excluída com sucesso"})
}
