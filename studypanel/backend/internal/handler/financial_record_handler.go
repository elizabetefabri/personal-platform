package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"studypanel-backend/internal/domain/entity"
	"studypanel-backend/internal/usecase"
	"studypanel-backend/pkg/response"
)

type FinancialRecordHandler struct {
	createUC *usecase.CreateFinancialRecordUseCase
	listUC   *usecase.ListFinancialRecordsUseCase
	getUC    *usecase.GetFinancialRecordUseCase
	updateUC *usecase.UpdateFinancialRecordUseCase
	deleteUC *usecase.DeleteFinancialRecordUseCase
}

func NewFinancialRecordHandler(
	createUC *usecase.CreateFinancialRecordUseCase,
	listUC *usecase.ListFinancialRecordsUseCase,
	getUC *usecase.GetFinancialRecordUseCase,
	updateUC *usecase.UpdateFinancialRecordUseCase,
	deleteUC *usecase.DeleteFinancialRecordUseCase,
) *FinancialRecordHandler {
	return &FinancialRecordHandler{createUC: createUC, listUC: listUC, getUC: getUC, updateUC: updateUC, deleteUC: deleteUC}
}

func (h *FinancialRecordHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/v1/financial-records", h.List)
	mux.HandleFunc("POST /api/v1/financial-records", h.Create)
	mux.HandleFunc("GET /api/v1/financial-records/{id}", h.Get)
	mux.HandleFunc("PUT /api/v1/financial-records/{id}", h.Update)
	mux.HandleFunc("DELETE /api/v1/financial-records/{id}", h.Delete)
}

func (h *FinancialRecordHandler) List(w http.ResponseWriter, r *http.Request) {
	recordType := r.URL.Query().Get("type")
	records, err := h.listUC.Execute(r.Context(), recordType)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	if records == nil {
		records = []*entity.FinancialRecord{}
	}
	response.JSON(w, http.StatusOK, records)
}

func (h *FinancialRecordHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateFinancialRecordInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "corpo da requisição inválido")
		return
	}
	record, err := h.createUC.Execute(r.Context(), input)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	response.JSON(w, http.StatusCreated, record)
}

func (h *FinancialRecordHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	record, err := h.getUC.Execute(r.Context(), id)
	if err != nil {
		if strings.Contains(err.Error(), "não encontrado") {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, record)
}

func (h *FinancialRecordHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var input usecase.UpdateFinancialRecordInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "corpo da requisição inválido")
		return
	}
	record, err := h.updateUC.Execute(r.Context(), id, input)
	if err != nil {
		if strings.Contains(err.Error(), "não encontrado") {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, record)
}

func (h *FinancialRecordHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := h.deleteUC.Execute(r.Context(), id); err != nil {
		if strings.Contains(err.Error(), "não encontrado") {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, map[string]string{"message": "registro excluído com sucesso"})
}
