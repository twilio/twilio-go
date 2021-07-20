package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadLimits(t *testing.T) {
	assert.Equal(t, 5, ReadLimits(nil, setLimit(5)))
	assert.Equal(t, 5, ReadLimits(setPageSize(10), setLimit(5)))
	assert.Equal(t, 1000, ReadLimits(nil, setLimit(5000)))
	assert.Equal(t, 10, ReadLimits(setPageSize(10), nil))
	assert.Equal(t, 50, ReadLimits(nil, nil))
}

func setLimit(limit int) *int {
	return &limit
}

func setPageSize(pageSize int) *int {
	return &pageSize
}
