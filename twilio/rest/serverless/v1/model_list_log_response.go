/*
 * Twilio - Serverless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ListLogResponse struct for ListLogResponse
type ListLogResponse struct {
	Logs []ServerlessV1ServiceEnvironmentLog `json:"Logs,omitempty"`
	Meta ListServiceResponseMeta `json:"Meta,omitempty"`
}
