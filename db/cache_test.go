package db

import (
	"context"
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
func TestRedisConnectionPing(t *testing.T) {
	rdb, err := GetCacheClientWithDB(0)
	assert.NoError(t, err)

	_, err = rdb.Ping(context.Background()).Result()
	assert.NoError(t, err)
}
