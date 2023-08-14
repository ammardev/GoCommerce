package cart

import "github.com/labstack/echo/v4"

func RegisterRoutes(router *echo.Echo) {
	router.GET("/carts/:sessionID", nil)
	router.POST("/carts/:sessionID/products", nil)
	router.DELETE("/carts/:sessionID/products/:id", nil)
}
