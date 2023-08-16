package cart

import "github.com/labstack/echo/v4"

func RegisterRoutes(router *echo.Echo) {
	router.GET("/cart", getCart)
	router.POST("/cart", addCartItem)
	router.PATCH("/cart/:cartItemId", changeCartItemQuantity)
	router.DELETE("/cart/:cartItemId", deleteCartItem)
}

func getCart(c echo.Context) error {
	// TODO: Return the cart. Get the session id from the headers
	return echo.ErrNotImplemented
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
