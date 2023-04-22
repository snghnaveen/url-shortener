package routers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
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

	var resp map[string]interface{}
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))

	assert.Contains(t, resp["message"], "working flawlessly")
}
