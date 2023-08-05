package products

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var products []Product

func RegisterRoutes(router *echo.Echo) {
	products = []Product{
		{
			ID:          1,
			Title:       "test",
			Description: "test test",
		},
		{
			ID:          2,
			Title:       "IPhone",
			Description: "Apple IPhone",
		},
		{
			ID:          3,
			Title:       "Macbook",
			Description: "Apple Macbook",
		},
	}

	router.GET("/products", listProducts)
	router.GET("/products/:id", showProduct)
	router.POST("/products", createProduct)
	router.PUT("/products/:id", updateProduct)
	router.DELETE("/products/:id", deleteProduct)
}

func listProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

func showProduct(c echo.Context) error {
	productID, _ := strconv.Atoi(c.Param("id"))

	for _, product := range products {
		if product.ID == productID {
			return c.JSON(http.StatusOK, product)
		}
	}

	return echo.ErrNotFound
}

func createProduct(c echo.Context) error {
	return c.String(http.StatusOK, "creating product")
}

func updateProduct(c echo.Context) error {
	return c.String(http.StatusOK, "updating product")
}

func deleteProduct(c echo.Context) error {
	return c.String(http.StatusOK, "deleting product")
}
