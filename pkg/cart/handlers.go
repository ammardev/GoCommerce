package cart

import (
	"net/http"

	"github.com/ammardev/gocommerce/pkg/products"
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
	cart, err := repository.GetCartBySessionId(c.Request().Header.Get("X-CART"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, cart)
}

func addCartItem(c echo.Context) error {
	type addToCartRequest struct {
		ProductId int `json:"product_id"`
		Quantity  int `json:"quantity"`
	}

	cart := Cart{
		SessionID: c.Request().Header.Get("X-CART"),
	}

	if cart.SessionID == "" {
		cart.NewSessionId()
		cart.Save()
	} else {
		cart.Load()
	}

	request := &addToCartRequest{}
	c.Bind(&request)

	item := CartItem{
		Product: products.Product{
			ID: int64(request.ProductId),
		},
		Quantity: request.Quantity,
	}

	cart.AddItem(item)

	return echo.ErrNotImplemented
}

func changeCartItemQuantity(c echo.Context) error {
	return echo.ErrNotImplemented
}

func deleteCartItem(c echo.Context) error {
	return echo.ErrNotImplemented
}
