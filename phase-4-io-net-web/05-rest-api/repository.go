package restapi

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var (
	ErrProductNotFound  = errors.New("product not found")
	ErrSKUAlreadyExists = errors.New("product with this SKU already exists")
)

// ProductRepository defines a thread-safe, in-memory repository for storing products.
type ProductRepository struct {
	mu       sync.RWMutex
	products map[string]Product
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		products: make(map[string]Product),
	}
}

// Create inserts a new product, ensuring SKU uniqueness.
func (r *ProductRepository) Create(p Product) (Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Check SKU uniqueness
	for _, existing := range r.products {
		if existing.SKU == p.SKU {
			return Product{}, ErrSKUAlreadyExists
		}
	}

	// Mock ID generation
	p.ID = fmt.Sprintf("prod-%d", len(r.products)+1)
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	r.products[p.ID] = p
	return p, nil
}

// GetByID retrieves a single product by its ID.
func (r *ProductRepository) GetByID(id string) (Product, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	p, exists := r.products[id]
	return p, exists
}

// GetAll returns a list of all products.
func (r *ProductRepository) GetAll() []Product {
	r.mu.RLock()
	defer r.mu.RUnlock()

	list := make([]Product, 0, len(r.products))
	for _, p := range r.products {
		list = append(list, p)
	}
	return list
}

// Update modifies fields on an existing product.
func (r *ProductRepository) Update(id string, req UpdateProductRequest) (Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	p, exists := r.products[id]
	if !exists {
		return Product{}, ErrProductNotFound
	}

	if req.Name != nil {
		p.Name = *req.Name
	}
	if req.Price != nil {
		p.Price = *req.Price
	}
	if req.Quantity != nil {
		p.Quantity = *req.Quantity
	}
	p.UpdatedAt = time.Now()

	r.products[id] = p
	return p, nil
}

// Delete removes a product by its ID.
func (r *ProductRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, exists := r.products[id]
	if !exists {
		return ErrProductNotFound
	}

	delete(r.products, id)
	return nil
}
