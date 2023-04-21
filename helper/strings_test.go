package helper_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"mini-api/helper"
)

func TestGetHostDomain(t *testing.T) {
	testcases := []struct {
		input    string
		expected string
	}{
		{input: "-", expected: "-"},
		{input: "localhost:1323", expected: "localhost"},
		{input: "api.sk88.dev", expected: "sk88.dev"},
	}
	for _, tc := range testcases {
		result := helper.GetHostDomain(tc.input)
		assert.Equal(t, tc.expected, result)
	}
}

func TestGetHostURL(t *testing.T) {
	invalidUrl := "postgres://user:abc{DEf1=ghi@example.com:5432"
	url1 := helper.GetHostURL(invalidUrl)
	url2 := helper.GetHostURL("http://localhost:1323/swagger/index.html")
	expected := "http://localhost:1323"
	assert.Equal(t, invalidUrl, url1)
	assert.Equal(t, expected, url2)
}
