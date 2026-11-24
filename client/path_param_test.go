package client

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatePathParam_Valid(t *testing.T) {
	for _, v := range []string{
		"FW00000000000000000000000000000000",
		"AC00000000000000000000000000000000",
		"+15551234567",
		"my name.txt",
		"...",
		".hidden",
		"a.b",
	} {
		got, err := ValidatePathParam(v)
		assert.NoError(t, err, v)
		assert.Equal(t, v, got)
	}
}

func TestValidatePathParam_Invalid(t *testing.T) {
	for _, v := range []string{"", ".", "..", "a/b", "a?b", "a#b", "../x", "/", "FW0/Executions/FN0/Context?x=1"} {
		got, err := ValidatePathParam(v)
		assert.Error(t, err, v)
		assert.True(t, errors.Is(err, ErrInvalidPathParam), v)
		assert.Equal(t, "", got)
	}
}
