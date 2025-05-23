package http

import (
	"net/http"

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

	e.GET("api/product", handler.GetProducts)
	e.GET("api/product/:id", handler.GetProduct)
}

func (h *handlerProduct) GetProducts(c echo.Context) error {
	products, err := h.usecase.GetProducts(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, products)
}

func (h *handlerProduct) GetProduct(c echo.Context) error {
	id := c.Param("id")
	product, err := h.usecase.GetProduct(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, product)
}
