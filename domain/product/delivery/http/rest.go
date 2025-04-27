package http

import (
	"github.com/foodieeoo/domain/product"
	"github.com/foodieeoo/models"
	"github.com/labstack/echo"
)

type handlerProduct struct {
	usecase product.Usecase
	config  *models.Config
}

func NewHandlerProduct(e *echo.Echo, usecase product.Usecase, config *models.Config) {
	handler := &handlerProduct{
		usecase: usecase,
		config:  config,
	}

	e.POST("api/v1/products", handler.CreateProduct)
}

func (h *handlerProduct) CreateProduct(c echo.Context) error {
	return nil
}
