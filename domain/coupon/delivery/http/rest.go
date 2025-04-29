package http

import (
	"github.com/foodieeoo/domain/coupon"
	"github.com/foodieeoo/models"
	"github.com/foodieeoo/shared/util"
	"github.com/labstack/echo"
)

type handlerCoupon struct {
	usecase coupon.Usecase
	config  *models.Config
}

func NewHandlerCoupon(e *echo.Echo, usecase coupon.Usecase, config *models.Config) {
	handler := &handlerCoupon{
		usecase: usecase,
		config:  config,
	}

	e.POST("api/coupon", handler.AddCouponCodes)
}

func (h *handlerCoupon) AddCouponCodes(c echo.Context) error {
	resp := h.usecase.AddCouponCodes(c)
	if resp.Error != nil {
		return util.ApiResponse(c, "error", resp.Data, resp.Message, resp.StatusCode, nil)
	}

	return util.ApiResponse(c, "success", resp.Data, resp.Message, resp.StatusCode, nil)
}
