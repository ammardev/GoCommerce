package cart

type addToCartRequest struct {
    ProductId int `json:"product_id"`
    Quantity  int `json:"quantity"`
    SessionId string `header:"X-CART"`
}


type setCartItemQuantityRequest struct {
    ProductId int64 `param:"cartItemId"`
    SessionId string `header:"X-CART"`
    Quantity int `json:"quantity"`
}

type deleteCartItemRequest struct {
    ProductId int64 `param:"cartItemId"`
    SessionId string `header:"X-CART"`
}
