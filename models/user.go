package models

import (
	"time"
)

/*
User represents a user of the application
*/
type User struct {
	ID    int64  `gorm:"column:id" json:"id"`
	Name  string `gorm:"column:name" json:"name"`
	Email string `gorm:"column:email" json:"email"`
	// Password    string    `gorm:"column:password" json:"password"`
	// Type        string    `gorm:"column:type" json:"type"`   will represent the type of user e.g., CUSTOMER, RESTAURANT_OWNER, ADMIN
	PhoneNumber string    `gorm:"column:phone_number" json:"phoneNumber"`
	CountryCode string    `gorm:"column:country_code" json:"countryCode"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updatedAt"`

	Orders   []Order   `gorm:"foreignKey:UserID" json:"orders"`   // One-to-Many relationship with Order
	Invoices []Invoice `gorm:"foreignKey:UserID" json:"invoices"` // One-to-Many relationship with Invoice
}
