package client

import (
	"errors"
	"fmt"
	"strings"
)

// ErrInvalidPathParam is returned (wrapped) by ValidatePathParam when a path
// parameter value could alter the request target. Use errors.Is to detect it.
var ErrInvalidPathParam = errors.New("invalid path parameter: must not be empty, '.', '..' or contain '/', '?', '#'")

// ValidatePathParam rejects path parameter values that would change which
// endpoint a request is sent to when interpolated into a URL path. It returns
// an error if the value is empty, exactly "." or "..", or contains '/', '?'
// or '#'. Otherwise the value is returned unchanged.
func ValidatePathParam(value string) (string, error) {
	if value == "" || value == "." || value == ".." || strings.ContainsAny(value, "/?#") {
		return "", fmt.Errorf("%w: %q", ErrInvalidPathParam, value)
	}
	return value, nil
}
