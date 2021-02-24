/*
 * Twilio - Chat
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// CreateMessageRequest struct for CreateMessageRequest
type CreateMessageRequest struct {
	// A valid JSON string that contains application-specific data.
	Attributes string `json:"Attributes,omitempty"`
	// The message to send to the channel. Can be an empty string or `null`, which sets the value as an empty string. You can send structured data in the body by serializing it as a string.
	Body string `json:"Body,omitempty"`
	// The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was created. The default value is the current time set by the Chat service. This parameter should only be used when a Chat's history is being recreated from a backup/separate source.
	DateCreated time.Time `json:"DateCreated,omitempty"`
	// The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was last updated.
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	// The [Identity](https://www.twilio.com/docs/chat/identity) of the new message's author. The default value is `system`.
	From string `json:"From,omitempty"`
	// The [Identity](https://www.twilio.com/docs/chat/identity) of the User who last updated the Message, if applicable.
	LastUpdatedBy string `json:"LastUpdatedBy,omitempty"`
	// The SID of the [Media](https://www.twilio.com/docs/chat/rest/media) to attach to the new Message.
	MediaSid string `json:"MediaSid,omitempty"`
}
