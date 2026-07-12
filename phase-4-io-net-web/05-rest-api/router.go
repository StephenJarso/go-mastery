package restapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes the Gin engine and configures v1 API routes.
func SetupRouter(repo *ProductRepository) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	// Register basic recovery and logging middleware
	r.Use(gin.Recovery())

	// Health check route
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "UP"})
	})

	// Version 1 API routes (e.g. /api/v1/products)
	v1 := r.Group("/api/v1")
	{
		products := v1.Group("/products")
		{
			products.GET("", ListProducts(repo))
			products.GET("/:id", GetProduct(repo))
			products.POST("", CreateProduct(repo))
			products.PUT("/:id", UpdateProduct(repo))
			products.DELETE("/:id", DeleteProduct(repo))
		}
	}

	return r
}
