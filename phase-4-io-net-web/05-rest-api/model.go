package restapi

import "time"

// Product represents a product resource in our inventory REST API.
// We use JSON struct tags for request/response serialization,
// and 'binding' tags for automatic Gin validation.
type Product struct {
	ID        string    `json:"id"`
	SKU       string    `json:"sku" binding:"required,min=3,max=10"`
	Name      string    `json:"name" binding:"required,min=2"`
	Price     float64   `json:"price" binding:"required,gt=0"`
	Quantity  int       `json:"quantity" binding:"required,gte=0"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UpdateProductRequest defines the schema for updating a product.
type UpdateProductRequest struct {
	Name     *string  `json:"name" binding:"omitempty,min=2"`
	Price    *float64 `json:"price" binding:"omitempty,gt=0"`
	Quantity *int     `json:"quantity" binding:"omitempty,gte=0"`
}
