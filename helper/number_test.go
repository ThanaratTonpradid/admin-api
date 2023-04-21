package helper_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"mini-api/helper"
)

func TestMin(t *testing.T) {
	testcases := []struct {
		a        int64
		b        int64
		expected int64
	}{
		{a: 0, b: 1, expected: 0},
		{a: 10, b: 1, expected: 1},
	}
	for _, tc := range testcases {
		min := helper.Min(tc.a, tc.b)
		assert.Equal(t, tc.expected, min)
	}
}

func TestMax(t *testing.T) {
	testcases := []struct {
		a        int64
		b        int64
		expected int64
	}{
		{a: 0, b: 1, expected: 1},
		{a: 10, b: 1, expected: 10},
	}
	for _, tc := range testcases {
		min := helper.Max(tc.a, tc.b)
		assert.Equal(t, tc.expected, min)
	}
}

func TestToNumberPtr(t *testing.T) {
	num := 4
	result := helper.ToPtr(4)
	assert.Equal(t, &num, result)
}
