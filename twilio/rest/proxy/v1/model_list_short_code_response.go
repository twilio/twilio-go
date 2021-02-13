/*
 * Twilio - Proxy
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListShortCodeResponse struct for ListShortCodeResponse
type ListShortCodeResponse struct {
	Meta       ListServiceResponseMeta   `json:"Meta,omitempty"`
	ShortCodes []ProxyV1ServiceShortCode `json:"ShortCodes,omitempty"`
}
