package products

import (
	"github.com/ammardev/gocommerce/internal/connections"
	"github.com/ammardev/gocommerce/internal/persistence"
)

const ITEMS_PER_PAGE = 5

type ProductRepository struct {
	persistence.BaseMySqlRepository
}

func (repo *ProductRepository) SelectPaginatedProducts(currentPage int) (*Products, error) {
	products := Products{}

	offset := (currentPage - 1) * ITEMS_PER_PAGE

	err := connections.DB.Select(&products, "select * from products limit ? offset ?", ITEMS_PER_PAGE, offset)
	if err != nil {
		return nil, err
	}

	return &products, nil
}

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

func (repo *ProductRepository) updateProductFieldsFromRequest(id int64, request *UpdateRequest) error {
	builder := repo.UpdatesBuilder

	builder.Add("title", request.Title)
	builder.Add("description", request.Description)
	builder.Add("price", request.Price)

	builder.Values = append(builder.Values, id)

	query := "update products set " + builder.GetQuery() + " where id=?"
	_, err := connections.DB.Exec(query, builder.Values...)

	return err
}

func (repo *ProductRepository) deleteProductById(id int64) error {
	_, err := connections.DB.Exec("delete from products where id = ?", id)
	return err
}
