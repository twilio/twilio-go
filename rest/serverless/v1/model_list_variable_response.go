/*
 * Twilio - Serverless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.23.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListVariableResponse struct for ListVariableResponse
type ListVariableResponse struct {
	Meta      ListServiceResponseMeta `json:"meta,omitempty"`
	Variables []ServerlessV1Variable  `json:"variables,omitempty"`
}
