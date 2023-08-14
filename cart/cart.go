package cart

import "github.com/ammardev/ecommerce-playground/products"

type Cart struct {
	ID       int64
	SessionID string
	Products []CartProduct
}

type CartProduct struct {
	Product products.Product
	PriceOnAddition float64
}

