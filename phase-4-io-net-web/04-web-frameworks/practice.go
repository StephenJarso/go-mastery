package webframeworks

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)


var _ = http.StatusOK
var _ = gin.New
var _ = echo.New

// Exercise 1: Gin Ping Setup
// Setup Gin router with GET /ping returning JSON {"message": "pong"} with 200 OK.
func SetupGinPing() *gin.Engine {
	// TODO: Implement
	return nil
}

// Exercise 2: Echo Ping Setup
// Setup Echo router with GET /ping returning JSON {"message": "pong"} with 200 OK.
func SetupEchoPing() *echo.Echo {
	// TODO: Implement
	return nil
}
