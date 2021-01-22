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
// BulkexportsV1ExportExportCustomJobReadResponse struct for BulkexportsV1ExportExportCustomJobReadResponse
type BulkexportsV1ExportExportCustomJobReadResponse struct {
	Jobs []BulkexportsV1ExportExportCustomJob `json:"Jobs,omitempty"`
	Meta BulkexportsV1ExportDayReadResponseMeta `json:"Meta,omitempty"`
}
