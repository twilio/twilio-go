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
	"strings"
	"time"
)

// Optional parameters for the method 'CreateServiceConversationWithParticipants'
type CreateServiceConversationWithParticipantsParams struct {
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

func (params *CreateServiceConversationWithParticipantsParams) SetXTwilioWebhookEnabled(XTwilioWebhookEnabled string) *CreateServiceConversationWithParticipantsParams {
	params.XTwilioWebhookEnabled = &XTwilioWebhookEnabled
	return params
}
func (params *CreateServiceConversationWithParticipantsParams) SetFriendlyName(FriendlyName string) *CreateServiceConversationWithParticipantsParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateServiceConversationWithParticipantsParams) SetUniqueName(UniqueName string) *CreateServiceConversationWithParticipantsParams {
	params.UniqueName = &UniqueName
	return params
}
func (params *CreateServiceConversationWithParticipantsParams) SetDateCreated(DateCreated time.Time) *CreateServiceConversationWithParticipantsParams {
	params.DateCreated = &DateCreated
	return params
}
func (params *CreateServiceConversationWithParticipantsParams) SetDateUpdated(DateUpdated time.Time) *CreateServiceConversationWithParticipantsParams {
	params.DateUpdated = &DateUpdated
	return params
}
func (params *CreateServiceConversationWithParticipantsParams) SetMessagingServiceSid(MessagingServiceSid string) *CreateServiceConversationWithParticipantsParams {
	params.MessagingServiceSid = &MessagingServiceSid
	return params
}
func (params *CreateServiceConversationWithParticipantsParams) SetAttributes(Attributes string) *CreateServiceConversationWithParticipantsParams {
	params.Attributes = &Attributes
	return params
}
func (params *CreateServiceConversationWithParticipantsParams) SetState(State string) *CreateServiceConversationWithParticipantsParams {
	params.State = &State
	return params
}
func (params *CreateServiceConversationWithParticipantsParams) SetTimersInactive(TimersInactive string) *CreateServiceConversationWithParticipantsParams {
	params.TimersInactive = &TimersInactive
	return params
}
func (params *CreateServiceConversationWithParticipantsParams) SetTimersClosed(TimersClosed string) *CreateServiceConversationWithParticipantsParams {
	params.TimersClosed = &TimersClosed
	return params
}
func (params *CreateServiceConversationWithParticipantsParams) SetBindingsEmailAddress(BindingsEmailAddress string) *CreateServiceConversationWithParticipantsParams {
	params.BindingsEmailAddress = &BindingsEmailAddress
	return params
}
func (params *CreateServiceConversationWithParticipantsParams) SetBindingsEmailName(BindingsEmailName string) *CreateServiceConversationWithParticipantsParams {
	params.BindingsEmailName = &BindingsEmailName
	return params
}
func (params *CreateServiceConversationWithParticipantsParams) SetParticipant(Participant []string) *CreateServiceConversationWithParticipantsParams {
	params.Participant = &Participant
	return params
}

// Create a new conversation with the list of participants in your account's default service
func (c *ApiService) CreateServiceConversationWithParticipants(ChatServiceSid string, params *CreateServiceConversationWithParticipantsParams) (*ConversationsV1ServiceConversationWithParticipants, error) {
	path := "/v1/Services/{ChatServiceSid}/ConversationWithParticipants"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)

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
		data.Set("State", *params.State)
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

	ps := &ConversationsV1ServiceConversationWithParticipants{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
