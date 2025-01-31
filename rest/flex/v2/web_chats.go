/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Flex
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"net/url"
)

// Optional parameters for the method 'CreateWebChannel'
type CreateWebChannelParams struct {
	// The Ui-Version HTTP request header
	UiVersion *string `json:"Ui-Version,omitempty"`
	// The SID of the Conversations Address. See [Address Configuration Resource](https://www.twilio.com/docs/conversations/api/address-configuration-resource) for configuration details. When a conversation is created on the Flex backend, the callback URL will be set to the corresponding Studio Flow SID or webhook URL in your address configuration.
	AddressSid *string `json:"AddressSid,omitempty"`
	// The Conversation's friendly name. See the [Conversation resource](https://www.twilio.com/docs/conversations/api/conversation-resource) for an example.
	ChatFriendlyName *string `json:"ChatFriendlyName,omitempty"`
	// The Conversation participant's friendly name. See the [Conversation Participant Resource](https://www.twilio.com/docs/conversations/api/conversation-participant-resource) for an example.
	CustomerFriendlyName *string `json:"CustomerFriendlyName,omitempty"`
	// The pre-engagement data.
	PreEngagementData *string `json:"PreEngagementData,omitempty"`
}

func (params *CreateWebChannelParams) SetUiVersion(UiVersion string) *CreateWebChannelParams {
	params.UiVersion = &UiVersion
	return params
}
func (params *CreateWebChannelParams) SetAddressSid(AddressSid string) *CreateWebChannelParams {
	params.AddressSid = &AddressSid
	return params
}
func (params *CreateWebChannelParams) SetChatFriendlyName(ChatFriendlyName string) *CreateWebChannelParams {
	params.ChatFriendlyName = &ChatFriendlyName
	return params
}
func (params *CreateWebChannelParams) SetCustomerFriendlyName(CustomerFriendlyName string) *CreateWebChannelParams {
	params.CustomerFriendlyName = &CustomerFriendlyName
	return params
}
func (params *CreateWebChannelParams) SetPreEngagementData(PreEngagementData string) *CreateWebChannelParams {
	params.PreEngagementData = &PreEngagementData
	return params
}

func (c *ApiService) CreateWebChannel(params *CreateWebChannelParams) (*FlexV2WebChannel, error) {
	path := "/v2/WebChats"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.AddressSid != nil {
		data.Set("AddressSid", *params.AddressSid)
	}
	if params != nil && params.ChatFriendlyName != nil {
		data.Set("ChatFriendlyName", *params.ChatFriendlyName)
	}
	if params != nil && params.CustomerFriendlyName != nil {
		data.Set("CustomerFriendlyName", *params.CustomerFriendlyName)
	}
	if params != nil && params.PreEngagementData != nil {
		data.Set("PreEngagementData", *params.PreEngagementData)
	}

	if params != nil && params.UiVersion != nil {
		headers["Ui-Version"] = *params.UiVersion
	}
	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV2WebChannel{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
