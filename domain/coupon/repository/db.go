package repository

import (
	"github.com/foodieeoo/domain/coupon"
	"github.com/foodieeoo/models"
	"gorm.io/gorm"
)

type mysqlHandler struct {
	session *gorm.DB
}

func NewCouponRepository(db *gorm.DB) coupon.Repository {
	return &mysqlHandler{
		session: db,
	}
}

func (h *mysqlHandler) CreateCoupons(tx *gorm.DB, coupons []models.Coupon) error {
	err := tx.Create(&coupons).Error
	return err
}

func (h *mysqlHandler) GetCoupons(code string) ([]models.Coupon, error) {
	var coupons []models.Coupon
	err := h.session.Where("code = ? AND is_active = ?", code, true).Find(&coupons).Error
	return coupons, err
}
