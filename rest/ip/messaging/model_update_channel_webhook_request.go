/*
 * Twilio - Ip_messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// UpdateChannelWebhookRequest struct for UpdateChannelWebhookRequest
type UpdateChannelWebhookRequest struct {
	ConfigurationFilters []string `json:"ConfigurationFilters,omitempty"`
	ConfigurationFlowSid string `json:"ConfigurationFlowSid,omitempty"`
	ConfigurationMethod string `json:"ConfigurationMethod,omitempty"`
	ConfigurationRetryCount int32 `json:"ConfigurationRetryCount,omitempty"`
	ConfigurationTriggers []string `json:"ConfigurationTriggers,omitempty"`
	ConfigurationUrl string `json:"ConfigurationUrl,omitempty"`
}
