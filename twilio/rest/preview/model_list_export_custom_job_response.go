/*
 * Twilio - Preview
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListExportCustomJobResponse struct for ListExportCustomJobResponse
type ListExportCustomJobResponse struct {
	Jobs []PreviewBulkExportsExportExportCustomJob `json:"Jobs,omitempty"`
	Meta ListDayResponseMeta                       `json:"Meta,omitempty"`
}
