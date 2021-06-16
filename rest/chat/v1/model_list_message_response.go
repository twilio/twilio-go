/*
 * Twilio - Chat
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.17.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListMessageResponse struct for ListMessageResponse
type ListMessageResponse struct {
	Messages []ChatV1ServiceChannelMessage `json:"messages,omitempty"`
	Meta     ListCredentialResponseMeta    `json:"meta,omitempty"`
}
