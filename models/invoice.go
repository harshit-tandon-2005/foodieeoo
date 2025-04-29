package models

import (
	"time"
)

/*
Order represents a customer's order
*/
type Invoice struct {
	ID                 int64      `gorm:"column:id" json:"id"`
	InvoiceNumber      string     `gorm:"column:invoice_number" json:"invoiceNumber"`
	OrderID            int64      `gorm:"column:order_id" json:"orderId"`
	UserID             int64      `gorm:"column:user_id" json:"userId"`
	TotalAmount        float64    `gorm:"column:total_amount" json:"totalAmount"`
	TotalPayableAmount float64    `gorm:"column:total_payable_amount" json:"totalPayableAmount"`
	Discount           float64    `gorm:"column:discount" json:"discount"`
	PaymentMethod      string     `gorm:"column:payment_method" json:"paymentMethod"` // e.g., NETBANKING, CREDIT_CARD, DEBIT_CARD, UPI
	PaymentStatus      string     `gorm:"column:payment_status" json:"paymentStatus"` // e.g., PENDING, COMPLETED, FAILED
	PaymentDate        *time.Time `gorm:"column:payment_date" json:"paymentDate"`
	CouponCode         string     `gorm:"column:coupon_code" json:"couponCode"`
	CreatedAt          time.Time  `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt          time.Time  `gorm:"column:updated_at" json:"updatedAt"`

	OrderItems []OrderItem `gorm:"foreignKey:OrderID" json:"orderItems"` // One-to-Many relationship with OrderItem
}
