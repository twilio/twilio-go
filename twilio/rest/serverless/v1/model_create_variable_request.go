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

// CreateVariableRequest struct for CreateVariableRequest
type CreateVariableRequest struct {
	// A string by which the Variable resource can be referenced. It can be a maximum of 128 characters.
	Key string `json:"Key"`
	// A string that contains the actual value of the Variable. It can be a maximum of 450 bytes in size.
	Value string `json:"Value"`
}
