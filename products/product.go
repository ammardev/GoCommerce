package products

import "github.com/ammardev/ecommerce-playground/connections"

type Product struct {
	ID          int64   `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func (product *Product) Load() error {
	return connections.DB.Get(&product, "select * from products where id = ?", product.ID)
}

func (product *Product) Save() error {
	result, err := connections.DB.NamedExec("insert into products (title, description, price) values (:title, :description, :price)", &product)
	if err != nil {
		return err
	}

	product.ID, err = result.LastInsertId()

	return err
}

func (product *Product) Update() error {
	_, err := connections.DB.NamedExec("update products set title=:title, description=:description, price=:price where id=:id", &product)
	return err
}

func (product *Product) Delete() error {
	_, err := connections.DB.Exec("delete from products where id = ?", product.ID)
	return err
}

type Products []Product

func (products *Products) Select() error {
	return connections.DB.Select(products, "select * from products")
}