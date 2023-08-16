package cart

import "github.com/ammardev/ecommerce-playground/products"

type Cart struct {
	ID        int64
	SessionID string
	Items     []CartItem
}

type CartItem struct {
	Product         products.Product
	Quantity        int
	PriceOnAddition float64
}
