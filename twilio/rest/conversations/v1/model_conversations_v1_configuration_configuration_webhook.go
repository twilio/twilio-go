/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ConversationsV1ConfigurationConfigurationWebhook struct for ConversationsV1ConfigurationConfigurationWebhook
type ConversationsV1ConfigurationConfigurationWebhook struct {
	AccountSid     *string                     `json:"AccountSid,omitempty"`
	Filters        *[]string                   `json:"Filters,omitempty"`
	Method         *ConfigurationWebhookMethod `json:"Method,omitempty"`
	PostWebhookUrl *string                     `json:"PostWebhookUrl,omitempty"`
	PreWebhookUrl  *string                     `json:"PreWebhookUrl,omitempty"`
	Target         *ConfigurationWebhookTarget `json:"Target,omitempty"`
	Url            *string                     `json:"Url,omitempty"`
}
