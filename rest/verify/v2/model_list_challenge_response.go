/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListChallengeResponse struct for ListChallengeResponse
type ListChallengeResponse struct {
	Challenges []VerifyV2Challenge                 `json:"challenges,omitempty"`
	Meta       ListVerificationAttemptResponseMeta `json:"meta,omitempty"`
}
