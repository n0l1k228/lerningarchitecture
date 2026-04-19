package product_handler

import (
	"encoding/json"
	"myavito/internal/dto"
	"net/http"
)

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req dto.ProductDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Error("failed to decode request", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.validator.Struct(req); err != nil {
		h.log.Error("failed to validate request", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.srv.ServiceCreateProduct(ctx, req.AuthorID, req.Price, req.Title, req.Description, req.Category); err != nil {
		h.log.Error("failed to product request", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(req); err != nil {
		h.log.Error("failed to product request", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *ProductHandler) SearchProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req dto.ProductSearch
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Error("failed to Decode request", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	items, err := h.srv.ServiceSearchProduct(ctx, req.Title)
	if err != nil {
		h.log.Error("failed to Decode request", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(items); err != nil {
		h.log.Error("failed to product request", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
