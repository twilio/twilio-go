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

    "github.com/twilio/twilio-go/client"
)


// Fetch the configuration of a conversation service
func (c *ApiService) FetchServiceConfiguration(ChatServiceSid string, ) (*ConversationsV1ServiceConfiguration, error) {
    path := "/v1/Services/{ChatServiceSid}/Configuration"
        path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)

data := url.Values{}
headers := make(map[string]interface{})




    resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ConversationsV1ServiceConfiguration{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Optional parameters for the method 'UpdateServiceConfiguration'
type UpdateServiceConfigurationParams struct {
    // The conversation-level role assigned to a conversation creator when they join a new conversation. See the [Conversation Role](https://www.twilio.com/docs/conversations/api/role-resource) for more info about roles.
    DefaultConversationCreatorRoleSid *string `json:"DefaultConversationCreatorRoleSid,omitempty"`
    // The conversation-level role assigned to users when they are added to a conversation. See the [Conversation Role](https://www.twilio.com/docs/conversations/api/role-resource) for more info about roles.
    DefaultConversationRoleSid *string `json:"DefaultConversationRoleSid,omitempty"`
    // The service-level role assigned to users when they are added to the service. See the [Conversation Role](https://www.twilio.com/docs/conversations/api/role-resource) for more info about roles.
    DefaultChatServiceRoleSid *string `json:"DefaultChatServiceRoleSid,omitempty"`
    // Whether the [Reachability Indicator](https://www.twilio.com/docs/chat/reachability-indicator) is enabled for this Conversations Service. The default is `false`.
    ReachabilityEnabled *bool `json:"ReachabilityEnabled,omitempty"`
}

func (params *UpdateServiceConfigurationParams) SetDefaultConversationCreatorRoleSid(DefaultConversationCreatorRoleSid string) (*UpdateServiceConfigurationParams){
    params.DefaultConversationCreatorRoleSid = &DefaultConversationCreatorRoleSid
    return params
}
func (params *UpdateServiceConfigurationParams) SetDefaultConversationRoleSid(DefaultConversationRoleSid string) (*UpdateServiceConfigurationParams){
    params.DefaultConversationRoleSid = &DefaultConversationRoleSid
    return params
}
func (params *UpdateServiceConfigurationParams) SetDefaultChatServiceRoleSid(DefaultChatServiceRoleSid string) (*UpdateServiceConfigurationParams){
    params.DefaultChatServiceRoleSid = &DefaultChatServiceRoleSid
    return params
}
func (params *UpdateServiceConfigurationParams) SetReachabilityEnabled(ReachabilityEnabled bool) (*UpdateServiceConfigurationParams){
    params.ReachabilityEnabled = &ReachabilityEnabled
    return params
}

// Update configuration settings of a conversation service
func (c *ApiService) UpdateServiceConfiguration(ChatServiceSid string, params *UpdateServiceConfigurationParams) (*ConversationsV1ServiceConfiguration, error) {
    path := "/v1/Services/{ChatServiceSid}/Configuration"
        path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)

data := url.Values{}
headers := make(map[string]interface{})

if params != nil && params.DefaultConversationCreatorRoleSid != nil {
    data.Set("DefaultConversationCreatorRoleSid", *params.DefaultConversationCreatorRoleSid)
}
if params != nil && params.DefaultConversationRoleSid != nil {
    data.Set("DefaultConversationRoleSid", *params.DefaultConversationRoleSid)
}
if params != nil && params.DefaultChatServiceRoleSid != nil {
    data.Set("DefaultChatServiceRoleSid", *params.DefaultChatServiceRoleSid)
}
if params != nil && params.ReachabilityEnabled != nil {
    data.Set("ReachabilityEnabled", fmt.Sprint(*params.ReachabilityEnabled))
}



    resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ConversationsV1ServiceConfiguration{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
