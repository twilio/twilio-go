/*
 * Twilio - Notify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.12.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListBindingResponse struct for ListBindingResponse
type ListBindingResponse struct {
	Bindings []NotifyV1ServiceBinding   `json:"Bindings,omitempty"`
	Meta     ListCredentialResponseMeta `json:"Meta,omitempty"`
}
