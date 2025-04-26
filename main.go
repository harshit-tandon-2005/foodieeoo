package main

import (
	"fmt"
	"os"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/foodieeoo/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/yaml.v2"
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

	e.Start(fmt.Sprintf(":%d", config.Port))
	gracehttp.Serve(e.Server)
}
