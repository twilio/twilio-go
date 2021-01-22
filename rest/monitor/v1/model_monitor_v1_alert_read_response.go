/*
 * Twilio - Monitor
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// MonitorV1AlertReadResponse struct for MonitorV1AlertReadResponse
type MonitorV1AlertReadResponse struct {
	Alerts []MonitorV1Alert `json:"Alerts,omitempty"`
	Meta MonitorV1AlertReadResponseMeta `json:"Meta,omitempty"`
}
