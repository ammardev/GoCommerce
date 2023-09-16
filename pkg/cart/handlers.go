package cart

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var repository CartRepository

func RegisterRoutes(router *echo.Echo) {
	router.GET("/cart", getCart)
	router.POST("/cart", addCartItem)
	router.PATCH("/cart/:cartItemId", changeCartItemQuantity)
	router.DELETE("/cart/:cartItemId", deleteCartItem)
}

func getCart(c echo.Context) error {
	cart, err := repository.GetCartWithItemsBySessionId(c.Request().Header.Get("X-CART"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, cart)
}

func addCartItem(c echo.Context) error {
	request := &AddToCartRequest{}
	c.Bind(&request)

    repository.addCartItem(c.Request().Header.Get("X-CART"), *request)

	return c.JSON(http.StatusOK, map[string]string{
        "status": "Success",
    })
}

func changeCartItemQuantity(c echo.Context) error {
	return echo.ErrNotImplemented
}

func deleteCartItem(c echo.Context) error {
	return echo.ErrNotImplemented
}
