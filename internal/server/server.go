package server

import (
	"github.com/ammardev/gocommerce/app"
	"github.com/ammardev/gocommerce/cart"
	"github.com/ammardev/gocommerce/products"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start() {
	router := echo.New()

	router.Use(middleware.Recover())

	router.HTTPErrorHandler = httpErrorHandler

	products.RegisterRoutes(router)
	cart.RegisterRoutes(router)

	address := app.GetEnv("HTTP_HOST", "127.0.0.1") + ":" + app.GetEnv("HTTP_PORT", "3000")
	router.Logger.Fatal(router.Start(address))
}
