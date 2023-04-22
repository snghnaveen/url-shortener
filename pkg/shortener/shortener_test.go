package shortener

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/snghnaveen/url-shortener/db"
	"github.com/snghnaveen/url-shortener/util"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Setenv(util.EnvKey, util.EnvTesting)
	result := m.Run()
	os.Exit(result)
}
func TestShortenURL(t *testing.T) {
	inURL := "https://github.com/snghnaveen/url-shortener"

	outShortenURL, err := ShortenURL(inURL)
	assert.NoError(t, err)

	assert.NotEmpty(t, outShortenURL)

	c, err := db.GetCacheClientWithDB(db.DB0)
	assert.NoError(t, err)

	val, err := c.Get(context.TODO(), outShortenURL).Result()
	assert.NoError(t, err)
	assert.Contains(t, inURL, val)
}

func TestFetchShortenURxL(t *testing.T) {
	c, err := db.GetCacheClientWithDB(db.DB0)
	assert.NoError(t, err)

	t.Run("when key is present", func(t *testing.T) {
		testURL := "https://github.com/snghnaveen/url-shortener"
		testShortenKey := "foo-bar"
		assert.NoError(t,
			c.Set(context.TODO(),
				testShortenKey,
				testURL, time.Second*100).Err(),
		)

		res, err := FetchShortenURL(testShortenKey)
		assert.NoError(t, err)
		assert.Equal(t, res, testURL)
	})
	t.Run("when key is not present", func(t *testing.T) {
		_, err := FetchShortenURL("foo-bar-not-present")
		assert.Error(t, err)
	})
}

func TestGetTopRequested(t *testing.T) {
	assert.NoError(t, ForTestCreateTestingData())

	// refer to ForTestCreateTestingData for test data i.e domain
	domain1 := "snghnaveen.1.io"
	domain2 := "snghnaveen.2.io"
	domain3 := "snghnaveen.3.io"

	out, err := GetTopRequested(3)
	assert.NoError(t, err)
	assert.Len(t, out, 3)

	for i, v := range out {
		rank := i + 1
		if rank == 1 {
			assert.Equal(t, v["domain"], domain1)
			assert.Equal(t, v["rank"], 1)
			assert.Equal(t, v["score"], float64(101))
		}

		if rank == 2 {
			assert.Equal(t, v["domain"], domain2)
			assert.Equal(t, v["rank"], 2)
			assert.Equal(t, v["score"], float64(51))
		}

		if rank == 3 {
			assert.Equal(t, v["domain"], domain3)
			assert.Equal(t, v["rank"], 3)
			assert.Equal(t, v["score"], float64(34))
		}
	}
}
