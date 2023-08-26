package main

import (
	"github.com/ammardev/ecommerce-playground/cart"
	"github.com/ammardev/ecommerce-playground/connections"
	"github.com/ammardev/ecommerce-playground/products"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	connections.NewMySqlConnection()
	defer connections.Close()

	router := echo.New()

	router.Use(middleware.Recover())

	products.RegisterRoutes(router)
	cart.RegisterRoutes(router)

	router.Logger.Fatal(router.Start(":3000"))
}
