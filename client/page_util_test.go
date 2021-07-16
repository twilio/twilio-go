package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadLimits(t *testing.T) {
	assert.Equal(t, 5, ReadLimits(nil, 5))
	assert.Equal(t, 5, ReadLimits(setPageSize(10), 5))
	assert.Equal(t, 1000, ReadLimits(nil, 5000))
	assert.Equal(t, 10, ReadLimits(setPageSize(10), 0))
	assert.Equal(t, 50, ReadLimits(nil, 0))
}

func setPageSize(pageSize int) *int {
	return &pageSize
}
