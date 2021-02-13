/*
 * Twilio - Notify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListBindingResponse struct for ListBindingResponse
type ListBindingResponse struct {
	Bindings []NotifyV1ServiceBinding   `json:"Bindings,omitempty"`
	Meta     ListCredentialResponseMeta `json:"Meta,omitempty"`
}
