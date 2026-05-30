package solutions

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)


func SetupGinPing() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	return r
}

func SetupEchoPing() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "pong"})
	})
	return e
}
