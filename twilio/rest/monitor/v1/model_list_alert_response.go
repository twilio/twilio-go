/*
 * Twilio - Monitor
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListAlertResponse struct for ListAlertResponse
type ListAlertResponse struct {
	Alerts []MonitorV1Alert      `json:"Alerts,omitempty"`
	Meta   ListAlertResponseMeta `json:"Meta,omitempty"`
}
