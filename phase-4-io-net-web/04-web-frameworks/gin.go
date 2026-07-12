package webframeworks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Gin is one of the most popular web frameworks in the Go ecosystem.
// It features a fast Radix tree router, built-in JSON/XML rendering,
// parameter binding/validation, and middle-ware chaining.
// Gin uses its own context type: *gin.Context, which abstracts ResponseWriter and Request.

type CreateItemRequest struct {
	Name  string `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required,gt=0"`
}

func SetupGinRouter() http.Handler {
	// For testing and output cleanliness, set Gin to ReleaseMode.
	// In development, you would use gin.DebugMode.
	gin.SetMode(gin.ReleaseMode)

	// Create a Gin engine with Default logger and recovery middleware.
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	// 1. Basic route
	r.GET("/", func(c *gin.Context) {
		// c.String writes string responses.
		c.String(http.StatusOK, "Hello from Gin!")
	})

	// 2. Route Grouping and parameters
	api := r.Group("/api")
	{
		// GET with path parameters (e.g., /api/items/42)
		api.GET("/items/:id", func(c *gin.Context) {
			// c.Param retrieves path parameters
			id := c.Param("id")
			// c.Query retrieves URL query parameters (e.g., /api/items/42?details=true)
			details := c.DefaultQuery("details", "false")

			// c.JSON serializes struct/map to JSON response
			c.JSON(http.StatusOK, gin.H{
				"router":  "gin",
				"item_id": id,
				"details": details,
			})
		})

		// POST with JSON request binding
		api.POST("/items", func(c *gin.Context) {
			var req CreateItemRequest

			// ShouldBindJSON binds the request JSON body to the struct
			// and automatically validates it using 'binding' struct tags (validator v10).
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusCreated, gin.H{
				"message": "Item created successfully",
				"name":    req.Name,
				"price":   req.Price,
			})
		})
	}

	return r
}
