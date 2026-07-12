package webframeworks

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Echo is a high performance, extensible, minimalist Go web framework.
// Similar to Gin, Echo uses its own Context type (*echo.Context) to manage requests
// and responses, and supports data binding, validation, and middleware.

func SetupEchoRouter() http.Handler {
	e := echo.New()

	// Suppress banner and logger info for clean test runs
	e.HideBanner = true
	e.HidePort = true

	// 1. Setup Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// 2. Basic route
	e.GET("/", func(c echo.Context) error {
		// c.String renders a plain text response.
		// Returns an error if writing the response fails.
		return c.String(http.StatusOK, "Hello from Echo!")
	})

	// 3. Group routes
	api := e.Group("/api")
	{
		// GET with path parameters (e.g. /api/items/42)
		api.GET("/items/:id", func(c echo.Context) error {
			// c.Param retrieves path parameters
			id := c.Param("id")
			// c.QueryParam retrieves query parameters (e.g. ?details=true)
			details := c.QueryParam("details")
			if details == "" {
				details = "false"
			}

			// c.JSON renders JSON response
			return c.JSON(http.StatusOK, map[string]string{
				"router":  "echo",
				"item_id": id,
				"details": details,
			})
		})

		// POST with body binding
		api.POST("/items", func(c echo.Context) error {
			// We can reuse the CreateItemRequest struct from gin.go
			var req CreateItemRequest

			// c.Bind binds request body to the struct based on Content-Type
			if err := c.Bind(&req); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
			}

			// In Echo, manual validation or custom validator configuration is typical.
			// Let's do simple manual checks since Echo does not validate by default.
			if req.Name == "" || req.Price <= 0 {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid name or price"})
			}

			return c.JSON(http.StatusCreated, map[string]interface{}{
				"message": "Item created successfully",
				"name":    req.Name,
				"price":   req.Price,
			})
		})
	}

	return e
}
