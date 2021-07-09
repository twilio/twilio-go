/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.18.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"net/url"
)

// Fetch the global configuration of conversations on your account
func (c *ApiService) FetchConfiguration() (*ConversationsV1Configuration, error) {
	path := "/v1/Configuration"

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1Configuration{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateConfiguration'
type UpdateConfigurationParams struct {
	// The SID of the default [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) to use when creating a conversation.
	DefaultChatServiceSid *string `json:"DefaultChatServiceSid,omitempty"`
	// Default ISO8601 duration when conversation will be switched to `closed` state. Minimum value for this timer is 10 minutes.
	DefaultClosedTimer *string `json:"DefaultClosedTimer,omitempty"`
	// Default ISO8601 duration when conversation will be switched to `inactive` state. Minimum value for this timer is 1 minute.
	DefaultInactiveTimer *string `json:"DefaultInactiveTimer,omitempty"`
	// The SID of the default [Messaging Service](https://www.twilio.com/docs/sms/services/api) to use when creating a conversation.
	DefaultMessagingServiceSid *string `json:"DefaultMessagingServiceSid,omitempty"`
}

func (params *UpdateConfigurationParams) SetDefaultChatServiceSid(DefaultChatServiceSid string) *UpdateConfigurationParams {
	params.DefaultChatServiceSid = &DefaultChatServiceSid
	return params
}
func (params *UpdateConfigurationParams) SetDefaultClosedTimer(DefaultClosedTimer string) *UpdateConfigurationParams {
	params.DefaultClosedTimer = &DefaultClosedTimer
	return params
}
func (params *UpdateConfigurationParams) SetDefaultInactiveTimer(DefaultInactiveTimer string) *UpdateConfigurationParams {
	params.DefaultInactiveTimer = &DefaultInactiveTimer
	return params
}
func (params *UpdateConfigurationParams) SetDefaultMessagingServiceSid(DefaultMessagingServiceSid string) *UpdateConfigurationParams {
	params.DefaultMessagingServiceSid = &DefaultMessagingServiceSid
	return params
}

// Update the global configuration of conversations on your account
func (c *ApiService) UpdateConfiguration(params *UpdateConfigurationParams) (*ConversationsV1Configuration, error) {
	path := "/v1/Configuration"

	data := url.Values{}
	if params != nil && params.DefaultChatServiceSid != nil {
		data.Set("DefaultChatServiceSid", *params.DefaultChatServiceSid)
	}
	if params != nil && params.DefaultClosedTimer != nil {
		data.Set("DefaultClosedTimer", *params.DefaultClosedTimer)
	}
	if params != nil && params.DefaultInactiveTimer != nil {
		data.Set("DefaultInactiveTimer", *params.DefaultInactiveTimer)
	}
	if params != nil && params.DefaultMessagingServiceSid != nil {
		data.Set("DefaultMessagingServiceSid", *params.DefaultMessagingServiceSid)
	}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1Configuration{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
