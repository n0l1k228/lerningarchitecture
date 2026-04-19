package user_handler

import (
	"myavito/internal/logger"
	"myavito/internal/service"

	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	srv       *service.UserService
	validator *validator.Validate
	log       *logger.Logger
}

func NewUserHandler(srv *service.UserService, val *validator.Validate, log *logger.Logger) *UserHandler {
	return &UserHandler{
		srv:       srv,
		validator: val,
		log:       log,
	}
}
