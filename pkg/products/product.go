package products

import (
	"github.com/ammardev/gocommerce/internal/connections"
)

type Product struct {
	ID          int64   `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func (product *Product) Update() error {
	_, err := connections.DB.NamedExec("update products set title=:title, description=:description, price=:price where id=:id", &product)
	return err
}

type Products []Product
