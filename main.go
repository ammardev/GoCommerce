package main

import (
	"github.com/ammardev/ecommerce-playground/cart"
	"github.com/ammardev/ecommerce-playground/connections"
	"github.com/ammardev/ecommerce-playground/products"
	"github.com/labstack/echo/v4"
)

func main() {
	connections.NewMySqlConnection()
	defer connections.Close()

	router := echo.New()

	products.RegisterRoutes(router)
	cart.RegisterRoutes(router)

	router.Logger.Fatal(router.Start(":3000"))
}
