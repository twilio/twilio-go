/*
 * Twilio - Preview
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// PreviewWirelessRatePlanReadResponse struct for PreviewWirelessRatePlanReadResponse
type PreviewWirelessRatePlanReadResponse struct {
	Meta PreviewBulkExportsExportDayReadResponseMeta `json:"Meta,omitempty"`
	RatePlans []PreviewWirelessRatePlan `json:"RatePlans,omitempty"`
}
