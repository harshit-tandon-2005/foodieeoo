package models

import (
	"time"
)

/*
Restaurant represents a restaurant entity
*/
type Restaurant struct {
	ID          int64     `gorm:"column:id" json:"id"`
	Name        string    `gorm:"column:name" json:"name"`
	Address     string    `gorm:"column:address" json:"address"`
	PhoneNumber string    `gorm:"column:phone_number" json:"phoneNumber"`
	CountryCode string    `gorm:"column:country_code" json:"countryCode"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updatedAt"`

	RestaurantProducts []RestaurantProduct `gorm:"foreignKey:RestaurantID" json:"restaurantProducts"` // Many-to-Many relationship with RestaurantProduct
	Orders             []Order             `gorm:"foreignKey:RestaurantID" json:"orders"`             // One-to-Many relationship with Order
}
