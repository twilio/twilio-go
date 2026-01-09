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

// RestErrorV1 provides information about an unsuccessful request using Twilio API Standards V1.0.
// This error format is used by newer Twilio APIs that follow the V1.0 standards.
type RestErrorV1 struct {
	Code           int                    `json:"code"`
	Message        string                 `json:"message"`
	HttpStatusCode int                    `json:"httpStatusCode"`
	Params         map[string]interface{} `json:"params"`
	UserError      bool                   `json:"userError"`
}

// Error implements the error interface for RestErrorV1.
func (err *RestErrorV1) Error() string {
	paramsJSON, _ := json.Marshal(err.Params)
	return fmt.Sprintf("Status: %d - ApiError %d: %s (params: %s, userError: %t)",
		err.HttpStatusCode, err.Code, err.Message, paramsJSON, err.UserError)
}

// GetCode returns the error code.
func (err *RestErrorV1) GetCode() int {
	return err.Code
}

// GetMessage returns the error message.
func (err *RestErrorV1) GetMessage() string {
	return err.Message
}

// GetHttpStatusCode returns the HTTP status code.
func (err *RestErrorV1) GetHttpStatusCode() int {
	return err.HttpStatusCode
}

// GetParams returns additional information about the error.
func (err *RestErrorV1) GetParams() map[string]interface{} {
	return err.Params
}

// GetUserError returns true if this is an error that depends on the end user's actions.
func (err *RestErrorV1) GetUserError() bool {
	return err.UserError
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
