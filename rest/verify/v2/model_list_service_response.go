/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListServiceResponse struct for ListServiceResponse
type ListServiceResponse struct {
	Meta     ListVerificationAttemptResponseMeta `json:"meta,omitempty"`
	Services []VerifyV2Service                   `json:"services,omitempty"`
}
