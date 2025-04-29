package usecase

import (
	"github.com/foodieeoo/domain/coupon"
	"gorm.io/gorm"
)

type couponUsecase struct {
	couponRepo coupon.Repository
	db         *gorm.DB
}

func NewCouponUsecase(couponRepo coupon.Repository, db *gorm.DB) coupon.Usecase {
	return &couponUsecase{
		couponRepo: couponRepo,
		db:         db,
	}
}
