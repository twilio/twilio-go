/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ApiV2010AccountValidationRequest struct for ApiV2010AccountValidationRequest
type ApiV2010AccountValidationRequest struct {
	AccountSid string `json:"AccountSid,omitempty"`
	CallSid string `json:"CallSid,omitempty"`
	FriendlyName string `json:"FriendlyName,omitempty"`
	PhoneNumber string `json:"PhoneNumber,omitempty"`
	ValidationCode string `json:"ValidationCode,omitempty"`
}
