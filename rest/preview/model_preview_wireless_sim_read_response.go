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
// PreviewWirelessSimReadResponse struct for PreviewWirelessSimReadResponse
type PreviewWirelessSimReadResponse struct {
	Meta PreviewBulkExportsExportDayReadResponseMeta `json:"Meta,omitempty"`
	Sims []PreviewWirelessSim `json:"Sims,omitempty"`
}
