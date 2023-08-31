package products

import (
	"github.com/ammardev/gocommerce/internal/http"
)

type ProductRequest struct {
	Title       string
	Description string
	Price       float64
}

func (request *ProductRequest) Validate() error {
	errors := make(http.ValidationErrors)

	if len(request.Title) == 0 {
		errors["title"] = "The title is required"
	}

	if len(request.Description) == 0 {
		errors["description"] = "The description is required"
	}

	if request.Price == 0 {
		errors["price"] = "The price should be greater than zero"
	}

	return errors.Check()
}
