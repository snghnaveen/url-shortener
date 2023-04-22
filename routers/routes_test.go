package routers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/snghnaveen/url-shortener/util"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Setenv(util.EnvKey, util.EnvTesting)
	result := m.Run()
	os.Exit(result)
}

func TestHealthCheck(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/v1/api/health/check", nil)
	assert.NoError(t, err)
	router := InitRouter()
	router.ServeHTTP(w, r)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.JSONEq(t, `{"message":"working flawlessly!!!"}`, w.Body.String())
}

func TestRoutesURLShorten(t *testing.T) {
	router := InitRouter()

	t.Run("shorten", func(t *testing.T) {
		w := httptest.NewRecorder()

		reqBody := `{"url":"https://snghnaveen.github.io/foo-bar-routes-test"}`
		r, err := http.NewRequest("POST", "/v1/api/shorten", strings.NewReader(reqBody))
		assert.NoError(t, err)
		router.ServeHTTP(w, r)

		assert.Equal(t, w.Code, http.StatusCreated)
		assert.JSONEq(t, `
		{
			"error": false,
			"data": {
				"shorten_key": "cYYx74UA",
				"shorten_url": "/resolve/cYYx74UA"
			}
		}
		`, w.Body.String())
	})

	// Depends on above test
	t.Run("resolve", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, err := http.NewRequest("GET", "/v1/api/resolve/cYYx74UA?type=json", nil)
		assert.NoError(t, err)
		router.ServeHTTP(w, r)

		assert.Equal(t, w.Code, http.StatusOK)
		assert.JSONEq(t, `
		{"error":false,"data":{"url":"https://snghnaveen.github.io/foo-bar-routes-test"}}
		`, w.Body.String())

	})

	t.Run("metrics", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, err := http.NewRequest("GET", "/v1/api/metrics-top-requested", nil)
		assert.NoError(t, err)
		router.ServeHTTP(w, r)

		assert.Equal(t, w.Code, http.StatusOK)
		assert.JSONEq(t, `
		{"error":false,"data":[{"rank":1,"score":1,"domain":"snghnaveen.github.io"}]}
		`, w.Body.String())
	})
}
