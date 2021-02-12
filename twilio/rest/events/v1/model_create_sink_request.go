/*
 * Twilio - Events
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CreateSinkRequest struct for CreateSinkRequest
type CreateSinkRequest struct {
	// A human readable description for the Sink **This value should not contain PII.**
	Description string `json:"Description"`
	// The information required for Twilio to connect to the provided Sink encoded as JSON.
	SinkConfiguration map[string]interface{} `json:"SinkConfiguration"`
	// The Sink type. Can only be \"kinesis\" or \"webhook\" currently.
	SinkType string `json:"SinkType"`
}
