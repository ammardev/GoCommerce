package cart

import (
	"database/sql"
	"log"

	"github.com/ammardev/gocommerce/internal/connections"
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

func (cart *Cart) Save() error {
	result, err := connections.DB.Exec("insert into carts (session_id) values (?)", cart.SessionID)

	if err != nil {
		return err
	}

	cart.ID, err = result.LastInsertId()

	return err
}

func (cart *Cart) NewSessionId() {
	// TODO: Generate the session id
	cart.SessionID = "TODO"
}

func (cart *Cart) AddItem(item CartItem) error {
	quantity := cart.GetItemQuantity(item)
	if quantity > 0 {
		cart.UpdateItemQuantity(item, quantity+item.Quantity)
		return nil
	}

	insertionQuery := "insert into cart_items (product_id, cart_id, quantity, price) values (?, ?, ?, ?)"
	_, err := connections.DB.Exec(insertionQuery, item.Product.ID, cart.ID, quantity, item.Product.Price)

	if err != nil {
		log.Println(err)
	}

	return nil
}

func (cart *Cart) GetItemQuantity(item CartItem) int {
	quantity := 0
	err := connections.DB.Get(&quantity, "select quantity from cart_items where cart_id=? and product_id=? limit 1", cart.ID, item.Product.ID)
	if err == sql.ErrNoRows {
		return 0
	} else if err != nil {
		log.Println(err)
	}

	return quantity
}

func (cart *Cart) UpdateItemQuantity(item CartItem, newQuantity int) {
	_, err := connections.DB.Exec("update cart_items set quantity=? where cart_id=? and product_id=?", newQuantity, cart.ID, item.Product.ID)
	if err != nil {
		log.Println(err)
	}
}
