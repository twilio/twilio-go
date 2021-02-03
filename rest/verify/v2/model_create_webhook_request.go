/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CreateWebhookRequest struct for CreateWebhookRequest
type CreateWebhookRequest struct {
	// The array of events that this Webhook is subscribed to. Possible event types: `*, factor.deleted, factor.created, factor.verified, challenge.approved, challenge.denied`
	EventTypes []string `json:"EventTypes"`
	// The string that you assigned to describe the webhook. **This value should not contain PII.**
	FriendlyName string `json:"FriendlyName"`
	// The webhook status. Default value is `enabled`. One of: `enabled` or `disabled`
	Status string `json:"Status,omitempty"`
	// The URL associated with this Webhook.
	WebhookUrl string `json:"WebhookUrl"`
}
