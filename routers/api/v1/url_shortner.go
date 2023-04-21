package v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snghnaveen/url-shortener/pkg/rest"
	"github.com/snghnaveen/url-shortener/pkg/shortener"
	"github.com/snghnaveen/url-shortener/util"
)

// Resolve returns or redirects the url from shorten url
func Resolve(c *gin.Context) {
	app := rest.Gin{C: c}

	reqShortenURL := c.Param("shorten_key")

	res, err := shortener.FetchShortenURL(reqShortenURL)

	if err != nil {
		if errors.Is(err, rest.ErrShortenURLNotFound) {
			app.Response(http.StatusBadRequest, "", err)
			return
		} else {
			app.Response(http.StatusInternalServerError, "", err)
			return
		}
	}

	if c.Query("type") == "json" {
		app.Response(http.StatusOK, gin.H{
			"url": res,
		}, nil)
	} else {
		// Not working on Microsoft edge browser.
		c.Redirect(http.StatusTemporaryRedirect, res)
	}
}

type ShortenReqBody struct {
	URL string `json:"url"`
}

// Shorten shortens a given url
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

// Metrics returns top requested domains
func Metrics(c *gin.Context) {
	fetchTopXRecords := 3
	app := rest.Gin{C: c}

	res, err := shortener.GetTopRequested(int64(util.GetConfig().MaxMetricsRecords))

	if err != nil {
		app.Response(http.StatusInternalServerError, "", err)
		return
	}

	app.Response(http.StatusOK, res, nil)
}
