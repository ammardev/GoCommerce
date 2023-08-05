package products

import (
	"log"
	"net/http"

	"github.com/ammardev/ecommerce-playground/connections"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(router *echo.Echo) {
	router.GET("/products", listProducts)
	router.GET("/products/:id", showProduct)
	router.POST("/products", createProduct)
	router.PUT("/products/:id", updateProduct)
	router.DELETE("/products/:id", deleteProduct)
}

func listProducts(c echo.Context) error {
	rows, err := connections.DB.Query("select * from products")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	products := []Product{}

	for rows.Next() {
		product := Product{}
		rows.Scan(&product.ID, &product.Title, &product.Description, &product.Price)
		products = append(products, product)
	}

	return c.JSON(http.StatusOK, products)
}

func showProduct(c echo.Context) error {
	row := connections.DB.QueryRow("select * from products where id = ?", c.Param("id"))
	product := Product{}
	err := row.Scan(&product.ID, &product.Title, &product.Description, &product.Price)
	if err != nil {
		return echo.ErrNotFound
	}

	return c.JSON(http.StatusOK, product)
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
