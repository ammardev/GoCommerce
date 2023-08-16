package products

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

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
	products := Products{}

	products.Select()

	return c.JSON(http.StatusOK, products)
}

func showProduct(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	product := Product{
		ID: id,
	}

	err := product.Load()
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

	product.Save()

	return c.JSON(http.StatusCreated, product)
}

func updateProduct(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	product := Product{
		ID: id,
	}
	c.Bind(&product)

	err := product.Update()
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, product)
}

func deleteProduct(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	product := Product{
		ID: id,
	}
	err := product.Delete()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Success",
	})
}
