/*
 * Twilio - Sync
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// SyncV1StreamMessage struct for SyncV1StreamMessage
type SyncV1StreamMessage struct {
	// Stream Message body
	Data *interface{} `json:"data,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
}
