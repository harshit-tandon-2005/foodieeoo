package usecase

import (
	"github.com/foodieeoo/domain/product"
)

type productUsecase struct {
	productRepo product.Repository
}

func NewProductUsecase(productRepo product.Repository) product.Usecase {
	return &productUsecase{
		productRepo: productRepo,
	}
}
