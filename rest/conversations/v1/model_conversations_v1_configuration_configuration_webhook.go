/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ConversationsV1ConfigurationConfigurationWebhook struct for ConversationsV1ConfigurationConfigurationWebhook
type ConversationsV1ConfigurationConfigurationWebhook struct {
	AccountSid string `json:"AccountSid,omitempty"`
	Filters []string `json:"Filters,omitempty"`
	Method string `json:"Method,omitempty"`
	PostWebhookUrl string `json:"PostWebhookUrl,omitempty"`
	PreWebhookUrl string `json:"PreWebhookUrl,omitempty"`
	Target string `json:"Target,omitempty"`
	Url string `json:"Url,omitempty"`
}
