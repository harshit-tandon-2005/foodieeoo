package repository

import (
	"github.com/foodieeoo/domain/order"
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
