/*
 * Twilio - Trusthub
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CreateTrustProductChannelEndpointAssignmentRequest struct for CreateTrustProductChannelEndpointAssignmentRequest
type CreateTrustProductChannelEndpointAssignmentRequest struct {
	// The SID of an channel endpoint
	ChannelEndpointSid string `json:"ChannelEndpointSid"`
	// The type of channel endpoint. eg: phone-number
	ChannelEndpointType string `json:"ChannelEndpointType"`
}
