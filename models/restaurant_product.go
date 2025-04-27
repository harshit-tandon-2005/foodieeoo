package models

import (
	"time"

	"gorm.io/gorm"
)

/*
RestaurantProduct is the join table between Restaurant and Product and represents a product offered by a specific restaurant
*/
type RestaurantProduct struct {
	gorm.Model
	ID           int64     `gorm:"column:id" json:"id"`
	RestaurantID int64     `gorm:"column:restaurant_id" json:"restaurantId"`
	ProductID    int64     `gorm:"column:product_id" json:"productId"`
	IsAvailable  bool      `gorm:"column:is_available" json:"isAvailable"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updatedAt"`

	Restaurant Restaurant `gorm:"foreignKey:RestaurantID" json:"restaurant"`
	Product    Product    `gorm:"foreignKey:ProductID" json:"product"`
}
