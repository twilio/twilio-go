/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CreateTaskChannelRequest struct for CreateTaskChannelRequest
type CreateTaskChannelRequest struct {
	// Whether the Task Channel should prioritize Workers that have been idle. If `true`, Workers that have been idle the longest are prioritized.
	ChannelOptimizedRouting bool `json:"ChannelOptimizedRouting,omitempty"`
	// A descriptive string that you create to describe the Task Channel. It can be up to 64 characters long.
	FriendlyName string `json:"FriendlyName"`
	// An application-defined string that uniquely identifies the Task Channel, such as `voice` or `sms`.
	UniqueName string `json:"UniqueName"`
}
