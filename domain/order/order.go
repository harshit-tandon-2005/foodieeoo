package order

import (
	"github.com/foodieeoo/models"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type Usecase interface {
	CreateOrder(ctx echo.Context) models.ApiUsescaseResponse
}

type Repository interface {
	CreateOrder(tx *gorm.DB, order *models.Order) (*models.Order, error)
	CreateOrderItems(tx *gorm.DB, orderItem []models.OrderItem) ([]models.OrderItem, error)
	CreateInvoice(tx *gorm.DB, invoice *models.Invoice) (*models.Invoice, error)
}
