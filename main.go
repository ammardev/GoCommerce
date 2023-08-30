package main

import (
	"net/http"

	"github.com/ammardev/ecommerce-playground/app"
	"github.com/ammardev/ecommerce-playground/cart"
	"github.com/ammardev/ecommerce-playground/connections"
	"github.com/ammardev/ecommerce-playground/products"

	_ "github.com/joho/godotenv/autoload"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	connections.NewMySqlConnection()
	defer connections.Close()

	router := echo.New()

	router.Use(middleware.Recover())

	router.HTTPErrorHandler = customHTTPErrorHandler

	products.RegisterRoutes(router)
	cart.RegisterRoutes(router)

	address := app.GetEnv("HTTP_HOST", "127.0.0.1") + ":" + app.GetEnv("HTTP_PORT", "3000")
	router.Logger.Fatal(router.Start(address))
}

func customHTTPErrorHandler(err error, c echo.Context) {
	skipLogging := false

	switch err.(type) {
	case *app.ValidationErrors:
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
