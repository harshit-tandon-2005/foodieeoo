package http

import (
	"net/http"

	"github.com/foodieeoo/domain/order"
	"github.com/foodieeoo/models"
	"github.com/foodieeoo/shared/util"
	"github.com/labstack/echo"
)

type handlerOrder struct {
	usecase order.Usecase
	config  *models.Config
}

func NewHandlerOrder(e *echo.Echo, usecase order.Usecase, config *models.Config) {
	handler := &handlerOrder{
		usecase: usecase,
		config:  config,
	}

	e.POST("api/order", handler.CreateOrder)
}

func (h *handlerOrder) CreateOrder(c echo.Context) error {
	return util.ApiResponse(c, "success", nil, "Order created successfully", http.StatusOK, nil)
}
