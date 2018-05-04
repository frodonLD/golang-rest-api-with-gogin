package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck handler
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"alive": true})
}
