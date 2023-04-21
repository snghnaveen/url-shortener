package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"message": "working flawlessly!!!",
	})
}
