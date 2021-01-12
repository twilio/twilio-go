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
// InlineResponse20011 struct for InlineResponse20011
type InlineResponse20011 struct {
	DeliveryReceipts []ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt `json:"delivery_receipts,omitempty"`
	Meta InlineResponse200Meta `json:"meta,omitempty"`
}
