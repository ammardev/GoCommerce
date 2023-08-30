package products

import "github.com/ammardev/ecommerce-playground/app"

type ProductRequest struct {
	Title       string
	Description string
	Price       float64
}

func (request *ProductRequest) Validate() *app.ValidationErrors {
	errors := make(app.ValidationErrors)

	if len(request.Title) == 0 {
		errors["title"] = "The title is required"
	}

	if len(request.Description) == 0 {
		errors["description"] = "The description is required"
	}

	if request.Price == 0 {
		errors["price"] = "The price should be greater than zero"
	}

	return &errors
}
