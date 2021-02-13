/*
 * Twilio - Events
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CreateSinkValidateRequest struct for CreateSinkValidateRequest
type CreateSinkValidateRequest struct {
	// A 34 character string that uniquely identifies the test event for a Sink being validated.
	TestId string `json:"TestId"`
}
