package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/snghnaveen/url-shortener/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiV1 := r.Group("/v1/api")

	apiV1.GET("/health/check", v1.HealthCheck)
	apiV1.POST("/shorten", v1.Shorten)
	apiV1.GET("/resolve/:shorten_key", v1.Resolve)
	apiV1.GET("/metrics-top-requested", v1.Metrics)
	return r
}
