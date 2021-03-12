/*
    * Twilio - Bulkexports
    *
    * This is the public Twilio REST API.
    *
    * API version: 1.10.0
    * Contact: support@twilio.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi
// ListDayResponse struct for ListDayResponse
type ListDayResponse struct {
	Days []BulkexportsV1ExportDay `json:"Days,omitempty"`
	Meta ListDayResponseMeta `json:"Meta,omitempty"`
}
