package helper_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"mini-api/helper"
)

func TestRandInt64(t *testing.T) {
	r := helper.RandInt64(5, 10)
	assert.Greater(t, r, int64(4))
	assert.Less(t, r, int64(11))
}

func TestRandInt64_Panic(t *testing.T) {
	f := func() { helper.RandInt64(1, 0) }
	assert.Panics(t, f)
}

func TestRandFloat64(t *testing.T) {
	r := helper.RandFloat64()
	assert.Greater(t, r, 0.0)
	assert.Less(t, r, 1.0)
}
