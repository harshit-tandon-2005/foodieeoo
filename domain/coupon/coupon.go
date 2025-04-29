package coupon

import (
	"github.com/labstack/echo"

	"github.com/foodieeoo/models"
	"gorm.io/gorm"
)

type Usecase interface {
	AddCouponCodes(ctx echo.Context) models.ApiUsescaseResponse
}

type Repository interface {
	CreateCoupons(tx *gorm.DB, coupons []models.Coupon) error
	GetCoupons(code string) ([]models.Coupon, error)
}
