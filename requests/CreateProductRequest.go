package requests

type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Price       float64 `json:"price" binding:"required,gte=10,lte=1000"`
	Description string  `json:"description" binding:"required"`
}
