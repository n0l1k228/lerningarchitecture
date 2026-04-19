package user_handler

import (
	"encoding/json"
	"log/slog"
	"myavito/internal/dto"
	"strconv"

	"net/http"
)

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req dto.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("failed decode request", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.validator.Struct(req); err != nil {
		slog.Error("failed validate request", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := h.srv.ServiceCreateUser(ctx, req.Name, req.Phone_number)
	if err != nil {
		slog.Error("failed Save user", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	slog.Info("user save")

	response := map[string]string{
		"id": strconv.Itoa(id),
	}
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		slog.Error("failed Encode request", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}
