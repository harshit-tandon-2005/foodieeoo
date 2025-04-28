package models

import (
	"time"
)

/*
Product represents a food item available in the platform
*/
type Product struct {
	ID          int64     `gorm:"column:id" json:"id"`
	Code        string    `gorm:"column:code" json:"code"`
	Name        string    `gorm:"column:name" json:"name"`
	Description string    `gorm:"column:description" json:"description"`
	Price       float64   `gorm:"column:price" json:"price"`
	Currency    string    `gorm:"column:currency" json:"currency"`
	Category    string    `gorm:"column:category" json:"category"` // e.g., APPETIZER, MAIN_COURSE, DESSERT, DRINK
	Type        string    `gorm:"column:type" json:"type"`         // e.g., VEGETARIAN, NON_VEGETARIAN, VEGAN
	CreatedAt   time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updatedAt"`

	OrderItems         []OrderItem         `gorm:"foreignKey:ProductID" json:"orderItems"`         // One-to-Many relationship with OrderItem
	RestaurantProducts []RestaurantProduct `gorm:"foreignKey:ProductID" json:"restaurantProducts"` // Many-to-Many relationship through RestaurantProduct
}
