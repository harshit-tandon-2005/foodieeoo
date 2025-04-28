package usecase

import (
	"context"
	"fmt"

	ProductModels "github.com/foodieeoo/domain/product/models"
)

func (u *productUsecase) GetProducts(ctx context.Context) ([]ProductModels.ProductResponse, error) {
	products, err := u.productRepo.GetProducts(ctx)
	if err != nil {
		fmt.Printf("Error getting products, [Error: %s]", err.Error())
		return nil, err
	}

	resp := []ProductModels.ProductResponse{}

	for _, product := range products {
		resp = append(resp, ProductModels.ProductResponse{
			ID:       product.ID,
			Name:     product.Name,
			Price:    product.Price,
			Currency: product.Currency,
			Category: product.Category,
		})
	}
	return resp, nil
}

func (u *productUsecase) GetProduct(ctx context.Context, id string) (*ProductModels.ProductResponse, error) {

	product, err := u.productRepo.GetProduct(ctx, id)
	if err != nil {
		fmt.Printf("Error getting product for [Id: %s], [Error: %s]", id, err.Error())
		return nil, err
	}

	resp := ProductModels.ProductResponse{}

	resp = ProductModels.ProductResponse{
		ID:       product.ID,
		Name:     product.Name,
		Price:    product.Price,
		Currency: product.Currency,
		Category: product.Category,
	}
	return &resp, nil
}
