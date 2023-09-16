package cart

type AddToCartRequest struct {
    ProductId int `json:"product_id"`
    Quantity  int `json:"quantity"`
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
