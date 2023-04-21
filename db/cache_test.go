package db

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Setenv("ENVIRONMENT", "testing")
	result := m.Run()
	os.Exit(result)
}
func TestRedisConnectionPing(t *testing.T) {
	rdb, err := GetCacheClientWithDB(0)
	assert.NoError(t, err)

	_, err = rdb.Ping(context.Background()).Result()
	assert.NoError(t, err)
}

func ForTestCreateTestingData() {
	// @todo
}
