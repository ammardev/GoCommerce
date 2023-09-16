package cart

import (
    "github.com/ammardev/gocommerce/pkg/products"
)

type Cart struct {
    ID        int64      `json:"id"`
    SessionID string     `json:"session_id" db:"session_id"`
    Items     []CartItem `json:"items"`
}

type CartItem struct {
    Product         products.Product `json:"product"`
    Quantity        int              `json:"quantity"`
    PriceOnAddition float64          `json:"price_on_addition" db:"price_on_addition"`
}
