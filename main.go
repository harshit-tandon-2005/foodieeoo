package main

import (
	"fmt"
	"os"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/foodieeoo/models"
	"github.com/foodieeoo/shared/database"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/yaml.v2"

	orderRepository "github.com/foodieeoo/domain/order/repository"
	productRepository "github.com/foodieeoo/domain/product/repository"

	orderUsecase "github.com/foodieeoo/domain/order/usecase"
	productUsecase "github.com/foodieeoo/domain/product/usecase"

	orderHandler "github.com/foodieeoo/domain/order/delivery/http"
	productHandler "github.com/foodieeoo/domain/product/delivery/http"
)

func main() {

	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Gzip())
	e.Use(middleware.CORS())

	// Read the config file
	configFile, _ := os.ReadFile("config.yml")

	// Define a variable of type Config
	config := models.Config{}

	// Unmarshal the YAML data into the config struct
	err := yaml.Unmarshal(configFile, &config)
	if err != nil {
		fmt.Printf("Error unmarshalling config data: %v\n", err)
		panic(err)
	}

	mysqlConfig := database.MysqlConfig{
		DatabaseHost:     config.DatabaseHost,
		DatabaseUser:     config.DatabaseUser,
		DatabasePassword: config.DatabasePassword,
		DatabasePort:     config.DatabasePort,
		DatabaseName:     config.DatabaseName,
	}

	mysql := database.NewMysql(mysqlConfig)

	mysqlSession, err := mysql.OpenMysqlConn()
	if err != nil {
		fmt.Printf("Error open mysql master connection, err: %s", err.Error())
		panic(err)
	}

	fmt.Printf("Starting server on port %d\n", config.ApplicationPort)

	orderRepo := orderRepository.NewOrderRepository(mysqlSession.Client)
	productRepo := productRepository.NewProductRepository(mysqlSession.Client)

	orderUsecase := orderUsecase.NewOrderUsecase(orderRepo, productRepo, mysqlSession.Client)
	productUsecase := productUsecase.NewProductUsecase(productRepo)

	orderHandler.NewHandlerOrder(e, orderUsecase, &config)
	productHandler.NewHandlerProduct(e, productUsecase, &config)

	e.Start(fmt.Sprintf(":%d", config.ApplicationPort))
	gracehttp.Serve(e.Server)
}
