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
// CreateFieldRequest struct for CreateFieldRequest
type CreateFieldRequest struct {
	// The unique name or sid of the FieldType. It can be any [Built-in Field Type](https://www.twilio.com/docs/assistant/api/built-in-field-types) or the unique_name or the Field Type sid of a custom Field Type.
	FieldType string `json:"FieldType"`
	// A user-provided string that uniquely identifies this resource as an alternative to the sid. Unique up to 64 characters long.
	UniqueName string `json:"UniqueName"`
}
