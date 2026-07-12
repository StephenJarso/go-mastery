package restapi

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// PRACTICE EXERCISE: API Extension (Restock Endpoint)
// In a real inventory management system, stock adjustment (replenishing inventory)
// is a separate transaction, rather than updating the whole product resource.
//
// Implement a Restock request structure and a handler:
// 1. POST /api/v1/products/:id/restock
// 2. Request body: {"amount": <integer>} where amount must be >= 1.
// 3. If quantity is valid and product exists, add the amount to the current stock.
// 4. Return the updated Product with 200 OK.
// 5. Handle errors:
//    - Return 400 Bad Request if request binding fails or amount is < 1.
//    - Return 404 Not Found if the product ID does not exist.

type RestockRequest struct {
	Amount int `json:"amount" binding:"required,gt=0"`
}

// RestockProduct adjusts the inventory levels of a product.
func RestockProduct(repo *ProductRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var req RestockRequest
		// 1. Bind and validate the incoming JSON.
		if err := c.ShouldBindJSON(&req); err != nil {
			WriteError(c, http.StatusBadRequest, "Invalid restock amount. Amount must be greater than 0", err.Error())
			return
		}

		// 2. Retrieve the current product state.
		repo.mu.Lock()
		p, exists := repo.products[id]
		if !exists {
			repo.mu.Unlock()
			WriteError(c, http.StatusNotFound, ErrProductNotFound.Error())
			return
		}

		// 3. Adjust the quantity and update the timestamp.
		p.Quantity += req.Amount
		p.UpdatedAt = timeNow() // Using helper to mock time in tests if necessary
		repo.products[id] = p
		repo.mu.Unlock()

		// 4. Return the updated product state.
		c.JSON(http.StatusOK, p)
	}
}

// timeNow is a mockable helper for time.Now
var timeNow = func() time.Time {
	return time.Now()
}

// RegisterPracticeRoute adds the restock route to the given router engine.
func RegisterPracticeRoute(r *gin.Engine, repo *ProductRepository) {
	// The restock endpoint is grouped under v1 products.
	r.POST("/api/v1/products/:id/restock", RestockProduct(repo))
}
