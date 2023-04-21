package resolver

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Setenv("ENVIRONMENT", "testing")
	result := m.Run()
	os.Exit(result)
}

func TestEncodeURL(t *testing.T) {
	testURL := "https://github.com/snghnaveen/url-shortener"
	out, err := EncodeURL(testURL)
	assert.NoError(t, err)
	assert.NotEmpty(t, out)
}

func TestFormatURL(t *testing.T) {
	t.Run("with valid url", func(t *testing.T) {
		testURL := "https://github.com/snghnaveen/url-shortener"
		out, err := FormatURL(testURL)
		assert.NoError(t, err)
		assert.Equal(t, "github.com", out.Hostname())
	})

	t.Run("with invalid url", func(t *testing.T) {
		testURL := "invalid-url"
		_, err := FormatURL(testURL)
		assert.Error(t, err)
	})
}

func TestAddScheme(t *testing.T) {
	assert.Contains(t, AddScheme("github.com/snghnaveen/url-shortener"), "http://")
}
