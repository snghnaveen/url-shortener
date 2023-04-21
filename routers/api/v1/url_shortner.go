package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snghnaveen/url-shortner/pkg/rest"
	"github.com/snghnaveen/url-shortner/pkg/shortener"
)

func Resolve(c *gin.Context) {
	app := rest.Gin{C: c}

	reqShortenURL := c.Param("shorten_key")

	res, err := shortener.FetchShortenURL(reqShortenURL)
	if err != nil {
		app.Response(http.StatusBadRequest, "", err)
		return
	}

	app.Response(http.StatusInternalServerError, gin.H{
		"url": res,
	}, nil)

	// TODO implement
	// c.Redirect(http.StatusTemporaryRedirect, "http://www.google.com")

}

type ShortenReqBody struct {
	URL string `json:"url"`
}

func Shorten(c *gin.Context) {
	app := rest.Gin{C: c}
	var req ShortenReqBody

	if err := c.ShouldBindJSON(&req); err != nil {
		app.Response(http.StatusBadRequest, "", err)
		return
	}

	shortenKey, err := shortener.ShortenURL(req.URL)
	if err != nil {
		app.Response(http.StatusInternalServerError, "", err)
		return
	}

	app.Response(http.StatusCreated, gin.H{
		"shorten_url": c.Request.Host + "/resolve/" + shortenKey,
		"shorten_key": shortenKey,
	}, nil)
}
