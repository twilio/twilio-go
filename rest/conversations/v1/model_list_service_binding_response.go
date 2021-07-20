/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListServiceBindingResponse struct for ListServiceBindingResponse
type ListServiceBindingResponse struct {
	Bindings []ConversationsV1ServiceServiceBinding `json:"bindings,omitempty"`
	Meta     ListConversationResponseMeta           `json:"meta,omitempty"`
}
