/*
 * Twilio - Proxy
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.12.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListSessionResponse struct for ListSessionResponse
type ListSessionResponse struct {
	Meta     ListServiceResponseMeta `json:"Meta,omitempty"`
	Sessions []ProxyV1ServiceSession `json:"Sessions,omitempty"`
}
