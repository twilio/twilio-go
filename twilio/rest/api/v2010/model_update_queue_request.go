/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// UpdateQueueRequest struct for UpdateQueueRequest
type UpdateQueueRequest struct {
	// A descriptive string that you created to describe this resource. It can be up to 64 characters long.
	FriendlyName string `json:"FriendlyName,omitempty"`
	// The maximum number of calls allowed to be in the queue. The default is 100. The maximum is 5000.
	MaxSize int32 `json:"MaxSize,omitempty"`
}
