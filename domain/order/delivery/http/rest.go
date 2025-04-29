package http

import (
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
	resp := h.usecase.CreateOrder(c)
	if resp.Error != nil {
		return util.ApiResponse(c, "error", resp.Data, resp.Message, resp.StatusCode, nil)
	}

	return util.ApiResponse(c, "success", resp.Data, resp.Message, resp.StatusCode, nil)
}
