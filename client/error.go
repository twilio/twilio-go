// Package error provides the interface for Twilio specific errors.
package client

import (
	"encoding/json"
	"fmt"
)

// TwilioRestError provides information about an unsuccessful request.
type TwilioRestError struct {
	Code     int                    `json:"code"`
	Details  map[string]interface{} `json:"details"`
	Message  string                 `json:"message"`
	MoreInfo string                 `json:"more_info"`
	Status   int                    `json:"status"`
}

func (err *TwilioRestError) Error() string {
	detailsJSON, _ := json.Marshal(err.Details)
	return fmt.Sprintf("Status: %d - ApiError %d: %s (%s) More info: %s",
		err.Status, err.Code, err.Message, detailsJSON, err.MoreInfo)
}

// KeyError represents an error when a required key is missing from response metadata.
// This error is typically thrown when expected metadata fields are not present
// in the response data, which may indicate a malformed or incomplete response.
type KeyError struct {
	Key     string
	Message string
}

func (err *KeyError) Error() string {
	return fmt.Sprintf("KeyError: required field '%s' is missing - %s", err.Key, err.Message)
}

// NewKeyError creates a new KeyError with the specified key and message.
func NewKeyError(key string, message string) *KeyError {
	return &KeyError{Key: key, Message: message}
}
