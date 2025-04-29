package models

import "time"

type Coupon struct {
	ID        int       `gorm:"column:id" json:"id"`
	Filename  string    `gorm:"column:filename" json:"filename"`
	Code      string    `gorm:"column:code" json:"code"`
	IsActive  bool      `gorm:"column:is_active" json:"isActive"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}
