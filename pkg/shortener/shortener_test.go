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

	c, err := db.GetCacheClientWithDB(DB0)
	assert.NoError(t, err)

	val, err := c.Get(context.TODO(), outShortenURL).Result()
	assert.NoError(t, err)
	assert.Contains(t, inURL, val)
}

func TestFetchShortenURxL(t *testing.T) {
	c, err := db.GetCacheClientWithDB(DB0)
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
	c, err := db.GetCacheClientWithDB(DB1)
	assert.NoError(t, err)

	// prepare some records
	url1 := "https://snghnaveen.1.io/path"
	url2 := "https://snghnaveen.2.io/path"
	url3 := "https://snghnaveen.3.io/path"
	url4 := "https://snghnaveen.4.io/path"
	url5 := "https://snghnaveen.5.io/path"

	for i := 1; i <= 100; i++ {
		if i%1 == 0 {
			assert.NoError(t, c.ZIncrBy(context.TODO(), reqCounterKey, 1, url1).Err())
		}

		if i%2 == 0 {
			assert.NoError(t, c.ZIncrBy(context.TODO(), reqCounterKey, 1, url2).Err())
		}

		if i%3 == 0 {
			assert.NoError(t, c.ZIncrBy(context.TODO(), reqCounterKey, 1, url3).Err())
		}

		if i%4 == 0 {
			assert.NoError(t, c.ZIncrBy(context.TODO(), reqCounterKey, 1, url4).Err())
		}

		if i%5 == 0 {
			assert.NoError(t, c.ZIncrBy(context.TODO(), reqCounterKey, 1, url5).Err())
		}
	}

	out, err := GetTopRequested(3)
	assert.NoError(t, err)
	assert.Len(t, out, 3)

	for i, v := range out {
		rank := i + 1
		if rank == 1 {
			assert.Equal(t, v["url"], url1)
			assert.Equal(t, v["rank"], 1)
			assert.Equal(t, v["score"], float64(100))
		}

		if rank == 2 {
			assert.Equal(t, v["url"], url2)
			assert.Equal(t, v["rank"], 2)
			assert.Equal(t, v["score"], float64(50))
		}

		if rank == 3 {
			assert.Equal(t, v["url"], url3)
			assert.Equal(t, v["rank"], 3)
			assert.Equal(t, v["score"], float64(33))
		}
	}
}
