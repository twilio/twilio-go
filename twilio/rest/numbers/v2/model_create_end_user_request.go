/*
 * Twilio - Numbers
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CreateEndUserRequest struct for CreateEndUserRequest
type CreateEndUserRequest struct {
	// The set of parameters that are the attributes of the End User resource which are derived End User Types.
	Attributes map[string]interface{} `json:"Attributes,omitempty"`
	// The string that you assigned to describe the resource.
	FriendlyName string `json:"FriendlyName"`
	// The type of end user of the Bundle resource - can be `individual` or `business`.
	Type string `json:"Type"`
}
