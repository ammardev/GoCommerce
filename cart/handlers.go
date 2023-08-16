package cart

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(router *echo.Echo) {
	router.GET("/cart", getCart)
	router.POST("/cart", addCartItem)
	router.PATCH("/cart/:cartItemId", changeCartItemQuantity)
	router.DELETE("/cart/:cartItemId", deleteCartItem)
}

func getCart(c echo.Context) error {
	cart := Cart{
		SessionID: c.Request().Header.Get("X-CART"),
	}

	err := cart.Load()
	if err == sql.ErrNoRows {
		return echo.ErrNotFound
	} else if err != nil {
		log.Fatal(err)
	}

	err = cart.LoadItems()
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, cart)
}

func addCartItem(c echo.Context) error {
	// TODO: Create a new cart if the header not exist. And add the item
	return echo.ErrNotImplemented
}

func changeCartItemQuantity(c echo.Context) error {
	return echo.ErrNotImplemented
}

func deleteCartItem(c echo.Context) error {
	return echo.ErrNotImplemented
}
