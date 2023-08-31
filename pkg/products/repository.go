package products

import "github.com/ammardev/gocommerce/internal/connections"

type ProductRepository struct{}

func (repo *ProductRepository) SelectProductById(id int64) (*Product, error) {
	product := Product{}

	err := connections.DB.Get(&product, "select * from products where id = ?", id)
	if err != nil {
		return nil, err
	}
	
	return &product, nil
}

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
