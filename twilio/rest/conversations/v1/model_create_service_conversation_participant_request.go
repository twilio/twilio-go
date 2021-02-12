/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)
// CreateServiceConversationParticipantRequest struct for CreateServiceConversationParticipantRequest
type CreateServiceConversationParticipantRequest struct {
	// An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \"{}\" will be returned.
	Attributes string `json:"Attributes,omitempty"`
	// The date that this resource was created.
	DateCreated time.Time `json:"DateCreated,omitempty"`
	// The date that this resource was last updated.
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	// A unique string identifier for the conversation participant as [Conversation User](https://www.twilio.com/docs/conversations/api/user-resource). This parameter is non-null if (and only if) the participant is using the Conversation SDK to communicate. Limited to 256 characters.
	Identity string `json:"Identity,omitempty"`
	// The address of the participant's device, e.g. a phone or WhatsApp number. Together with the Proxy address, this determines a participant uniquely. This field (with proxy_address) is only null when the participant is interacting from an SDK endpoint (see the 'identity' field).
	MessagingBindingAddress string `json:"MessagingBindingAddress,omitempty"`
	// The address of the Twilio phone number that is used in Group MMS. Communication mask for the Conversation participant with Identity.
	MessagingBindingProjectedAddress string `json:"MessagingBindingProjectedAddress,omitempty"`
	// The address of the Twilio phone number (or WhatsApp number) that the participant is in contact with. This field, together with participant address, is only null when the participant is interacting from an SDK endpoint (see the 'identity' field).
	MessagingBindingProxyAddress string `json:"MessagingBindingProxyAddress,omitempty"`
	// The SID of a conversation-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the participant.
	RoleSid string `json:"RoleSid,omitempty"`
}
