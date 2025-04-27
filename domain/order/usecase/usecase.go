package usecase

import (
	"github.com/foodieeoo/domain/order"
)

type orderUsecase struct {
	orderRepo order.Repository
}

func NewOrderUsecase(orderRepo order.Repository) order.Usecase {
	return &orderUsecase{
		orderRepo: orderRepo,
	}
}
