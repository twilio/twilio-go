/*
 * Twilio - Flex
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListWebChannelResponse struct for ListWebChannelResponse
type ListWebChannelResponse struct {
	FlexChatChannels []FlexV1WebChannel      `json:"flex_chat_channels,omitempty"`
	Meta             ListChannelResponseMeta `json:"meta,omitempty"`
}
