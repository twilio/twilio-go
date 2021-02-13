/*
 * Twilio - Preview
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListFleetResponse struct for ListFleetResponse
type ListFleetResponse struct {
	Fleets []PreviewDeployedDevicesFleet `json:"Fleets,omitempty"`
	Meta   ListDayResponseMeta           `json:"Meta,omitempty"`
}
