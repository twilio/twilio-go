/*
    * Twilio - Studio
    *
    * This is the public Twilio REST API.
    *
    * API version: 1.10.0
    * Contact: support@twilio.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi
// ListFlowRevisionResponse struct for ListFlowRevisionResponse
type ListFlowRevisionResponse struct {
	Meta ListFlowResponseMeta `json:"Meta,omitempty"`
	Revisions []StudioV2FlowFlowRevision `json:"Revisions,omitempty"`
}
