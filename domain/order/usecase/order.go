package usecase

import (
	"errors"
	"fmt"
	"net/http"

	OrderModels "github.com/foodieeoo/domain/order/models"
	"github.com/foodieeoo/models"
	"github.com/foodieeoo/shared/constants"
	"github.com/foodieeoo/shared/util"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func (u *orderUsecase) CreateOrder(ctx echo.Context) models.ApiUsescaseResponse {

	req := OrderModels.CreateOrderRequest{}
	if err := ctx.Bind(&req); err != nil {
		return util.SetUsecaseResponse(nil, err, http.StatusBadRequest, "INVALID_REQUEST", "Invalid Request")
	}

	fmt.Printf("Request body for create order: %+v", req)

	// Hardcoding Resturant ID and user ID
	resturantID := 1
	userID := 1
	// TODO: Can test if a certain Restraunt Id and User Id exist or not, since it is hardcoded above won't add that validation now

	resp := u.ValidateProduct(ctx, resturantID, req)
	if resp.Error != nil {
		return resp
	}

	availableProducts := resp.Data.([]OrderModels.AvailableProducts)
	discount := 0.0

	return u.CreateOrderUtil(ctx, availableProducts, userID, req, discount)

}

func (u *orderUsecase) ValidateProduct(ctx echo.Context, resturantID int, req OrderModels.CreateOrderRequest) models.ApiUsescaseResponse {

	availableProducts := []OrderModels.AvailableProducts{}

	for _, item := range req.Items {
		product, err := u.productRepo.GetProduct(ctx.Request().Context(), item.ProductID)
		if err != nil {
			fmt.Printf("Error getting product: %s", err.Error())
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return util.SetUsecaseResponse(nil, err, http.StatusNotFound, "PRODUCT_NOT_FOUND", constants.PRODUCT_NOT_AVAILABLE)
			}
			return util.SetUsecaseResponse(nil, err, http.StatusInternalServerError, "PRODUCT_NOT_FOUND", constants.INTERNAL_SERVER_ERROR)
		}

		fmt.Printf("Product: %+v", product)

		resturantProduct, err := u.productRepo.GetResturantProduct(ctx.Request().Context(), resturantID, item.ProductID)
		if err != nil {
			fmt.Printf("Error getting resturant products: %s", err.Error())
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return util.SetUsecaseResponse(nil, err, http.StatusNotFound, "PRODUCT_NOT_FOUND", constants.PRODUCT_NOT_AVAILABLE)
			}
			return util.SetUsecaseResponse(nil, err, http.StatusInternalServerError, "PRODUCT_NOT_FOUND", constants.INTERNAL_SERVER_ERROR)
		}

		if !resturantProduct.IsAvailable {
			fmt.Printf("Product not available in Resturant: %+v", resturantProduct)
			return util.SetUsecaseResponse(nil, err, http.StatusNotFound, "PRODUCT_NOT_FOUND", constants.PRODUCT_NOT_AVAILABLE)
		}
		availableProducts = append(availableProducts, OrderModels.AvailableProducts{
			Products:           *product,
			RestaurantProducts: *resturantProduct,
		})
	}

	return util.SetUsecaseResponse(availableProducts, nil, http.StatusOK, "SUCCESS", "Success")
}

func (u *orderUsecase) CreateOrderUtil(ctx echo.Context, availableProducts []OrderModels.AvailableProducts, userID int, req OrderModels.CreateOrderRequest, discount float64) models.ApiUsescaseResponse {

	resturantID := availableProducts[0].RestaurantProducts.RestaurantID
	orderItems := []models.OrderItem{}
	totalAmount := 0.0

	tx := u.db.Begin()
	defer tx.Rollback()

	order := &models.Order{
		OrderId:      uuid.New().String(),
		UserID:       int64(userID),
		RestaurantID: int64(resturantID),
		Status:       constants.ORDER_STATUS_PENDING,
	}

	order, err := u.orderRepo.CreateOrder(tx, order)
	if err != nil {
		fmt.Printf("Error creating order [Order: %+v] Error: %s", order, err.Error())
		return util.SetUsecaseResponse(nil, err, http.StatusInternalServerError, "ORDER_CREATION_FAILED", constants.INTERNAL_SERVER_ERROR)
	}

	for index, item := range availableProducts {
		price := item.Products.Price * float64(req.Items[index].Quantity)
		orderItems = append(orderItems, models.OrderItem{
			ProductID: item.Products.ID,
			Quantity:  req.Items[index].Quantity,
			OrderID:   order.ID,
			Price:     price,
		})
		totalAmount += price
	}

	orderItems, err = u.orderRepo.CreateOrderItems(tx, orderItems)
	if err != nil {
		fmt.Printf("Error creating order items [OrderItems: %+v] Error: %s", orderItems, err.Error())
		return util.SetUsecaseResponse(nil, err, http.StatusInternalServerError, "ORDER_ITEMS_CREATION_FAILED", constants.INTERNAL_SERVER_ERROR)
	}

	invoice := &models.Invoice{
		OrderID:            order.ID,
		InvoiceNumber:      uuid.New().String(),
		UserID:             int64(userID),
		TotalAmount:        totalAmount,
		Discount:           discount,
		TotalPayableAmount: totalAmount - discount,
		PaymentMethod:      constants.PAYMENT_METHOD_CREDIT_CARD,
		PaymentStatus:      constants.PAYMENT_STATUS_PENDING,
	}

	invoice, err = u.orderRepo.CreateInvoice(tx, invoice)
	if err != nil {
		fmt.Printf("Error creating invoice [Invoice: %+v] Error: %s", invoice, err.Error())
		return util.SetUsecaseResponse(nil, err, http.StatusInternalServerError, "INVOICE_CREATION_FAILED", constants.INTERNAL_SERVER_ERROR)
	}

	err = tx.Commit().Error
	if err != nil {
		fmt.Printf("Error committing transaction: %s", err.Error())
		return util.SetUsecaseResponse(nil, err, http.StatusInternalServerError, "TRANSACTION_COMMIT_FAILED", constants.INTERNAL_SERVER_ERROR)
	}

	productResponse := []OrderModels.ProductResponse{}
	for _, item := range availableProducts {
		productResponse = append(productResponse, OrderModels.ProductResponse{
			ID:       item.Products.ID,
			Name:     item.Products.Name,
			Price:    item.Products.Price,
			Category: item.Products.Category,
		})
	}
	resp := OrderModels.CreateOrderResponse{
		OrderID:  order.ID,
		Items:    req.Items,
		Products: productResponse,
	}

	return util.SetUsecaseResponse(resp, nil, http.StatusCreated, "SUCCESS", "Successfully created order")
}
