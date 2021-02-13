/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListServiceResponse struct for ListServiceResponse
type ListServiceResponse struct {
	Meta     ListServiceResponseMeta `json:"Meta,omitempty"`
	Services []VerifyV2Service       `json:"Services,omitempty"`
}
