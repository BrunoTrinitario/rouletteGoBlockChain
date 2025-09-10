package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRouter devuelve un router configurado con todas las rutas
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", pingHandler)
	return r
}

// Handlers
func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
