/*
 * Twilio - Bulkexports
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// UpdateExportConfigurationRequest struct for UpdateExportConfigurationRequest
type UpdateExportConfigurationRequest struct {
	// If true, Twilio will automatically generate every day's file when the day is over.
	Enabled bool `json:"Enabled,omitempty"`
	// Sets whether Twilio should call a webhook URL when the automatic generation is complete, using GET or POST. The actual destination is set in the webhook_url
	WebhookMethod string `json:"WebhookMethod,omitempty"`
	// Stores the URL destination for the method specified in webhook_method.
	WebhookUrl string `json:"WebhookUrl,omitempty"`
}
