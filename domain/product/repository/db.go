package repository

import (
	"context"

	"github.com/foodieeoo/domain/product"
	"github.com/foodieeoo/models"
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

func (h *mysqlHandler) GetProduct(ctx context.Context, id string) (*models.Product, error) {
	product := models.Product{}
	if err := h.session.WithContext(ctx).Where("id = ?", id).Unscoped().First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (h *mysqlHandler) GetProducts(ctx context.Context) ([]models.Product, error) {
	products := []models.Product{}
	if err := h.session.WithContext(ctx).Unscoped().Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}