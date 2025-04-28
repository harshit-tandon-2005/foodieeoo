package product

import (
	"context"

	ProductModels "github.com/foodieeoo/domain/product/models"
	"github.com/foodieeoo/models"
)

type Usecase interface {
	GetProduct(ctx context.Context, id string) (*ProductModels.ProductResponse, error)
	GetProducts(ctx context.Context) ([]ProductModels.ProductResponse, error)
}

type Repository interface {
	GetProduct(ctx context.Context, id string) (*models.Product, error)
	GetProducts(ctx context.Context) ([]models.Product, error)
}
