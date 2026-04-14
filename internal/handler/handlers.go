package handler

import (
	"context"
	"encoding/json"
	"log/slog"
	"myavito/internal/dto"
	"myavito/internal/repository"

	"net/http"

	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	Repo repository.Repo
	Ctx  context.Context
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("failed decode request", "error", err)
	}
	validator := validator.New()
	if err := validator.Struct(req); err != nil {
		slog.Error("failed validate request", "error", err)
	}

	h.Repo.SaveUser(h.Ctx, &req)
}
