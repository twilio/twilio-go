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
// ServerlessV1ServiceEnvironmentVariableReadResponse struct for ServerlessV1ServiceEnvironmentVariableReadResponse
type ServerlessV1ServiceEnvironmentVariableReadResponse struct {
	Meta ServerlessV1ServiceReadResponseMeta `json:"meta,omitempty"`
	Variables []ServerlessV1ServiceEnvironmentVariable `json:"variables,omitempty"`
}
