package cart

import (
	"database/sql"

	"github.com/ammardev/gocommerce/internal/connections"
	"github.com/ammardev/gocommerce/internal/persistence"
	"github.com/ammardev/gocommerce/pkg/products"
	"github.com/go-sql-driver/mysql"
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

    return &cart, nil
}

func (repo *CartRepository) GetCartWithItemsBySessionId(sessionId string) (*Cart, error) {
    cart, err := repo.GetCartBySessionId(sessionId)
    if err != nil {
        return nil, err
    }

	cart.Items, err = repo.getCartItems(cart.ID)
	if err != nil {
		return nil, err
	}

	return cart, nil
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

func (repo *CartRepository) addCartItem(sessionId string, request AddToCartRequest) error {
    cart, err := repo.GetCartBySessionId(sessionId)
    if err == sql.ErrNoRows {
        err = nil
        cart = &Cart{
            SessionID: sessionId,
        }

        result, _ := connections.DB.Exec("insert into carts (session_id) values (?)", cart.SessionID)
        cart.ID, _ = result.LastInsertId()
    } else if err != nil {
        return err
    }

    productsRepo := products.ProductRepository{}
    product, err := productsRepo.SelectProductById(int64(request.ProductId))
    if err != nil {
        return err
    }

	insertionQuery := "insert into cart_items (product_id, cart_id, quantity, price) values (?, ?, ?, ?)"
    _, err = connections.DB.Exec(insertionQuery, product.ID, cart.ID, request.Quantity, product.Price)
    if mysqlerr, ok := err.(*mysql.MySQLError); ok && mysqlerr.Number == 1062 {
        err = nil
        updateQuery := "update cart_items set quantity = quantity + ? where cart_id=? and product_id=?"
        _, err = connections.DB.Exec(updateQuery, request.Quantity, cart.ID, product.ID)
    } else if err != nil {
        return err
    }

    return nil
}

func (repo *CartRepository) setQuantity(sessionId string, productId int64, quantity int) error {
    updateQuery := `
        update cart_items
        join carts on carts.id = cart_items.cart_id
        set quantity = ?
        where
            carts.session_id = ? and
            cart_items.product_id = ?
    `
    _, err := connections.DB.Exec(updateQuery, quantity, sessionId, productId)
    if err != nil {
        return err
    }

    return nil
}
