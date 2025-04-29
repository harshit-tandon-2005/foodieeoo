package models

import (
	SharedModels "github.com/foodieeoo/models"
)

type CreateOrderRequest struct {
	CouponCode string      `json:"couponCode"`
	Items      []OrderItem `json:"items"`
}

type OrderItem struct {
	ProductID string `json:"productId"`
	Quantity  int    `json:"quantity"`
}

type AvailableProducts struct {
	Products           SharedModels.Product
	RestaurantProducts SharedModels.RestaurantProduct
}

type CreateOrderResponse struct {
	OrderID  int64             `json:"id"`
	Items    []OrderItem       `json:"items"`
	Products []ProductResponse `json:"products"`
}

type ProductResponse struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}
