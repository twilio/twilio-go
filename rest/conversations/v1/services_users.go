/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.18.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateServiceUser'
type CreateServiceUserParams struct {
	// The X-Twilio-Webhook-Enabled HTTP request header
	XTwilioWebhookEnabled *string `json:"X-Twilio-Webhook-Enabled,omitempty"`
	// The JSON Object string that stores application-specific data. If attributes have not been set, `{}` is returned.
	Attributes *string `json:"Attributes,omitempty"`
	// The string that you assigned to describe the resource.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The application-defined string that uniquely identifies the resource's User within the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource). This value is often a username or an email address, and is case-sensitive.
	Identity *string `json:"Identity,omitempty"`
	// The SID of a service-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the user.
	RoleSid *string `json:"RoleSid,omitempty"`
}

func (params *CreateServiceUserParams) SetXTwilioWebhookEnabled(XTwilioWebhookEnabled string) *CreateServiceUserParams {
	params.XTwilioWebhookEnabled = &XTwilioWebhookEnabled
	return params
}
func (params *CreateServiceUserParams) SetAttributes(Attributes string) *CreateServiceUserParams {
	params.Attributes = &Attributes
	return params
}
func (params *CreateServiceUserParams) SetFriendlyName(FriendlyName string) *CreateServiceUserParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateServiceUserParams) SetIdentity(Identity string) *CreateServiceUserParams {
	params.Identity = &Identity
	return params
}
func (params *CreateServiceUserParams) SetRoleSid(RoleSid string) *CreateServiceUserParams {
	params.RoleSid = &RoleSid
	return params
}

// Add a new conversation user to your service
func (c *ApiService) CreateServiceUser(ChatServiceSid string, params *CreateServiceUserParams) (*ConversationsV1ServiceServiceUser, error) {
	path := "/v1/Services/{ChatServiceSid}/Users"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Identity != nil {
		data.Set("Identity", *params.Identity)
	}
	if params != nil && params.RoleSid != nil {
		data.Set("RoleSid", *params.RoleSid)
	}

	if params != nil && params.XTwilioWebhookEnabled != nil {
		headers["X-Twilio-Webhook-Enabled"] = *params.XTwilioWebhookEnabled
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ServiceServiceUser{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteServiceUser'
type DeleteServiceUserParams struct {
	// The X-Twilio-Webhook-Enabled HTTP request header
	XTwilioWebhookEnabled *string `json:"X-Twilio-Webhook-Enabled,omitempty"`
}

func (params *DeleteServiceUserParams) SetXTwilioWebhookEnabled(XTwilioWebhookEnabled string) *DeleteServiceUserParams {
	params.XTwilioWebhookEnabled = &XTwilioWebhookEnabled
	return params
}

// Remove a conversation user from your service
func (c *ApiService) DeleteServiceUser(ChatServiceSid string, Sid string, params *DeleteServiceUserParams) error {
	path := "/v1/Services/{ChatServiceSid}/Users/{Sid}"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.XTwilioWebhookEnabled != nil {
		headers["X-Twilio-Webhook-Enabled"] = *params.XTwilioWebhookEnabled
	}

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Fetch a conversation user from your service
func (c *ApiService) FetchServiceUser(ChatServiceSid string, Sid string) (*ConversationsV1ServiceServiceUser, error) {
	path := "/v1/Services/{ChatServiceSid}/Users/{Sid}"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ServiceServiceUser{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListServiceUser'
type ListServiceUserParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListServiceUserParams) SetPageSize(PageSize int) *ListServiceUserParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of all conversation users in your service
func (c *ApiService) ListServiceUser(ChatServiceSid string, params *ListServiceUserParams) (*ListServiceUserResponse, error) {
	path := "/v1/Services/{ChatServiceSid}/Users"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListServiceUserResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Retrieve a single page of  records from the API. Request is executed immediately.
func (c *ApiService) ServicesUsersPage(ChatServiceSid string, params *ListServiceUserParams, pageToken string, pageNumber string) (*ListServiceUserResponse, error) {
	path := "/v1/Services/{ChatServiceSid}/Users"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	data.Set("PageToken", pageToken)
	data.Set("PageNumber", pageNumber)

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListServiceUserResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Lists ServicesUsers records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) ServicesUsersList(ChatServiceSid string, params *ListServiceUserParams, limit int) ([]ListServiceUserResponse, error) {
	params.SetPageSize(c.requestHandler.ReadLimits(params.PageSize, limit))
	response, err := c.ListServiceUser(ChatServiceSid, params)
	if err != nil {
		return nil, err
	}

	page := client.NewPage(c.baseURL, response)

	resp := c.requestHandler.List(page, limit, 0)
	ret := make([]ListServiceUserResponse, len(resp))

	for i := range resp {
		jsonStr, _ := json.Marshal(resp[i])
		ps := ListServiceUserResponse{}
		if err := json.Unmarshal(jsonStr, &ps); err != nil {
			return ret, err
		}

		ret[i] = ps
	}

	return ret, nil
}

//Streams ServicesUsers records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) ServicesUsersStream(ChatServiceSid string, params *ListServiceUserParams, limit int) (chan interface{}, error) {
	params.SetPageSize(c.requestHandler.ReadLimits(params.PageSize, limit))
	response, err := c.ListServiceUser(ChatServiceSid, params)
	if err != nil {
		return nil, err
	}

	page := client.NewPage(c.baseURL, response)

	ps := ListServiceUserResponse{}
	return c.requestHandler.Stream(page, limit, 0, ps), nil
}

// Optional parameters for the method 'UpdateServiceUser'
type UpdateServiceUserParams struct {
	// The X-Twilio-Webhook-Enabled HTTP request header
	XTwilioWebhookEnabled *string `json:"X-Twilio-Webhook-Enabled,omitempty"`
	// The JSON Object string that stores application-specific data. If attributes have not been set, `{}` is returned.
	Attributes *string `json:"Attributes,omitempty"`
	// The string that you assigned to describe the resource.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The SID of a service-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the user.
	RoleSid *string `json:"RoleSid,omitempty"`
}

func (params *UpdateServiceUserParams) SetXTwilioWebhookEnabled(XTwilioWebhookEnabled string) *UpdateServiceUserParams {
	params.XTwilioWebhookEnabled = &XTwilioWebhookEnabled
	return params
}
func (params *UpdateServiceUserParams) SetAttributes(Attributes string) *UpdateServiceUserParams {
	params.Attributes = &Attributes
	return params
}
func (params *UpdateServiceUserParams) SetFriendlyName(FriendlyName string) *UpdateServiceUserParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateServiceUserParams) SetRoleSid(RoleSid string) *UpdateServiceUserParams {
	params.RoleSid = &RoleSid
	return params
}

// Update an existing conversation user in your service
func (c *ApiService) UpdateServiceUser(ChatServiceSid string, Sid string, params *UpdateServiceUserParams) (*ConversationsV1ServiceServiceUser, error) {
	path := "/v1/Services/{ChatServiceSid}/Users/{Sid}"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.RoleSid != nil {
		data.Set("RoleSid", *params.RoleSid)
	}

	if params != nil && params.XTwilioWebhookEnabled != nil {
		headers["X-Twilio-Webhook-Enabled"] = *params.XTwilioWebhookEnabled
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ServiceServiceUser{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
