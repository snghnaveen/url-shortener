package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck will be used if service is doing fine
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"message": "working flawlessly!!!",
	})
}
