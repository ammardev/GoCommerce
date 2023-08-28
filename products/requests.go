package products

type ProductRequest struct {
	Title       string  `validate:"required"`
	Description string  `validate:"required"`
	Price       float64 `validate:"required"`
}
