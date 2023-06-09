package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsProdMode(t *testing.T) {
	t.Setenv(EnvKey, EnvProd)
	assert.True(t, IsProdMode())

	t.Setenv(EnvKey, EnvTesting)
	assert.False(t, IsProdMode())
}
