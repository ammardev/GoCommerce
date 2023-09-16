package cart

import (
    net_http "net/http"
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

	return c.JSON(net_http.StatusOK, cart)
}

func addCartItem(c echo.Context) error {
	request := &addToCartRequest{}
	c.Bind(&request)

    repository.addCartItem(*request)

	return c.JSON(net_http.StatusOK, map[string]string{
        "status": "Success",
    })
}

func changeCartItemQuantity(c echo.Context) error {
    request := &setCartItemQuantityRequest{}
    c.Bind(request)

    err := repository.setQuantity(request.SessionId, request.ProductId, request.Quantity)
	if err != nil {
		return err
	}

	return c.JSON(net_http.StatusOK, map[string]string{
        "status": "Success",
    })
}

func deleteCartItem(c echo.Context) error {
    request := &deleteCartItemRequest{}
    c.Bind(request)

    err := repository.removeCartItem(request.SessionId, request.ProductId)
    if err != nil {
        return err
    }

    return c.JSON(net_http.StatusOK, map[string]string{
        "status": "Success",
    })
}
