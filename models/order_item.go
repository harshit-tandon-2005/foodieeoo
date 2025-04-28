package models

import (
	"time"
)

/*
OrderItem represents a single item within an order
*/
type OrderItem struct {
	ID        int64     `gorm:"column:id" json:"id"`
	OrderID   int64     `gorm:"column:order_id" json:"orderId"`
	ProductID int64     `gorm:"column:product_id" json:"productId"`
	Quantity  int       `gorm:"column:quantity" json:"quantity"`
	Price     float64   `gorm:"column:price" json:"price"` // Price at the time of order
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}
