package webframeworks

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

// PRACTICE EXERCISE: Router Comparison
// Implement a standard API spec across BOTH Gin and Echo.
//
// The API Spec:
// 1. GET /ping -> returns JSON {"message": "pong"} with 200 OK
// 2. GET /users/:id -> returns JSON {"user_id": "<id>"} with 200 OK
// 3. POST /users -> binds JSON {"username": "<name>"}, returns JSON {"status": "created", "username": "<name>"} with 201 Created.
//    If username is empty, return 400 Bad Request.

type UserSpecRequest struct {
	Username string `json:"username" binding:"required"`
}

// SetupUnifiedAPIRoutes returns the requested router implementation for comparison.
func SetupUnifiedAPIRoutes(engineType string) (http.Handler, error) {
	switch engineType {
	case "gin":
		return setupGinUnified(), nil
	case "echo":
		return setupEchoUnified(), nil
	default:
		return nil, errors.New("unsupported engine type")
	}
}

// setupGinUnified builds the unified spec in Gin
func setupGinUnified() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{"user_id": id})
	})

	r.POST("/users", func(c *gin.Context) {
		var req UserSpecRequest
		if err := c.ShouldBindJSON(&req); err != nil || req.Username == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "username is required"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"status":   "created",
			"username": req.Username,
		})
	})

	return r
}

// setupEchoUnified builds the unified spec in Echo
func setupEchoUnified() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "pong"})
	})

	e.GET("/users/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.JSON(http.StatusOK, map[string]string{"user_id": id})
	})

	e.POST("/users", func(c echo.Context) error {
		// In Echo we use the same struct tag (json) since it binding relies on standard JSON tags.
		var req UserSpecRequest
		if err := c.Bind(&req); err != nil || req.Username == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "username is required"})
		}

		return c.JSON(http.StatusCreated, map[string]string{
			"status":   "created",
			"username": req.Username,
		})
	})

	return e
}
