package repository

import (
	"github.com/foodieeoo/domain/product"
	"gorm.io/gorm"
)

type mysqlHandler struct {
	session *gorm.DB
}

func NewProductRepository(db *gorm.DB) product.Repository {
	return &mysqlHandler{
		session: db,
	}
}
