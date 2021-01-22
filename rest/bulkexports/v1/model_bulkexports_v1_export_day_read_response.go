/*
 * Twilio - Bulkexports
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// BulkexportsV1ExportDayReadResponse struct for BulkexportsV1ExportDayReadResponse
type BulkexportsV1ExportDayReadResponse struct {
	Days []BulkexportsV1ExportDay `json:"Days,omitempty"`
	Meta BulkexportsV1ExportDayReadResponseMeta `json:"Meta,omitempty"`
}
