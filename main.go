package main

import (
	"github.com/ammardev/ecommerce-playground/products"
	"github.com/labstack/echo/v4"
)

func main() {
	router := echo.New()

	products.RegisterRoutes(router)

	router.Logger.Fatal(router.Start(":3000"))
}
