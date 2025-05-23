package usecase

import (
	"github.com/foodieeoo/domain/coupon"
	"github.com/foodieeoo/domain/order"
	"github.com/foodieeoo/domain/product"
	"gorm.io/gorm"
)

type orderUsecase struct {
	orderRepo   order.Repository
	productRepo product.Repository
	couponRepo  coupon.Repository
	db          *gorm.DB
}

func NewOrderUsecase(orderRepo order.Repository, productRepo product.Repository, couponRepo coupon.Repository, db *gorm.DB) order.Usecase {
	return &orderUsecase{
		orderRepo:   orderRepo,
		productRepo: productRepo,
		couponRepo:  couponRepo,
		db:          db,
	}
}
