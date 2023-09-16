package cart

type AddToCartRequest struct {
    ProductId int `json:"product_id"`
    Quantity  int `json:"quantity"`
}

