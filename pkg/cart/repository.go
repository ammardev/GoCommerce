package cart

import (
	"github.com/ammardev/gocommerce/internal/connections"
	"github.com/ammardev/gocommerce/internal/persistence"
	"github.com/ammardev/gocommerce/pkg/products"
)

type CartRepository struct {
	persistence.BaseMySqlRepository
}

func (repo *CartRepository) GetCartBySessionId(sessionId string) (*Cart, error) {
	cart := Cart{}

	err := connections.DB.Get(&cart, "select * from carts where session_id=? limit 1", sessionId)
	if err != nil {
		return nil, err
	}

	cart.Items, err = repo.getCartItems(cart.ID)
	if err != nil {
		return nil, err
	}

	return &cart, nil
}

type cartItemResults struct {
    CartItem
    products.Product
}

func (repo *CartRepository) getCartItems(cartId int64) ([]CartItem, error) {
    results := []cartItemResults{}

    query := `
        select products.id, products.title, products.price, cart_items.price as price_on_addition, cart_items.quantity
        from cart_items
        join products on products.id = cart_items.product_id
        where cart_id=?
    `

	err := connections.DB.Select(&results, query, cartId)
    if err != nil {
        return nil, err
    }

    items := []CartItem{}

    for _, result := range results {
        result.CartItem.Product = result.Product
        items = append(items, result.CartItem)
    }

    return items, nil
}
