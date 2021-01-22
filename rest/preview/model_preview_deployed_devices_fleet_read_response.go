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
// PreviewDeployedDevicesFleetReadResponse struct for PreviewDeployedDevicesFleetReadResponse
type PreviewDeployedDevicesFleetReadResponse struct {
	Fleets []PreviewDeployedDevicesFleet `json:"Fleets,omitempty"`
	Meta PreviewBulkExportsExportDayReadResponseMeta `json:"Meta,omitempty"`
}
