/*
 * Twilio - Proxy
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListSessionResponse struct for ListSessionResponse
type ListSessionResponse struct {
	Meta     ListServiceResponseMeta `json:"Meta,omitempty"`
	Sessions []ProxyV1ServiceSession `json:"Sessions,omitempty"`
}
