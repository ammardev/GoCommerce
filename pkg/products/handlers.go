package products

import (
	"fmt"
	net_http "net/http"
	"strconv"

	"github.com/ammardev/gocommerce/internal/connections"
	"github.com/ammardev/gocommerce/internal/http"
	"github.com/labstack/echo/v4"
)

var (
	repository ProductRepository
)

func RegisterRoutes(router *echo.Echo) {
	repository = ProductRepository{}

	router.GET("/products", listProducts)
	router.GET("/products/stream", productsStream)
	router.GET("/products/:id", showProduct)
	router.POST("/products", createProduct)
	router.PATCH("/products/:id", updateProduct)
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

	return c.JSON(net_http.StatusOK, products)
}

func productsStream(c echo.Context) error {
    sse := http.ServerSentEventManager{}
    sse.SetHeadersForContext(c)

    redisSubscription := connections.Redis.Subscribe(c.Request().Context(), "products")
    defer redisSubscription.Close()

    sse.Channel = redisSubscription.Channel()

    sse.Serve(c)

    return nil
}

func showProduct(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	product, err := repository.SelectProductById(id)
	if err != nil {
		return err
	}

	return c.JSON(net_http.StatusOK, product)
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

    connections.Redis.Publish(c.Request().Context(), "products", fmt.Sprintf("New Product: %d", product.ID))

	return c.JSON(net_http.StatusCreated, product)
}

func updateProduct(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	request := UpdateRequest{}

	c.Bind(&request)

	 err := repository.updateProductFieldsFromRequest(id, &request)
	if err != nil {
		return err
	}

    connections.Redis.Publish(c.Request().Context(), "products", "New Product: " + c.Param("id"))

	return c.JSON(net_http.StatusOK, map[string]string{
		"message": "Product Updated",
	})
}

func deleteProduct(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	err := repository.deleteProductById(id)
	if err != nil {
		return err
	}

	return c.JSON(net_http.StatusOK, map[string]string{
		"message": "Product deleted",
	})
}
