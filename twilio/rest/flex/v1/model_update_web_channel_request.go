/*
 * Twilio - Flex
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// UpdateWebChannelRequest struct for UpdateWebChannelRequest
type UpdateWebChannelRequest struct {
	// The chat status. Can only be `inactive`.
	ChatStatus string `json:"ChatStatus,omitempty"`
	// The post-engagement data.
	PostEngagementData string `json:"PostEngagementData,omitempty"`
}
