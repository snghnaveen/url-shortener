package rest

import "github.com/gin-gonic/gin"

type Gin struct {
	C *gin.Context
}

type Response struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
}

func (g *Gin) Response(httpCode int, respBody interface{}, err error) {
	var hasError bool
	if err != nil {
		hasError = true
		respBody = err.Error()
	}

	g.C.JSON(httpCode,
		Response{
			Error: hasError,
			Data:  respBody,
		})
}
