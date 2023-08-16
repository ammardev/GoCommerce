package cart

import (
	"github.com/ammardev/ecommerce-playground/connections"
	"github.com/ammardev/ecommerce-playground/products"
)

type Cart struct {
	ID        int64      `json:"id"`
	SessionID string     `json:"session_id"`
	Items     []CartItem `json:"items"`
}

type CartItem struct {
	Product         products.Product `json:"product"`
	Quantity        int              `json:"quantity"`
	PriceOnAddition float64          `json:"price_on_addition" db:"price_on_addition"`
}

func (cart *Cart) Load() error {
	cart.Items = []CartItem{}

	if cart.ID > 0 {
		return connections.DB.Get(&cart.SessionID, "select session_id from carts where id=? limit 1", cart.ID)
	}

	return connections.DB.Get(&cart.ID, "select id from carts where session_id=? limit 1", cart.SessionID)
}

type cartItemPivot struct {
	CartItem
	products.Product
}

func (cart *Cart) LoadItems() error {
	query := `
		select cart_items.price as price_on_addition, cart_items.quantity, products.* from cart_items
		join products on products.id = cart_items.product_id
		where cart_id=?
	`

	pivotRecords := []cartItemPivot{}

	err := connections.DB.Unsafe().Select(&pivotRecords, query, cart.ID)

	for _, record := range pivotRecords {
		record.CartItem.Product = record.Product
		cart.Items = append(cart.Items, record.CartItem)
	}

	return err
}
