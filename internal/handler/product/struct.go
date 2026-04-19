package product_handler

import (
	"myavito/internal/logger"
	"myavito/internal/service"

	"github.com/go-playground/validator/v10"
)

type ProductHandler struct {
	srv       *service.ProductService
	validator *validator.Validate
	log       *logger.Logger
}

func NewProductHandler(srv *service.ProductService, val *validator.Validate, log *logger.Logger) *ProductHandler {
	return &ProductHandler{
		srv:       srv,
		validator: val,
		log:       log,
	}
}
