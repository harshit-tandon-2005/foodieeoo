package repository

import (
	"github.com/foodieeoo/domain/order"
	"github.com/foodieeoo/models"
	"gorm.io/gorm"
)

type mysqlHandler struct {
	session *gorm.DB
}

func NewOrderRepository(db *gorm.DB) order.Repository {
	return &mysqlHandler{
		session: db,
	}
}

func (h *mysqlHandler) CreateOrder(tx *gorm.DB, order *models.Order) (*models.Order, error) {
	err := tx.Create(&order).Error
	return order, err
}

func (h *mysqlHandler) CreateOrderItems(tx *gorm.DB, orderItem []models.OrderItem) ([]models.OrderItem, error) {
	err := tx.Create(&orderItem).Error
	return orderItem, err
}

func (h *mysqlHandler) CreateInvoice(tx *gorm.DB, invoice *models.Invoice) (*models.Invoice, error) {
	err := tx.Create(&invoice).Error
	return invoice, err
}
