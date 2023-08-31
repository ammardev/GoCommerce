package products

import "github.com/ammardev/gocommerce/internal/connections"

type ProductRepository struct{}

func (repo *ProductRepository) createProductFromRequest(request *ProductRequest) (*Product, error) {

	result, err := connections.DB.NamedExec(`
		insert into products (title, description, price)
		values (:title, :description, :price)
	`, &request)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	return &Product{
		ID:          id,
		Title:       request.Title,
		Description: request.Description,
		Price:       request.Price,
	}, nil
}
