/*
 * Twilio - Serverless
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ServerlessV1ServiceEnvironmentLogReadResponse struct for ServerlessV1ServiceEnvironmentLogReadResponse
type ServerlessV1ServiceEnvironmentLogReadResponse struct {
	Logs []ServerlessV1ServiceEnvironmentLog `json:"Logs,omitempty"`
	Meta ServerlessV1ServiceReadResponseMeta `json:"Meta,omitempty"`
}
