/*
 * Twilio - Serverless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.16.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListFunctionVersionResponse struct for ListFunctionVersionResponse
type ListFunctionVersionResponse struct {
	FunctionVersions []ServerlessV1ServiceFunctionFunctionVersion `json:"function_versions,omitempty"`
	Meta             ListServiceResponseMeta                      `json:"meta,omitempty"`
}
