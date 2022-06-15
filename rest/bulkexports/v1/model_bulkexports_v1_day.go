/*
 * Twilio - Bulkexports
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// BulkexportsV1Day struct for BulkexportsV1Day
type BulkexportsV1Day struct {
	// The date when resource is created
	CreateDate *string `json:"create_date,omitempty"`
	// The date of the data in the file
	Day *string `json:"day,omitempty"`
	// The friendly name specified when creating the job
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The type of communication – Messages, Calls, Conferences, and Participants
	ResourceType *string `json:"resource_type,omitempty"`
	// Size of the file in bytes
	Size *int `json:"size,omitempty"`
}
