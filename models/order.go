package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

/*
Order represents a customer's order
*/
type Order struct {
	gorm.Model
	ID           int64     `gorm:"column:id" json:"id"`
	UserID       int64     `gorm:"column:user_id" json:"userId"`
	RestaurantID int64     `gorm:"column:restaurant_id" json:"restaurantId"`
	TotalAmount  float64   `gorm:"column:total_amount" json:"totalAmount"`
	Status       string    `gorm:"column:status" json:"status"` // e.g., PENDING, CONFIRMED, PREPARING, DELIVERED, CANCELLED
	CreatedAt    time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updatedAt"`

	OrderItems []OrderItem `gorm:"foreignKey:OrderID" json:"orderItems"` // One-to-Many relationship with OrderItem
}
