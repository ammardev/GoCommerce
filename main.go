package main

import (
	"net/http"

	"github.com/ammardev/ecommerce-playground/app"
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

	app.NewValidator()

	router.HTTPErrorHandler = customHTTPErrorHandler

	products.RegisterRoutes(router)
	cart.RegisterRoutes(router)

	router.Logger.Fatal(router.Start(":3000"))
}

func customHTTPErrorHandler(err error, c echo.Context) {
	skipLogging := false

	switch err.(type) {
	case *app.ValidationError:
		c.JSON(http.StatusUnprocessableEntity, err)
		skipLogging = true
	case *echo.HTTPError:
		c.JSON(http.StatusInternalServerError, echo.HTTPError{
			Message: err.Error(),
		})
		skipLogging = true
	default:
		c.JSON(http.StatusInternalServerError, echo.HTTPError{
			Message: "Internal Error",
		})
	}

	if !skipLogging {
		c.Logger().Error(err)
	}
}
