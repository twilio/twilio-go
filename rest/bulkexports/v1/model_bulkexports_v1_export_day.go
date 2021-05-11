/*
 * Twilio - Bulkexports
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// BulkexportsV1ExportDay struct for BulkexportsV1ExportDay
type BulkexportsV1ExportDay struct {
	CreateDate   *string `json:"create_date,omitempty"`
	Day          *string `json:"day,omitempty"`
	FriendlyName *string `json:"friendly_name,omitempty"`
	ResourceType *string `json:"resource_type,omitempty"`
	Size         *int32  `json:"size,omitempty"`
}
