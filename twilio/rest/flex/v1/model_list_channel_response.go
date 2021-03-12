/*
    * Twilio - Flex
    *
    * This is the public Twilio REST API.
    *
    * API version: 1.10.0
    * Contact: support@twilio.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi
// ListChannelResponse struct for ListChannelResponse
type ListChannelResponse struct {
	FlexChatChannels []FlexV1Channel `json:"FlexChatChannels,omitempty"`
	Meta ListChannelResponseMeta `json:"Meta,omitempty"`
}
