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
	"fmt"
	"net/url"

	"strings"
)

// Fetch push notification service settings
func (c *ApiService) FetchServiceNotification(ChatServiceSid string) (*ConversationsV1ServiceServiceConfigurationServiceNotification, error) {
	path := "/v1/Services/{ChatServiceSid}/Configuration/Notifications"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ServiceServiceConfigurationServiceNotification{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateServiceNotification'
type UpdateServiceNotificationParams struct {
	// Whether to send a notification when a participant is added to a conversation. The default is `false`.
	AddedToConversationEnabled *bool `json:"AddedToConversation.Enabled,omitempty"`
	// The name of the sound to play when a participant is added to a conversation and `added_to_conversation.enabled` is `true`.
	AddedToConversationSound *string `json:"AddedToConversation.Sound,omitempty"`
	// The template to use to create the notification text displayed when a participant is added to a conversation and `added_to_conversation.enabled` is `true`.
	AddedToConversationTemplate *string `json:"AddedToConversation.Template,omitempty"`
	// Weather the notification logging is enabled.
	LogEnabled *bool `json:"LogEnabled,omitempty"`
	// Whether the new message badge is enabled. The default is `false`.
	NewMessageBadgeCountEnabled *bool `json:"NewMessage.BadgeCountEnabled,omitempty"`
	// Whether to send a notification when a new message is added to a conversation. The default is `false`.
	NewMessageEnabled *bool `json:"NewMessage.Enabled,omitempty"`
	// The name of the sound to play when a new message is added to a conversation and `new_message.enabled` is `true`.
	NewMessageSound *string `json:"NewMessage.Sound,omitempty"`
	// The template to use to create the notification text displayed when a new message is added to a conversation and `new_message.enabled` is `true`.
	NewMessageTemplate *string `json:"NewMessage.Template,omitempty"`
	// Whether to send a notification to a user when they are removed from a conversation. The default is `false`.
	RemovedFromConversationEnabled *bool `json:"RemovedFromConversation.Enabled,omitempty"`
	// The name of the sound to play to a user when they are removed from a conversation and `removed_from_conversation.enabled` is `true`.
	RemovedFromConversationSound *string `json:"RemovedFromConversation.Sound,omitempty"`
	// The template to use to create the notification text displayed to a user when they are removed from a conversation and `removed_from_conversation.enabled` is `true`.
	RemovedFromConversationTemplate *string `json:"RemovedFromConversation.Template,omitempty"`
}

func (params *UpdateServiceNotificationParams) SetAddedToConversationEnabled(AddedToConversationEnabled bool) *UpdateServiceNotificationParams {
	params.AddedToConversationEnabled = &AddedToConversationEnabled
	return params
}
func (params *UpdateServiceNotificationParams) SetAddedToConversationSound(AddedToConversationSound string) *UpdateServiceNotificationParams {
	params.AddedToConversationSound = &AddedToConversationSound
	return params
}
func (params *UpdateServiceNotificationParams) SetAddedToConversationTemplate(AddedToConversationTemplate string) *UpdateServiceNotificationParams {
	params.AddedToConversationTemplate = &AddedToConversationTemplate
	return params
}
func (params *UpdateServiceNotificationParams) SetLogEnabled(LogEnabled bool) *UpdateServiceNotificationParams {
	params.LogEnabled = &LogEnabled
	return params
}
func (params *UpdateServiceNotificationParams) SetNewMessageBadgeCountEnabled(NewMessageBadgeCountEnabled bool) *UpdateServiceNotificationParams {
	params.NewMessageBadgeCountEnabled = &NewMessageBadgeCountEnabled
	return params
}
func (params *UpdateServiceNotificationParams) SetNewMessageEnabled(NewMessageEnabled bool) *UpdateServiceNotificationParams {
	params.NewMessageEnabled = &NewMessageEnabled
	return params
}
func (params *UpdateServiceNotificationParams) SetNewMessageSound(NewMessageSound string) *UpdateServiceNotificationParams {
	params.NewMessageSound = &NewMessageSound
	return params
}
func (params *UpdateServiceNotificationParams) SetNewMessageTemplate(NewMessageTemplate string) *UpdateServiceNotificationParams {
	params.NewMessageTemplate = &NewMessageTemplate
	return params
}
func (params *UpdateServiceNotificationParams) SetRemovedFromConversationEnabled(RemovedFromConversationEnabled bool) *UpdateServiceNotificationParams {
	params.RemovedFromConversationEnabled = &RemovedFromConversationEnabled
	return params
}
func (params *UpdateServiceNotificationParams) SetRemovedFromConversationSound(RemovedFromConversationSound string) *UpdateServiceNotificationParams {
	params.RemovedFromConversationSound = &RemovedFromConversationSound
	return params
}
func (params *UpdateServiceNotificationParams) SetRemovedFromConversationTemplate(RemovedFromConversationTemplate string) *UpdateServiceNotificationParams {
	params.RemovedFromConversationTemplate = &RemovedFromConversationTemplate
	return params
}

// Update push notification service settings
func (c *ApiService) UpdateServiceNotification(ChatServiceSid string, params *UpdateServiceNotificationParams) (*ConversationsV1ServiceServiceConfigurationServiceNotification, error) {
	path := "/v1/Services/{ChatServiceSid}/Configuration/Notifications"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)

	data := url.Values{}
	if params != nil && params.AddedToConversationEnabled != nil {
		data.Set("AddedToConversation.Enabled", fmt.Sprint(*params.AddedToConversationEnabled))
	}
	if params != nil && params.AddedToConversationSound != nil {
		data.Set("AddedToConversation.Sound", *params.AddedToConversationSound)
	}
	if params != nil && params.AddedToConversationTemplate != nil {
		data.Set("AddedToConversation.Template", *params.AddedToConversationTemplate)
	}
	if params != nil && params.LogEnabled != nil {
		data.Set("LogEnabled", fmt.Sprint(*params.LogEnabled))
	}
	if params != nil && params.NewMessageBadgeCountEnabled != nil {
		data.Set("NewMessage.BadgeCountEnabled", fmt.Sprint(*params.NewMessageBadgeCountEnabled))
	}
	if params != nil && params.NewMessageEnabled != nil {
		data.Set("NewMessage.Enabled", fmt.Sprint(*params.NewMessageEnabled))
	}
	if params != nil && params.NewMessageSound != nil {
		data.Set("NewMessage.Sound", *params.NewMessageSound)
	}
	if params != nil && params.NewMessageTemplate != nil {
		data.Set("NewMessage.Template", *params.NewMessageTemplate)
	}
	if params != nil && params.RemovedFromConversationEnabled != nil {
		data.Set("RemovedFromConversation.Enabled", fmt.Sprint(*params.RemovedFromConversationEnabled))
	}
	if params != nil && params.RemovedFromConversationSound != nil {
		data.Set("RemovedFromConversation.Sound", *params.RemovedFromConversationSound)
	}
	if params != nil && params.RemovedFromConversationTemplate != nil {
		data.Set("RemovedFromConversation.Template", *params.RemovedFromConversationTemplate)
	}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ServiceServiceConfigurationServiceNotification{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
