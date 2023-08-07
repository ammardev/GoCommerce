package products

import (
	"database/sql"
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
	products := []Product{}

	err := connections.DB.Select(&products, "select * from products")
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, products)
}

func showProduct(c echo.Context) error {
	product := Product{}
	err := connections.DB.Get(&product, "select * from products where id = ?", c.Param("id"))

	if err == sql.ErrNoRows {
		return echo.ErrNotFound
	}

	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	return c.JSON(http.StatusOK, product)
}

func createProduct(c echo.Context) error {
	product := Product{}
	c.Bind(&product)

	_, err := connections.DB.NamedExec("insert into products (title, description, price) values (:title, :description, :price)", &product)
	if err != nil {
		log.Fatal(err)
	}

	return c.String(http.StatusOK, "creating product")
}

func updateProduct(c echo.Context) error {
	return c.String(http.StatusOK, "updating product")
}

func deleteProduct(c echo.Context) error {
	_, err := connections.DB.Exec("delete from products where id = ?", c.Param("id"))
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Success",
	})
}
