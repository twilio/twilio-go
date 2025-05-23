/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Conversations
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

// Optional parameters for the method 'CreateConversationWithParticipants'
type CreateConversationWithParticipantsParams struct {
	// The X-Twilio-Webhook-Enabled HTTP request header
	XTwilioWebhookEnabled *string `json:"X-Twilio-Webhook-Enabled,omitempty"`
	// The human-readable name of this conversation, limited to 256 characters. Optional.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource's `sid` in the URL.
	UniqueName *string `json:"UniqueName,omitempty"`
	// The date that this resource was created.
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The date that this resource was last updated.
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// The unique ID of the [Messaging Service](https://www.twilio.com/docs/messaging/api/service-resource) this conversation belongs to.
	MessagingServiceSid *string `json:"MessagingServiceSid,omitempty"`
	// An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\"{}\\\" will be returned.
	Attributes *string `json:"Attributes,omitempty"`
	//
	State *string `json:"State,omitempty"`
	// ISO8601 duration when conversation will be switched to `inactive` state. Minimum value for this timer is 1 minute.
	TimersInactive *string `json:"Timers.Inactive,omitempty"`
	// ISO8601 duration when conversation will be switched to `closed` state. Minimum value for this timer is 10 minutes.
	TimersClosed *string `json:"Timers.Closed,omitempty"`
	// The default email address that will be used when sending outbound emails in this conversation.
	BindingsEmailAddress *string `json:"Bindings.Email.Address,omitempty"`
	// The default name that will be used when sending outbound emails in this conversation.
	BindingsEmailName *string `json:"Bindings.Email.Name,omitempty"`
	// The participant to be added to the conversation in JSON format. The JSON object attributes are as parameters in [Participant Resource](https://www.twilio.com/docs/conversations/api/conversation-participant-resource). The maximum number of participants that can be added in a single request is 10.
	Participant *[]string `json:"Participant,omitempty"`
}

func (params *CreateConversationWithParticipantsParams) SetXTwilioWebhookEnabled(XTwilioWebhookEnabled string) *CreateConversationWithParticipantsParams {
	params.XTwilioWebhookEnabled = &XTwilioWebhookEnabled
	return params
}
func (params *CreateConversationWithParticipantsParams) SetFriendlyName(FriendlyName string) *CreateConversationWithParticipantsParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateConversationWithParticipantsParams) SetUniqueName(UniqueName string) *CreateConversationWithParticipantsParams {
	params.UniqueName = &UniqueName
	return params
}
func (params *CreateConversationWithParticipantsParams) SetDateCreated(DateCreated time.Time) *CreateConversationWithParticipantsParams {
	params.DateCreated = &DateCreated
	return params
}
func (params *CreateConversationWithParticipantsParams) SetDateUpdated(DateUpdated time.Time) *CreateConversationWithParticipantsParams {
	params.DateUpdated = &DateUpdated
	return params
}
func (params *CreateConversationWithParticipantsParams) SetMessagingServiceSid(MessagingServiceSid string) *CreateConversationWithParticipantsParams {
	params.MessagingServiceSid = &MessagingServiceSid
	return params
}
func (params *CreateConversationWithParticipantsParams) SetAttributes(Attributes string) *CreateConversationWithParticipantsParams {
	params.Attributes = &Attributes
	return params
}
func (params *CreateConversationWithParticipantsParams) SetState(State string) *CreateConversationWithParticipantsParams {
	params.State = &State
	return params
}
func (params *CreateConversationWithParticipantsParams) SetTimersInactive(TimersInactive string) *CreateConversationWithParticipantsParams {
	params.TimersInactive = &TimersInactive
	return params
}
func (params *CreateConversationWithParticipantsParams) SetTimersClosed(TimersClosed string) *CreateConversationWithParticipantsParams {
	params.TimersClosed = &TimersClosed
	return params
}
func (params *CreateConversationWithParticipantsParams) SetBindingsEmailAddress(BindingsEmailAddress string) *CreateConversationWithParticipantsParams {
	params.BindingsEmailAddress = &BindingsEmailAddress
	return params
}
func (params *CreateConversationWithParticipantsParams) SetBindingsEmailName(BindingsEmailName string) *CreateConversationWithParticipantsParams {
	params.BindingsEmailName = &BindingsEmailName
	return params
}
func (params *CreateConversationWithParticipantsParams) SetParticipant(Participant []string) *CreateConversationWithParticipantsParams {
	params.Participant = &Participant
	return params
}

// Create a new conversation with the list of participants in your account's default service
func (c *ApiService) CreateConversationWithParticipants(params *CreateConversationWithParticipantsParams) (*ConversationsV1ConversationWithParticipants, error) {
	path := "/v1/ConversationWithParticipants"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}
	if params != nil && params.DateCreated != nil {
		data.Set("DateCreated", fmt.Sprint((*params.DateCreated).Format(time.RFC3339)))
	}
	if params != nil && params.DateUpdated != nil {
		data.Set("DateUpdated", fmt.Sprint((*params.DateUpdated).Format(time.RFC3339)))
	}
	if params != nil && params.MessagingServiceSid != nil {
		data.Set("MessagingServiceSid", *params.MessagingServiceSid)
	}
	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.State != nil {
		data.Set("State", fmt.Sprint(*params.State))
	}
	if params != nil && params.TimersInactive != nil {
		data.Set("Timers.Inactive", *params.TimersInactive)
	}
	if params != nil && params.TimersClosed != nil {
		data.Set("Timers.Closed", *params.TimersClosed)
	}
	if params != nil && params.BindingsEmailAddress != nil {
		data.Set("Bindings.Email.Address", *params.BindingsEmailAddress)
	}
	if params != nil && params.BindingsEmailName != nil {
		data.Set("Bindings.Email.Name", *params.BindingsEmailName)
	}
	if params != nil && params.Participant != nil {
		for _, item := range *params.Participant {
			data.Add("Participant", item)
		}
	}

	if params != nil && params.XTwilioWebhookEnabled != nil {
		headers["X-Twilio-Webhook-Enabled"] = *params.XTwilioWebhookEnabled
	}
	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ConversationWithParticipants{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
