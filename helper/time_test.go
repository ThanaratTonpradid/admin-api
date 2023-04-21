package helper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"mini-api/helper"
)

func TestParseBangkokTime(t *testing.T) {
	dateTime := "2006-01-02 15:04:05"
	layout := "2006-01-02 15:04:05"
	bangkokTime, err := helper.ParseBangkokTime(layout, dateTime)
	assert.Nil(t, err)
	assert.Equal(t, "2006-01-02T15:04:05+07:00", bangkokTime.Format(time.RFC3339))
}
