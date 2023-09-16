package products

import (
	"net/http"
	"strconv"
	"time"

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

	return c.JSON(http.StatusOK, products)
}

func productsStream(c echo.Context) error {
    // SSE headers
    c.Response().Header().Add("Content-Type", "text/event-stream")
    c.Response().Header().Set("Cache-Control", "no-cache")
    c.Response().Header().Set("Connection", "keep-alive")

    // Make messages channel
    sseChannel := make(chan string)

    go func() {
        for {
            sseChannel <- "Hello"
            time.Sleep(5 * time.Second)
        }
    }()

    // Close the messages channel
    defer func() {
        close(sseChannel)
        sseChannel = nil
    }()

    for {
        select {
        case message := <- sseChannel:
            c.Response().Writer.Write([]byte("data: " + message + "\n\n"))
            c.Response().Flush()
        case <- c.Request().Context().Done():
            // Connection closed
            return nil
        }
    }
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
	request := UpdateRequest{}

	c.Bind(&request)

	 err := repository.updateProductFieldsFromRequest(id, &request)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Product Updated",
	})
}

func deleteProduct(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	err := repository.deleteProductById(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Product deleted",
	})
}
