package restapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIError represents a standardized RFC-7807 style error response.
type APIError struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Details []string `json:"details,omitempty"`
}

// WriteError sends a formatted APIError to the client.
func WriteError(c *gin.Context, status int, message string, details ...string) {
	c.JSON(status, APIError{
		Status:  status,
		Message: message,
		Details: details,
	})
}

// ListProducts returns all products.
func ListProducts(repo *ProductRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		products := repo.GetAll()
		c.JSON(http.StatusOK, products)
	}
}

// GetProduct retrieves a product by its ID.
func GetProduct(repo *ProductRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		product, exists := repo.GetByID(id)
		if !exists {
			WriteError(c, http.StatusNotFound, "Product not found")
			return
		}
		c.JSON(http.StatusOK, product)
	}
}

// CreateProduct creates a new product and validates input.
func CreateProduct(repo *ProductRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var p Product
		// ShouldBindJSON binds JSON and triggers validation (using validator v10)
		if err := c.ShouldBindJSON(&p); err != nil {
			WriteError(c, http.StatusBadRequest, "Invalid request body", err.Error())
			return
		}

		created, err := repo.Create(p)
		if err != nil {
			if err == ErrSKUAlreadyExists {
				WriteError(c, http.StatusConflict, err.Error())
				return
			}
			WriteError(c, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		c.JSON(http.StatusCreated, created)
	}
}

// UpdateProduct updates product details partially.
func UpdateProduct(repo *ProductRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var req UpdateProductRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			WriteError(c, http.StatusBadRequest, "Invalid request body", err.Error())
			return
		}

		updated, err := repo.Update(id, req)
		if err != nil {
			if err == ErrProductNotFound {
				WriteError(c, http.StatusNotFound, err.Error())
				return
			}
			WriteError(c, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		c.JSON(http.StatusOK, updated)
	}
}

// DeleteProduct removes a product by ID.
func DeleteProduct(repo *ProductRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		err := repo.Delete(id)
		if err != nil {
			if err == ErrProductNotFound {
				WriteError(c, http.StatusNotFound, err.Error())
				return
			}
			WriteError(c, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		// REST best practice: Return HTTP 204 No Content for successful deletes
		c.Status(http.StatusNoContent)
	}
}
