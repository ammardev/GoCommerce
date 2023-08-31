package products

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var (
	repository ProductRepository
)

func RegisterRoutes(router *echo.Echo) {
	repository = ProductRepository{}

	router.GET("/products", listProducts)
	router.GET("/products/:id", showProduct)
	router.POST("/products", createProduct)
	router.PUT("/products/:id", updateProduct)
	router.DELETE("/products/:id", deleteProduct)
}

func listProducts(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page == 0 {
		page = 1
	}

	products, err := repository.SelectPaginatedProducts(page)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, products)
}

func showProduct(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	product, err := repository.SelectProductById(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, product)
}

func createProduct(c echo.Context) error {
	request := ProductRequest{}
	c.Bind(&request)

	var err error

	err = request.Validate()
	if err != nil {
		return err
	}

	product, err := repository.createProductFromRequest(&request)

	if err != nil {
		return err
	}

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
		log.Panic(err)
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
		log.Panicf("%+v\n", err)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Success",
	})
}
